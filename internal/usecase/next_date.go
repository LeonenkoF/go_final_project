package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	Year  = "y"
	Month = "m"
	Day   = "d"
	Week  = "w"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {

	repeatData := strings.Split(repeat, " ")

	if len(repeatData) == 0 {
		return "", fmt.Errorf("repeat cannot be empty")
	}
	parsedDate, err := time.Parse("20060102", date)
	if err != nil {
		return "", fmt.Errorf("runtime error time.Parse")
	}

	switch repeatData[0] {
	case Year:
		for {
			parsedDate = parsedDate.AddDate(1, 0, 0)
			if parsedDate.After(now) {
				break
			}
		}
	case Day:
		if len(repeatData) != 2 {
			return "", fmt.Errorf("wrong days count")
		}
		days, err := strconv.Atoi(repeatData[1])
		if err != nil {
			return "", fmt.Errorf("wrong days count, err:%s", err)
		}

		if days > 400 || days < 0 {
			return "", fmt.Errorf("wrong days count")
		}
		for {
			parsedDate = parsedDate.AddDate(0, 0, days)
			if parsedDate.After(now) {
				break
			}
		}
	default:
		return "", fmt.Errorf("unsupported formatted")
	}

	return parsedDate.Format("20060102"), nil

}
