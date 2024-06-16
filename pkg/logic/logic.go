package logic

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func IsCorrectEnter(text string) (string, float64, error) {
	words := strings.Split(text, "-")
	if len(words) != 2 {
		err := errors.New("неверный формат ввода")
		return "", 0.0, err
	} else {
		sum, err := strconv.ParseFloat(words[1], 64)
		if err != nil {
			return "", 0.0, err
		}

		return words[0], sum, nil
	}
}

func GetTime() string {
	currentTime := time.Now().Format("01/02/2006")
	return currentTime
}

func HandleTime(interval string) string {
	currenctTime := time.Now()
	if interval == "week" {
		weekDay := int(currenctTime.Weekday())
		switch weekDay {
		case 0:
			return currenctTime.AddDate(0, 0, -6).Format("01/02/2006")
		case 1:
			return currenctTime.AddDate(0, 0, 0).Format("01/02/2006")
		case 2:
			return currenctTime.AddDate(0, 0, -1).Format("01/02/2006")
		case 3:
			return currenctTime.AddDate(0, 0, -2).Format("01/02/2006")
		case 4:
			return currenctTime.AddDate(0, 0, -3).Format("01/02/2006")
		case 5:
			return currenctTime.AddDate(0, 0, -4).Format("01/02/2006")
		case 6:
			return currenctTime.AddDate(0, 0, -5).Format("01/02/2006")
		default:
			return currenctTime.AddDate(0, 0, 0).Format("01/02/2006")
		}
	} else if interval == "month" {
		return time.Date(currenctTime.Year(), currenctTime.Month(), 01, 0, 0, 0, 0, time.Local).Format("01/02/2006")
	} else if interval == "year" {
		return time.Date(currenctTime.Year(), 01, 01, 0, 0, 0, 0, time.Local).Format("01/02/2006")
	} else {
		return "01/01/2000"
	}
}

func GetDateMonth(month string) (startDateInterval, endDateInterval string) {
	currenctTime := time.Now()
	switch month {
	case "Jan":
		startDateInterval := time.Date(currenctTime.Year(), 01, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Fab":
		startDateInterval := time.Date(currenctTime.Year(), 02, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Mar":
		startDateInterval := time.Date(currenctTime.Year(), 03, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Apr":
		startDateInterval := time.Date(currenctTime.Year(), 04, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "May":
		startDateInterval := time.Date(currenctTime.Year(), 05, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Jun":
		startDateInterval := time.Date(currenctTime.Year(), 06, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Jul":
		startDateInterval := time.Date(currenctTime.Year(), 07, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Avg":
		startDateInterval := time.Date(currenctTime.Year(), 8, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Sep":
		startDateInterval := time.Date(currenctTime.Year(), 9, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Oct":
		startDateInterval := time.Date(currenctTime.Year(), 10, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Nov":
		startDateInterval := time.Date(currenctTime.Year(), 11, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	case "Dec":
		startDateInterval := time.Date(currenctTime.Year(), 12, 01, 0, 0, 0, 0, time.Local)
		endDateInterval := startDateInterval.AddDate(0, 1, 0)
		return startDateInterval.Format("01/02/2006"), endDateInterval.Format("01/02/2006")
	default:
		return "01/01/2000", "01/01/2001"
	}

}
