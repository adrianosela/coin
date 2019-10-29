package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is a CPEN 442 Coin API client
type Client struct {
	url string
}

// NewClient is the Client constructor
func NewClient(url string) (*Client, error) {
	if url == "" {
		return nil, errors.New("no url provided")
	}
	return &Client{
		url: url,
	}, nil
}

// LastCoin returns the hash of the last coin mined
func (c *Client) LastCoin() (string, error) {
	// build request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/last_coin", c.url), nil)
	if err != nil {
		return "", fmt.Errorf("could not build http request: %s", err)
	}
	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not build http request: %s", err)
	}
	// read response
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("could not read http response: %s", err)
	}
	// unmarshall json response
	var body struct {
		CoinID string `json:"coin_id"`
	}
	if err = json.Unmarshal(bodyBytes, &body); err != nil {
		return "", fmt.Errorf("could not unmarshall http response: %s", err)
	}
	// return coin_id
	return body.CoinID, nil
}

// VerifyCoin base64 encodes a coin blob and submits it for verification
func (c *Client) VerifyCoin(blob, id string) (bool, error) {
	// build request body
	body := struct {
		CoinBlob string `json:"coin_blob"`
		MinerID  string `json:"id_of_miner"`
	}{
		CoinBlob: base64.StdEncoding.EncodeToString([]byte(blob)),
		MinerID:  id,
	}
	// marshall built json
	byt, err := json.Marshal(&body)
	if err != nil {
		return false, fmt.Errorf("could not marshall request body: %s", err)
	}
	// build request
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/verify_example_coin", c.url),
		bytes.NewBuffer(byt))
	if err != nil {
		return false, fmt.Errorf("could not build http request: %s", err)
	}
	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("could not send http request: %s", err)
	}
	// check status code
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil
}

// ClaimCoin base64 encodes a coin blob and submits it to the API
func (c *Client) ClaimCoin(blob, id string) (bool, error) {
	// build request body
	body := struct {
		CoinBlob string `json:"coin_blob"`
		MinerID  string `json:"id_of_miner"`
	}{
		CoinBlob: base64.StdEncoding.EncodeToString([]byte(blob)),
		MinerID:  id,
	}
	// marshall built json
	byt, err := json.Marshal(&body)
	if err != nil {
		return false, fmt.Errorf("could not marshall request body: %s", err)
	}
	// build request
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/claim_coin", c.url),
		bytes.NewBuffer(byt))
	if err != nil {
		return false, fmt.Errorf("could not build http request: %s", err)
	}
	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("could not send http request: %s", err)
	}
	// check status code
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil
}
