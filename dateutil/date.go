package dateutil

import (
	"strconv"
	"strings"
	"time"
)

type Date struct {
}

const (
	year            = "2006"
	month           = "01"
	day             = "02"
	hour            = "15"
	minute          = "04"
	second          = "05"
	complete        = "2006-01-02 15:05:05"
	stringToTimeOne = "2006-01-02 15:04:05"
	stringToTimeTow = "2006-01-02"

	DateTimeStr = "YYYY-MM-DD hh:mm:ss"
	DateStr     = "YYYY-MM-DD"
)

// FormatToString 时间格式化成字符串  Time forms a string
func FormatToString(t time.Time, s string) string {
	replace, b := replace(s)
	if !b {
		return t.Format(complete)
	}
	if IsZero(t) {
		return Now().Format(complete)
	}
	return t.Format(replace)
}

func StrToDate(timeStr string, s string) (time.Time, error) {
	tmpS, b := replace(s)
	if b {
		return time.ParseInLocation(tmpS, timeStr, time.Local)
	}
	return time.ParseInLocation(stringToTimeOne, timeStr, time.Local)
}

func CompareDate(date1, date2 time.Time) int {
	if date1.After(date2) {
		return 1
	} else if date1.Before(date2) {
		return -1
	} else {
		return 0
	}
}

//字符串替换
func replace(s string) (string, bool) {
	flag := false
	if strings.Contains(s, "YYYY") {
		s = strings.Replace(s, "YYYY", year, 1)
		flag = true
	}
	if strings.Contains(s, "MM") {
		s = strings.Replace(s, "MM", month, 1)
		flag = true
	}
	if strings.Contains(s, "DD") {
		s = strings.Replace(s, "DD", day, 1)
		flag = true
	}
	if strings.Contains(s, "hh") {
		s = strings.Replace(s, "hh", hour, 1)
		flag = true
	}
	if strings.Contains(s, "mm") {
		s = strings.Replace(s, "mm", minute, 1)
		flag = true
	}
	if strings.Contains(s, "ss") {
		s = strings.Replace(s, "ss", second, 1)
		flag = true
	}
	return s, flag
}

// IsZero 判断时间是否为零 Determine whether the time is zero
func IsZero(t time.Time) bool {
	return t.IsZero()
}

// Now 获取当前时间  Get the current time
func Now() time.Time {
	return time.Now()
}

// InterpretStringToTimestamp String to timestamp 字符串转时间戳
func InterpretStringToTimestamp(timeStr string, strFormat string) (int64, error) {
	var t int64
	loc, _ := time.LoadLocation("Local")
	replace, _ := replace(strFormat)
	t1, err := time.ParseInLocation(replace, timeStr, loc)
	if err != nil {
		return 0, err
	}
	t = t1.Unix()
	return t, err
}

// UnixToTime 时间戳转时间
// UnixToTime timestamp to time
func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// GetWeekDay 获取周几方法
//How to get the day of the week
func GetWeekDay(t time.Time) int {
	return int(t.Weekday())
}

// MinuteAddOrSub 时间分钟加减计算
func MinuteAddOrSub(t time.Time, num int64) time.Time {
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "m")
	return t.Add(m)
}

// HourAddOrSub 时间小时加减计算
func HourAddOrSub(t time.Time, num int64) time.Time {
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "h")
	return t.Add(m)
}

// DayAddOrSub 时间天加减计算
func DayAddOrSub(t time.Time, num int64) time.Time {
	num = num * 24
	s := strconv.FormatInt(num, 10)
	var m time.Duration
	m, _ = time.ParseDuration(s + "h")
	return t.Add(m)
}
