package main

import (
	"fmt"
	"time"
	"github.com/gen2brain/beeep"
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
			beeep.Notify("Pomodoro Timer", fmt.Sprintf("%s is over!", sessionType), "")
			return
		case t := <-ticker.C:
			minutes--
			fmt.Printf("%s - %d minutes left (last tick at %s)\n", sessionType, minutes, t.Format("15:04:05"))
			if minutes%5 == 0 {
                beeep.Notify("Pomodoro Timer", fmt.Sprintf("%s - %d minutes left", sessionType, minutes), "")
            }
		}
	}
}
