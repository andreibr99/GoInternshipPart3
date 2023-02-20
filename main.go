package main

import (
	"GoInternshipPart3/datamanager"
	"fmt"
)

func main() {
	payday := 31
	daysUntilNext, paydayDate := datamanager.DaysUntilNextPayday(payday)
	fmt.Printf("There are %v days until the payday that is on: %v\n", daysUntilNext, paydayDate)
	paydayList := datamanager.PaydayList(payday)
	for _, v := range paydayList {
		fmt.Println(v)
	}
}
