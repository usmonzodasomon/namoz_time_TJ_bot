package handler

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"regexp"
	"strings"
	"time"
)

func (h *Handler) TaqvimHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	times, err := getNamazTimes()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	fmt.Println(times)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        times,
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func getNamazTimes() (string, error) {
	// ChromeDP options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	// Create context with options
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create browser context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Set timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Variable to store HTML
	var htmlContent string

	// Get HTML from the page
	fmt.Println("Starting page loading...")
	err := chromedp.Run(ctx,
		chromedp.Navigate("http://www.taqvim.tj/"),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		return "", fmt.Errorf("browser error: %v", err)
	}

	fmt.Println("Parsing prayer times from the central mosque...")
	masjidiMarkaziTimes, err := getMasjidiMarkaziTimes(htmlContent)
	if err == nil {
		fmt.Println("Found prayer times for the central mosque")
		return masjidiMarkaziTimes, nil
	}
	return "error", nil
}

func getMasjidiMarkaziTimes(htmlContent string) (string, error) {
	// Regular expression to find the central mosque table
	pattern := regexp.MustCompile(`<table id="table_namoz_time_today">.*?Вақтҳои намоз дар масҷиди Марказии шаҳри Душанбе.*?</table>`)
	tableHTML := pattern.FindString(htmlContent)

	if tableHTML == "" {
		return "", fmt.Errorf("central mosque prayer times table not found")
	}

	// Parse date
	datePattern := regexp.MustCompile(`Имрӯз: (\d{2}-\d{2}-\d{4})`)
	dateMatch := datePattern.FindStringSubmatch(tableHTML)
	date := ""
	if len(dateMatch) > 1 {
		date = dateMatch[1]
	}

	// Parse prayer times
	prayerPattern := regexp.MustCompile(`<th class="th_namoz_time_today">(Бомдод|Пешин|Аср|Шом|Хуфтан)</th><td class="td_namoz_time_today">(\d{2}:\d{2})</td>`)
	prayerMatches := prayerPattern.FindAllStringSubmatch(tableHTML, -1)

	if len(prayerMatches) == 0 {
		return "", fmt.Errorf("prayer times not found in the table")
	}

	// Format response
	var response strings.Builder
	response.WriteString(fmt.Sprintf("Вақтҳои намоз дар масҷиди Марказии шаҳри Душанбе, Имрӯз: %s\n", date))

	for _, match := range prayerMatches {
		prayerName := match[1]
		prayerTime := match[2]
		response.WriteString(fmt.Sprintf("%s: %s\n", prayerName, prayerTime))
	}

	return response.String(), nil
}
