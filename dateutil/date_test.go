package dateutil

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatToString(t *testing.T) {
	toString := FormatToString(time.Now(), DateTimeStr)
	fmt.Println(toString)
	date, err := StrToDate(toString, DateTimeStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(date)
}
