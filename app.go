package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculate(a float64, p float64) float64 {
	percent := p / 100.00
	result := a + (a * percent)
	return result
}

func showMenu() {
	fmt.Println("Sales Tax Calculator")
	fmt.Println("================================")
}

func showTotal(r float64) {
	fmt.Println("Total")
	fmt.Println("================================")
	fmt.Printf("$%.2f\n================================\n\n", r)
}

func saveToFile(amount float64, percent float64, result float64) {
	file, err := os.OpenFile("calculations.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "Amount: $%.2f, Tax Rate: %.2f%%,\nR=at\nResult: $%.2f", amount, percent, result)
	writer.Flush()
}

func mainLoop() {
	for {
		var amount float64
		var percent float64

		showMenu()
		fmt.Print("Enter the amount ($): ")
		fmt.Scanln(&amount)
		fmt.Print("Enter the sales tax (%): ")
		fmt.Scanln(&percent)

		result := calculate(amount, percent)

		showTotal(result)

		saveToFile(amount, percent, result)

		var quit string

		fmt.Println("Do you want to quit the app? (y/n): ")
		fmt.Scanln(&quit)

		if quit == "y" {
			break
		} else if quit == "n" {
			continue
		} else {
			continue
		}
	}
}

func main() {
	mainLoop()
}
