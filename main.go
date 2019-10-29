package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	id = "110b3259f9f189e603ad4142660d6945"
)

func main() {
	client, err := NewClient("http://cpen442coin.ece.ubc.ca")
	if err != nil {
		log.Fatal(err)
	}

	last, err := client.LastCoin()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("fetched last coin hash from API: %s\n", last) 
	fmt.Println("mining...")

	rand.Seed(time.Now().UTC().UnixNano())
	c, err := Mine(last, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("mined coin with blob base64: %s\n", base64.StdEncoding.EncodeToString([]byte(c.b)))

	ok, err := client.ClaimCoin(c.b, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("coin claim result: %v\n", ok)
}
