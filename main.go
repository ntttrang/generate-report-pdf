package main

import (
	"fmt"
	"generate-report-pdf/service"
)

func main() {
	fmt.Println("Start: Generate report")
	err := service.GenerateReport()
	if err != nil {
		fmt.Println("Generate report failed: ", err)
	}
	fmt.Println("End: Generate report")
}
