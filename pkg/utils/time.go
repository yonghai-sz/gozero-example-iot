package utils

import (
	"fmt"
	"time"
)

func ParseTime(value string, loc *time.Location) (timeIns time.Time, err error) {
	timeLayout := "060102150405"
	timeIns, err = time.ParseInLocation(timeLayout, value, loc)
	return
}

func FormatTime(timeIns time.Time) string {
	year := timeIns.Year()     // 年
	month := timeIns.Month()   // 月
	day := timeIns.Day()       // 日
	hour := timeIns.Hour()     // 小时
	minute := timeIns.Minute() // 分钟
	second := timeIns.Second() // 秒
	yearDivision := 100        // 年取后两位
	timeFormat := "%02d%02d%02d%02d%02d%02d"
	return fmt.Sprintf(timeFormat, year%yearDivision, month, day, hour, minute, second)
}
