package telegram

import (
	"log"
	"time"
)

var UpdateTimeTimeSleep = 1 * time.Hour

// Вызываем UpdateTimer каждый час чтобы обновить время намаза на сервере
// Если возникает ошибка, вызываем снова функцию парсера через минуту
// Если получаем ошибку 3 раза подряд, выходим из цикла и ждём следующего часа

func (b *Bot) UpdateTimeProcedure() {
	for {
		Time := time.Now().Format("02.01.2006")
		queryCnt := 0
		for {
			queryCnt++
			if err := b.Parser.Parse(Time); err != nil {
				log.Println("error parsing time: ", err.Error())
				if queryCnt == 3 {
					break
				}
				time.Sleep(1 * time.Minute)
			} else {
				break
			}
		}

		time.Sleep(UpdateTimeTimeSleep)
	}
}
