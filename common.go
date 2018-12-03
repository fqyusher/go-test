package common

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Calculate seconds of one day
func GetOneDaySeconds(convertDate string) int64 {
	dd, _ := time.Parse("20060102", convertDate)
	return dd.Unix()
}

func GetRandNumber(avgNum float64, randSeed float64) float64 {
	return avgNum - randSeed/2 + rand.Float64()*randSeed
}

// Get certain precision of float
func Round(f float64, n int) (result float64) {
	pow10Num := math.Pow10(n)
	s := strconv.FormatFloat(f, 'f', 10, 64)
	sf, _ := strconv.ParseFloat(s, 64)
	index := strings.Index(s, ".")
	roundType := 0
	for i := index + n + 1; i < len(s); i++ {
		// according to ASCII, 0 is 48, and growing by number adding
		if s[i] >= 53 {
			roundType = 1
			break
		} else if s[i] < 52 {
			break
		}
	}

	if roundType == 1 {
		return math.Ceil(sf*pow10Num) / pow10Num
	}
	return math.Floor(sf*pow10Num) / pow10Num
}

// Get values of per-hour in on day
func GetPerHourValues() ([24]float64, float64) {
	var (
		hours [24]float64
		total float64
	)

	for i := 0; i < 24; i++ {
		if i < 2 {
			hours[i] = Round(5 + rand.Float64() * 4, 2)
		} else if i < 7 {
			hours[i] = Round(2.5 + rand.Float64() * 3, 2)
		} else if i < 10 {
			hours[i] = Round(4 + rand.Float64() * 5, 2)
		} else if i < 13 {
			hours[i] = Round(6.5  + rand.Float64() * 5, 2)
		} else if i < 16 {
			hours[i] = Round(8.5 + rand.Float64() * 5, 2)
		} else if i < 17 {
			hours[i] = Round(11 + rand.Float64() * 2, 2)
		} else if i < 21 {
			hours[i] = Round(13 + rand.Float64() * 5, 2)
		} else if i < 23 {
			hours[i] = Round(10 + rand.Float64() * 4, 2)
		} else {
			hours[i] = Round(7 + rand.Float64() * 5, 2)
		}
		total += hours[i]
	}

	return hours, total
}

// Get imei
// iType: 0-android, 1-ios
func GetRandImeiString(strLength int, iType int) (imei string) {
	seedStr := "0123456789"
	if iType == 0 {
		seedStr += "abcdef"
	} else {
		seedStr += "ABCDEF"
	}

	for i := 0; i < strLength; i++ {
		if i == 8 || i == 12 || i == 16 || i == 20 {
			imei += "-"
		}
		imei += string(seedStr[rand.Intn(len(seedStr))])
	}
	return
}

// Get next month
func AddOneMonth(yearMonth string) string {
	m, _ := strconv.Atoi(yearMonth[4:])
	if m == 12 {
		y, _ := strconv.Atoi(yearMonth[0:4])
		return strconv.Itoa(y+1) + "01"
	}

	position := 5
	if m >= 9 {
		position = 4
	}

	return yearMonth[0:position] + strconv.Itoa(m+1)
}

// Get next day
func AddOneDay(convertDate string) string {
	dd, _ := time.Parse("20060102", convertDate)
	nextDay := dd.AddDate(0, 0, 1)
	return nextDay.Format("20060102")
}

// Calculate days of one month
func CalDaysOfMonth(convertDate string) int64 {
	dd, _ := time.Parse("20060102", convertDate)
	lastDay := dd.AddDate(0, 0, -1)
	_, _, days := lastDay.Date()

	return int64(days)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
