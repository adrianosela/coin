package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const prev = "2e3a8e88a060cedcd9ac7b74fadd58e0"
const id = "110b3259f9f189e603ad4142660d6945"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	c, err := Mine(prev, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mined coin with blob base64: %s\n", base64.StdEncoding.EncodeToString([]byte(c.b)))
}
