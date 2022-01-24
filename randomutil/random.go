package randomutil

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomCharacter(targetStr string, count int) string {
	bytes := []byte(targetStr)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// pseudo-random string
func GetRandomString(count int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return GetRandomCharacter(str, count)
}

// pseudo-random numeric
func GetRandomNumeric(count int) string {
	str := "0123456789"
	return GetRandomCharacter(str, count)
}

// pseudo-random number in [min,max)
func RandInt(min int, max int) int64 {
	if min > max {
		panic(fmt.Sprintf("min: %d can not greater then max: %d", min, max))
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int64(min + r.Intn(max-min))
}
