package main

import "time"

func main() {
	for {
		pingContainers()
		time.Sleep(10 * time.Second)
	}
}
