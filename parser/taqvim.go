package parser

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"regexp"
)

const (
	taqvimUrl = "http://www.taqvim.tj/"
)

func GetTaqvimNamazTime() (*types.TaqvimTime, error) {
	url := launcher.New().
		Bin("/usr/bin/chromium-browser").
		Headless(true).
		NoSandbox(true).
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(taqvimUrl)
	page.MustWaitLoad()

	htmlContent := page.MustHTML()

	taqvimTimes, err := getMasjidiMarkaziTimes(htmlContent)
	if err != nil {
		return nil, fmt.Errorf("error: could not find prayer times")
	}
	return taqvimTimes, nil
}

func getMasjidiMarkaziTimes(htmlContent string) (*types.TaqvimTime, error) {
	pattern := regexp.MustCompile(`<table id="table_namoz_time_today">.*?Вақтҳои намоз дар масҷиди Марказии шаҳри Душанбе.*?</table>`)
	tableHTML := pattern.FindString(htmlContent)

	if tableHTML == "" {
		return nil, fmt.Errorf("central mosque prayer times table not found")
	}

	prayerPattern := regexp.MustCompile(`<th class="th_namoz_time_today">(Бомдод|Пешин|Аср|Шом|Хуфтан)</th><td class="td_namoz_time_today">(\d{2}:\d{2})</td>`)
	prayerMatches := prayerPattern.FindAllStringSubmatch(tableHTML, -1)

	if len(prayerMatches) == 0 {
		return nil, fmt.Errorf("prayer times not found in the table")
	}

	taqvimTime := &types.TaqvimTime{}
	for _, match := range prayerMatches {
		prayerName := match[1]
		prayerTime := match[2]

		switch prayerName {
		case "Бомдод":
			taqvimTime.Fajr = prayerTime
		case "Пешин":
			taqvimTime.Zuhr = prayerTime
		case "Аср":
			taqvimTime.Asr = prayerTime
		case "Шом":
			taqvimTime.Maghrib = prayerTime
		case "Хуфтан":
			taqvimTime.Isha = prayerTime
		}
	}

	return taqvimTime, nil
}
