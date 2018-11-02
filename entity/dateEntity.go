package entity

import (
	"errors"
	"fmt"
	"strconv"
)

type Date struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

/*
 *	get and set
 */
func (m_date Date) GetYear() int {
	return m_date.Year
}

func (m_date *Date) SetYear(t_year int) {
	m_date.Year = t_year
}

func (m_date Date) GetMonth() int {
	return m_date.Month
}

func (m_date *Date) SetMonth(t_month int) {
	m_date.Month = t_month
}

func (m_date Date) GetDay() int {
	return m_date.Day
}

func (m_date *Date) SetDay(t_day int) {
	m_date.Day = t_day
}

func (m_date Date) GetHour() int {
	return m_date.Hour
}

func (m_date *Date) SetHour(t_hour int) {
	m_date.Hour = t_hour
}

func (m_date Date) GetMinute() int {
	return m_date.Minute
}

func (m_date *Date) SetMinute(t_minute int) {
	m_date.Minute = t_minute
}

/**
 *   @brief check whether the date is valid or not
 *   @return the bool indicate valid or not
 */
func IsValid(t_date Date) error {
	if t_date.Year < 1000 || t_date.Year > 9999 {
		return errors.New("invalid date")
	}
	if t_date.Month < 1 || t_date.Month > 12 {
		return errors.New("invalid date")
	}
	var days = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if t_date.Year%400 == 0 || t_date.Year%4 == 0 && t_date.Year%100 != 0 {
		days[2]++
	}
	if t_date.Day < 1 || t_date.Day > days[t_date.Month] {
		return errors.New("invalid date")
	}
	if t_date.Hour < 0 || t_date.Hour > 23 {
		return errors.New("invalid date")
	}
	if t_date.Minute < 0 || t_date.Minute > 59 {
		return errors.New("invalid date")
	}
	return nil
}

/**
 * @brief convert string to int
 */
func String2Int(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

/**
 * @brief convert a string to date, if the format is not correct return
 * 0000-00-00/00:00
 * @return a date
 */
func StringToDate(t_dateString string) (Date, error) {
	emptyDate := Date{0, 0, 0, 0, 0}
	len := len(t_dateString)
	info := [5]int{0, 0, 0, 0, 0}
	if len != 16 {
		return emptyDate, errors.New("date formal error")
	}
	for i, j := 0, 0; i < len; i++ {
		switch i {
		case 4, 7:
			if t_dateString[i] != '-' {
				return emptyDate, errors.New("date formal error")
			}
			j++
		case 10:
			if t_dateString[i] != '/' {
				return emptyDate, errors.New("date formal error")
			}
			j++
		case 13:
			if t_dateString[i] != ':' {
				return emptyDate, errors.New("date formal error")
			}
			j++
		default:
			if t_dateString[i] < '0' || t_dateString[i] > '9' {
				return emptyDate, errors.New("date formal error")
			} else {
				info[j] = info[j]*10 + String2Int(t_dateString[i:i+1])
			}
		}
	}
	resultDate := Date{info[0], info[1], info[2], info[3], info[4]}
	if err := IsValid(resultDate); err != nil {
		return emptyDate, err
	} else {
		return resultDate, nil
	}
}

/**
 *   @brief convert int to  string
 */
func Int2String(a int) string {
	result := strconv.Itoa(a)
	return result
}

/**
 * @brief convert a date to string, if the date is invalid return
 * 0000-00-00/00:00
 */
func DateToString(t_date Date) (string, error) {
	if err := IsValid(t_date); err != nil {
		return "0000-00-00/00:00", err
	} else {
		var dateString string
		dateString = Int2String(t_date.Year) + "-" + Int2String(t_date.Month) + "-" +
			Int2String(t_date.Day) + "/" + Int2String(t_date.Hour) + ":" + Int2String(t_date.Minute)
		return dateString, nil
	}
}

/**
 * @brief check whether the CurrentDate is equal to the t_date
 */
func (m_date Date) Equal(t_date Date) bool {
	return m_date.Year == t_date.Year &&
		m_date.Month == t_date.Month &&
		m_date.Day == t_date.Day &&
		m_date.Hour == t_date.Hour &&
		m_date.Minute == t_date.Minute
}

/**
 * @brief check whether the CurrentDate is  greater than the t_date
 */
func (m_date Date) Greater(t_date Date) bool {
	if m_date.Year > t_date.Year {
		return true
	}
	if m_date.Year < t_date.Year {
		return false
	}
	if m_date.Month > t_date.Month {
		return true
	}
	if m_date.Month < t_date.Month {
		return false
	}
	if m_date.Day > t_date.Day {
		return true
	}
	if m_date.Day < t_date.Day {
		return false
	}
	if m_date.Hour > t_date.Hour {
		return true
	}
	if m_date.Hour < t_date.Hour {
		return false
	}
	if m_date.Minute > t_date.Minute {
		return true
	}
	return false
}

/**
 * @brief check whether the CurrentDate is  less than the t_date
 */
func (m_date Date) Less(t_date Date) bool {
	if m_date.Year < t_date.Year {
		return true
	}
	if m_date.Year > t_date.Year {
		return false
	}
	if m_date.Month < t_date.Month {
		return true
	}
	if m_date.Month > t_date.Month {
		return false
	}
	if m_date.Day < t_date.Day {
		return true
	}
	if m_date.Day > t_date.Day {
		return false
	}
	if m_date.Hour < t_date.Hour {
		return true
	}
	if m_date.Hour > t_date.Hour {
		return false
	}
	if m_date.Minute < t_date.Minute {
		return true
	}
	return false
}

/**
 * @brief check whether the CurrentDate is  greater or equal than the t_date
 */
func (m_date Date) GreaterOrEqual(t_date Date) bool {
	return !m_date.Less(t_date)
}

/**
 * @brief check whether the CurrentDate is  less than or equal to the t_date
 */
func (m_date Date) LessOrEqual(t_date Date) bool {
	return !m_date.Greater(t_date)
}
