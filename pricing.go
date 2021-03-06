package goanda

import (
	"bufio"
	"encoding/json"
	"strings"
)

// GetPricing returns pricing for the given instruments.
func (c *Client) GetPricing(accountID string, instruments []string) (prices []Pricing, err error) {
	result := &struct {
		Prices []Pricing
	}{}

	_, err = c.rc.R().
		SetResult(result).
		SetQueryParam("instruments", strings.Join(instruments, ",")).
		Get("/accounts/" + accountID + "/pricing")

	if err != nil {
		return
	}

	prices = result.Prices

	return
}

// GetPricingStream returns pricing for the given instruments.
func (c *Client) GetPricingStream(accountID string, instruments []string) (ch chan Pricing, err error) {
	ch = make(chan Pricing)

	resp, err := c.rc.R().
		SetDoNotParseResponse(true).
		SetQueryParam("instruments", strings.Join(instruments, ",")).
		Get(HostURLStream + "/accounts/" + accountID + "/pricing/stream")

	r := bufio.NewReader(resp.RawBody())

	go func() {
		for {
			b, err := r.ReadBytes('\n')
			if err != nil {
				continue
			}

			var price Pricing
			if err := json.Unmarshal(b, &price); err != nil {
				continue
			}

			ch <- price
		}
	}()

	return
}
