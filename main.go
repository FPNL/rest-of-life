package main

import (
	"fmt"
	"time"
)

var (
	Now    float64 = float64(time.Now().UnixNano())
	Future float64 = float64(time.Date(1993+36, 2, 6, 12, 0, 0, 0, time.Local).UnixNano())
	p              = fmt.Println
)

func main() {
	Count()
}

func Count() {
	remainingPercentage := (Future - Now) / Future
	p("我的人生已經過了...", PassedTime(remainingPercentage)+"%")
	p("我人生還有...", RestTime(remainingPercentage)+"%")

}

func PassedTime(time float64) string {

	hasPassedPercentage := fmt.Sprintf("%.2f", 100-time*100)
	return hasPassedPercentage
}

func RestTime(time float64) string {
	stringMinusTime := fmt.Sprintf("%.2f", time*100) // float to String

	return stringMinusTime
}
