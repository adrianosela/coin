package main

import (
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	hashSize = 32
	blobSize = 32
	coinName = "CPEN 442 Coin2019"
	hashPref = "00000000"
)

var (
	// program keeps track of already used blobs
	checked = make(map[string]bool)
)

// Coin is a crypto currency
type Coin struct {
	p string // hash of previous coin
	h string // hash of current coin
	b string // coin blob - proof of work
	m string // miner id
}

// Mine mines coin
func Mine(prev, id string) (*Coin, error) {
	// validate previous hash
	if len(prev) != hashSize {
		return nil, errors.New("hash of previous coin must be 32 bytes long")
	}
	// create coin object
	c := &Coin{
		p: prev,
		m: id,
		h: "", // not computed
	}
	// do work
	for !strings.HasPrefix(c.h, hashPref) {
		// new unused random blob value
		for {
			blob := [blobSize]byte{}
			io.ReadFull(rand.Reader, blob[:])
			c.b = string(blob[:])
			if _, ok := checked[c.b]; !ok {
				checked[c.b] = true
				break
			}
		}
		// MD5(coin_name + hash_of_preceding_coin + coin_blob + id_of_miner)
		data := []byte(fmt.Sprintf("%s%s%s%s", coinName, c.p, c.b, c.m))
		c.h = fmt.Sprintf("%x", md5.Sum(data))
	}
	return c, nil
}
