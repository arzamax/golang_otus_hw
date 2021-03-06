package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	result, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalln("Error with exact time fetching:", err)
	}
	t := time.Now().Truncate(time.Second)
	fmt.Println("current time:", t)
	fmt.Println("exact time:", result.Truncate(time.Second))
}
