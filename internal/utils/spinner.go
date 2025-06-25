package utils

import (
	"fmt"
	"time"
)

func StartSpinner(message string) func() {
	stop := make(chan bool)
	go func() {
		frames := []rune{'|', '/', '-', '\\'}
		i := 0
		for {
			select {
			case <-stop:
				fmt.Printf("\r%s... Done!\n", message)
				return
			default:
				fmt.Printf("\r%s... %c", message, frames[i%len(frames)])
				time.Sleep(100 * time.Millisecond)
				i++
			}
		}
	}()
	return func() {
		stop <- true
	}
}
