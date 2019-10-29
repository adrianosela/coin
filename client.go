package main

// Client is a CPEN 442 Coin API client
type Client struct {
	url string
}

// NewClient is the Client constructor
func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

// LastCoin returns the hash of the last coin mined
func (c *Client) LastCoin() (string, error) {
	// TODO
	return "", nil
}

// VerifyCoin base64 encodes a coin blob and verifies it with the API
func (c *Client) VerifyCoin(blob string) (bool, error) {
	// TODO
	return false, nil
}

// PublishCoin base64 encodes a coin blob and submits it to the API
func (c *Client) PublishCoin(blob string) (bool, error) {
	// TODO
	return false, nil
}
