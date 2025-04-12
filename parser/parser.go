package parser

import (
	"bytes"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const shuroiulamoUrl = "https://shuroiulamo.tj/tj/namaz/ntime"

func GetShuroNamazTimes(month, year string) ([]types.NamazTime, error) {
	formData := url.Values{}
	formData.Set("fmonth", month)
	formData.Set("fyear", year)
	formData.Set("fday", "0")

	// Отправка POST-запроса
	response, err := http.Post(shuroiulamoUrl, "application/x-www-form-urlencoded", bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	html := string(body)
	lines := strings.Split(html, "\n")

	namazTimes := make([]types.NamazTime, 0)
	cnt := -1
	var namazTime types.NamazTime
	for _, line := range lines {
		if strings.Contains(line, "<td>") {
			cnt++
			data := strings.ReplaceAll(line, "<td>", "")
			data = strings.ReplaceAll(data, "</td>", "")
			data = strings.TrimSpace(data)
			if cnt == 1 {
				namazTime.Date = data
			} else {
				sl := strings.Split(data, " - ")
				if cnt == 3 {
					namazTime.FajrFrom = sl[0]
					namazTime.FajrTo = sl[1]
				} else if cnt == 4 {
					namazTime.ZuhrFrom = sl[0]
					namazTime.ZuhrTo = sl[1]
				} else if cnt == 5 {
					namazTime.AsrFrom = sl[0]
					namazTime.AsrTo = sl[1]
				} else if cnt == 7 {
					namazTime.MaghribFrom = sl[0]
					namazTime.MaghribTo = sl[1]
				} else if cnt == 8 {
					namazTime.IshaFrom = sl[0]
					namazTime.IshaTo = sl[1]
					namazTimes = append(namazTimes, namazTime)
					cnt = -1
				}
			}

		}
	}
	return namazTimes, nil
}
