package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		var choice int

		fmt.Println("Pomodoro Timer")
		fmt.Println("1. Start Work Session")
		fmt.Println("2. Start Short Break")
		fmt.Println("3. Start Long Break")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			startTimer(25, "Work Session")
		case 2:
			startTimer(5, "Short Break")
		case 3:
			startTimer(15, "Long Break")
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func startTimer(minutes int, sessionType string) {
	fmt.Printf("Starting %s for %d minutes...\n", sessionType, minutes)

	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)

	go func() {
		time.Sleep(time.Duration(minutes) * time.Minute)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Printf("%s is over!\n", sessionType)
			return
		case t := <-ticker.C:
			minutes--
			fmt.Printf("%s - %d minutes left (last tick at %s)\n", sessionType, minutes, t.Format("15:04:05"))
		}
	}
}
