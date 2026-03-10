package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

func currentDayOfTheWeek() string {
	t := TimeNow()
	days := []string{
		"Воскресенье",
		"Понедельник",
		"Вторник",
		"Среда",
		"Четверг",
		"Пятница",
		"Суббота",
	}
	return days[int(t.Weekday())]
}

func dayOrNight() string {
	t := TimeNow()
	if hour := t.Hour(); hour >= 10 && hour <= 22 {
		return "День"
	} else {
		return "Ночь"
	}
}

func nextFriday() int {
	t := TimeNow()
	today := int(t.Weekday())
	friday := int(time.Friday)

	if today <= friday {
		return friday - today
	}
	return 7 - (today - friday)
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	correct := currentDayOfTheWeek()

	return strings.ToLower(answer) == strings.ToLower(correct)
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}

	correct := dayOrNight()

	return strings.ToLower(answer) == strings.ToLower(correct), nil
}
