package service

import (
	"crypto/sha1"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// HelperTest is the testing support struct.
type HelperTest struct{}

// RandStringBytes returns a random string
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandomString returns a random string
func RandomString(len int) string {
	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomNumber(97, 122))
	}

	return string(bytes)
}

// RandomNumber returns a random number.
func RandomNumber(min int64, max int64) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return int(min + rand.Int63n(max-min+1))
}

//RandomTime return a time in a day
//hx is the upper limit of the hour
//hy is the lower limit of the hour
//mx is the upper limit of the minute
//my is the upper limit of the minute
func RandomTime(hx, hy, mx, my int64) string {
	hour := strconv.Itoa(RandomNumber(hx, hy))
	minute := strconv.Itoa(RandomNumber(mx, my))

	if len(hour) < 2 {
		hour = "0" + hour
	}

	if len(minute) < 2 {
		minute = "0" + minute
	}

	return hour + ":" + minute + ":" + "00"
}

//RandomDate returns a random string format as YYYY-MM-DD
//Take in 2 parameters: day, month, year.
func RandomDate(day, month, year int) string {
	dd := strconv.Itoa(day)
	mm := strconv.Itoa(month)
	yy := strconv.Itoa(year)

	if len(dd) < 1 {
		dd = "0" + dd
	}

	if len(mm) < 1 {
		mm = "0" + mm
	}

	return yy + "-" + mm + "-" + dd
}

// This function generate the key based on the time of generating.
// Return a SHA1 hash string of the given information.
func GenerateKeyDoTest(ID int, EntityCode int, TestID int) string {
	var entityType string
	switch EntityCode {
	case 1:
		entityType = "Admin"
	case 2:
		entityType = "Student"
	case 3:
		entityType = "Teacher"
	}
	f_i := strconv.Itoa(ID)
	f_t := strconv.Itoa(TestID)

	key := strings.Join([]string{entityType, f_i, f_t}, "")
	h := sha1.New()

	h.Write([]byte(key))
	result := h.Sum(nil)

	return string(result)
}

func GenerateValueDoTest(Value int) string {
	f_c := strconv.Itoa(Value)
	value := strings.Join([]string{f_c}, "")

	return value
}
