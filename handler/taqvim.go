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
	url := launcher.New().
		Bin("/usr/bin/chromium-browser").
		Headless(true).
		NoSandbox(true).
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// Загружаем страницу
	page := browser.MustPage("http://www.taqvim.tj/")
	page.MustWaitLoad()
	time.Sleep(5 * time.Second) // подождём, пока загрузится JS

	// Получаем HTML-контент страницы
	htmlContent, err := page.HTML()
	if err != nil {
		return "", fmt.Errorf("failed to get HTML: %v", err)
	}

	// Сохраняем HTML для отладки
	err = os.WriteFile("NamazTime.html", []byte(htmlContent), 0644)
	if err != nil {
		log.Printf("Warning: could not save HTML to file: %v", err)
	}

	fmt.Println("Parsing prayer times from the central mosque...")

	// Парсим HTML
	masjidiMarkaziTimes, err := getMasjidiMarkaziTimes(htmlContent)
	if err == nil {
		fmt.Println("Found prayer times for the central mosque")
		return masjidiMarkaziTimes, nil
	}
	return "Sorry, information not found. Check the NamazTime.html file for debugging.", nil
}

func getMasjidiMarkaziTimes(htmlContent string) (string, error) {
	pattern := regexp.MustCompile(`<table id="table_namoz_time_today">.*?Вақтҳои намоз дар масҷиди Марказии шаҳри Душанбе.*?</table>`)
	tableHTML := pattern.FindString(htmlContent)

	if tableHTML == "" {
		return "", fmt.Errorf("central mosque prayer times table not found")
	}

	datePattern := regexp.MustCompile(`Имрӯз: (\d{2}-\d{2}-\d{4})`)
	dateMatch := datePattern.FindStringSubmatch(tableHTML)
	date := ""
	if len(dateMatch) > 1 {
		date = dateMatch[1]
	}

	prayerPattern := regexp.MustCompile(`<th class="th_namoz_time_today">(Бомдод|Пешин|Аср|Шом|Хуфтан)</th><td class="td_namoz_time_today">(\d{2}:\d{2})</td>`)
	prayerMatches := prayerPattern.FindAllStringSubmatch(tableHTML, -1)

	if len(prayerMatches) == 0 {
		return "", fmt.Errorf("prayer times not found in the table")
	}

	var response strings.Builder
	response.WriteString(fmt.Sprintf("Вақтҳои намоз дар масҷиди Марказии шаҳри Душанбе, Имрӯз: %s\n", date))

	for _, match := range prayerMatches {
		prayerName := match[1]
		prayerTime := match[2]
		response.WriteString(fmt.Sprintf("%s: %s\n", prayerName, prayerTime))
	}

	return response.String(), nil
}
