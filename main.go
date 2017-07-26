package main

import (
	"fmt"
)

import "time"

func main() {

	ticker := time.NewTicker(time.Millisecond * 20000)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			SyncRepo()
		}
	}()

	select {}

}
