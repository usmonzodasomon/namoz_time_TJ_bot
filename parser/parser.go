package parser

import (
	"bytes"
	"echobot/types"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const urlString = "https://shuroiulamo.tj/tj/namaz/ntime"

type Parser struct {
	NamazTime *types.NamazTime
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(time string) error {
	// Подготовка данных формы
	formData := url.Values{}
	formData.Set("fmonth", time[3:5])
	formData.Set("fyear", time[6:])
	formData.Set("fday", time[0:2])

	// Отправка POST-запроса
	response, err := http.Post(urlString, "application/x-www-form-urlencoded", bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	bodyString := string(body)

	// Извлечение данных между <td></td> с помощью регулярного выражения
	re := regexp.MustCompile(`<td>(.*?)<\/td>`)
	matches := re.FindAllStringSubmatch(bodyString, -1)
	records := []string{}
	// Обработка данных из совпадений
	for _, match := range matches {
		if len(match) > 1 {
			records = append(records, match[1])
		}
	}
	if len(records) == 0 {
		return errors.New("len of records is 0")
	}

	stringNamazTime := types.StringNamazTime{
		Today:   records[1],
		Fajr:    records[3],
		Dhuhr:   records[4],
		Asr:     records[5],
		Maghrib: records[7],
		Isha:    records[8],
	}

	namazTime, err := GetNamazTime(stringNamazTime)
	if err != nil {
		return err
	}

	p.NamazTime = &namazTime
	return nil
}

func returnDay(s string) (time.Time, error) {
	return time.Parse("02.01.2006", s)
}

func GetNamazTime(v types.StringNamazTime) (types.NamazTime, error) {
	today, err := returnDay(v.Today)
	if err != nil {
		return types.NamazTime{}, err
	}
	fajr, err := returnNamazTime(v.Fajr)
	if err != nil {
		return types.NamazTime{}, err
	}
	dhuhr, err := returnNamazTime(v.Dhuhr)
	if err != nil {
		return types.NamazTime{}, err
	}
	asr, err := returnNamazTime(v.Asr)
	if err != nil {
		return types.NamazTime{}, err
	}
	maghrib, err := returnNamazTime(v.Maghrib)
	if err != nil {
		return types.NamazTime{}, err
	}
	isha, err := returnNamazTime(v.Isha)
	if err != nil {
		return types.NamazTime{}, err
	}
	return types.NamazTime{
		Today: today,
		Namaz: [5]types.NamazTimeStruct{fajr, dhuhr, asr, maghrib, isha},
	}, nil
}

func returnNamazTime(s string) (types.NamazTimeStruct, error) {
	v := strings.Split(s, " - ")
	from, err := time.Parse("15:04", v[0])
	if err != nil {
		return types.NamazTimeStruct{}, err
	}
	to, err := time.Parse("15:04", v[1])
	if err != nil {
		return types.NamazTimeStruct{}, err
	}
	return types.NamazTimeStruct{From: from, To: to}, nil
}
