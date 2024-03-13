package parser

import (
	"bytes"
	"echobot/types"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const urlString = "https://shuroiulamo.tj/tj/namaz/ntime"

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(month, year string) ([]types.NamazTime, error) {
	formData := url.Values{}
	formData.Set("fmonth", month)
	formData.Set("fyear", year)
	formData.Set("fday", "0")

	// Отправка POST-запроса
	response, err := http.Post(urlString, "application/x-www-form-urlencoded", bytes.NewBufferString(formData.Encode()))
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

//
//func returnDay(s string) (time.Time, error) {
//	return time.Parse("02.01.2006", s)
//}
//
//func GetNamazTime(v types.StringNamazTime) (types.NamazTime, error) {
//	today, err := returnDay(v.Today)
//	if err != nil {
//		return types.NamazTime{}, err
//	}
//	fajr, err := returnNamazTime(v.Fajr)
//	if err != nil {
//		return types.NamazTime{}, err
//	}
//	dhuhr, err := returnNamazTime(v.Dhuhr)
//	if err != nil {
//		return types.NamazTime{}, err
//	}
//	asr, err := returnNamazTime(v.Asr)
//	if err != nil {
//		return types.NamazTime{}, err
//	}
//	maghrib, err := returnNamazTime(v.Maghrib)
//	if err != nil {
//		return types.NamazTime{}, err
//	}
//	isha, err := returnNamazTime(v.Isha)
//	if err != nil {
//		return types.NamazTime{}, err
//	}
//	return types.NamazTime{
//		Today: today,
//		Namaz: [5]types.NamazTimeStruct{fajr, dhuhr, asr, maghrib, isha},
//	}, nil
//}
//
//func returnNamazTime(s string) (types.NamazTimeStruct, error) {
//	v := strings.Split(s, " - ")
//	from, err := time.Parse("15:04", v[0])
//	if err != nil {
//		return types.NamazTimeStruct{}, err
//	}
//	to, err := time.Parse("15:04", v[1])
//	if err != nil {
//		return types.NamazTimeStruct{}, err
//	}
//	return types.NamazTimeStruct{From: from, To: to}, nil
//}
