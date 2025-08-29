package main

import (
	"fmt"
	"time"
)

func CacheManager(comms Comms, interval time.Duration) {
	for {
		select {
		case <-time.After(interval):
			fmt.Println()
			ColorPrint(Yellow, "---------------------------")
			ColorPrint(Yellow, fmt.Sprintf("Clearing cache... interval: %v", interval))
			ColorPrint(Yellow, "---------------------------")
			fmt.Println()
		}

	}
}
