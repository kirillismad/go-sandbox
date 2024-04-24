package examples

import (
	"fmt"
	"time"
)

func DemoDateTime() {
	demoTimeFormat()
}

func demoTimeFormat() {
	// Day 					 2 or 02 or _2
	// Day of Week 			 Monday or Mon
	// Month 				 01 or 1 or Jan or January
	// Year 				 2006 or 06
	// Hour 				 03 or 3 or 15
	// Minutes 				 04 or 4
	// Seconds 				 05 or 5
	// Milli Seconds  (ms) 	 .000 (.999)
	// Micro Seconds (μs) 	 .000000 (.999999)
	// Nano Seconds (ns) 	 .000000000 (.999999999)
	// am/pm 				 PM or pm
	// Timezone 			 MST
	// Timezone offset 		 Z0700 or Z070000 or Z07 or Z07:00 or Z07:00:00  or -0700 or  -070000 or -07 or -07:00 or -07:00:00

	now := time.Now().UTC()

	fmt.Println("Default formatting", now)
	//Format castom
	fmt.Printf("Castom format: %v\n", now.Format("02.01.2006 - 15:04:05Z07:00"))

	fmt.Println("ISO 8601:", now.Format(time.RFC3339))

	//Format YYYY-MM-DD
	fmt.Printf("YYYY-MM-DD: %s\n", now.Format("2006-01-02"))
	//Format YY-MM-DD
	fmt.Printf("YY-MM-DD: %s\n", now.Format("06-01-02"))

	//Format YYYY-#{MonthName}-DD
	fmt.Printf("YYYY-#{MonthName}-DD: %s\n", now.Format("2006-Jan-02"))

	//Format HH:MM:SS
	fmt.Printf("HH:MM:SS: %s\n", now.Format("03:04:05"))

	//Format HH:MM:SS Millisecond
	fmt.Printf("HH:MM:SS Millisecond: %s\n", now.Format("03:04:05 .999"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS: %s\n", now.Format("2006-Jan-02 Monday 03:04:05"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset: %s\n", now.Format("2006-Jan-02 Monday 03:04:05 PM MST -07:00"))
}
