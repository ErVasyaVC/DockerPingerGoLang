package main

import (
	"fmt"
	"os"
)

func main() {
	countainerHost, err := os.Hostname()
	fmt.Println(countainerHost, err)
}
