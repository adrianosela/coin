package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	prev  = "2e3a8e88a060cedcd9ac7b74fadd58e0"
	id    = "110b3259f9f189e603ad4142660d6945"
	found = "T2JEe2cYC7HC5DMrgmmKiiUnG9XDAuEyW9YftTNyWWU="
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	// blob, _ := base64.StdEncoding.DecodeString(found)
	// fmt.Println(tryBlob(prev, string(blob), id))

	c, err := Mine(prev, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mined coin with blob base64: %s\n", base64.StdEncoding.EncodeToString([]byte(c.b)))
}
