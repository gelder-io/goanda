package goanda

import (
	"errors"
)

// PlaceOrder places the given order.
func (c *Client) PlaceOrder(accountID string, order interface{}) (err error) {
	if !isOrderRequest(order) {
		err = errors.New("given order is not an order request")
		return
	}

	request := &struct {
		Order interface{} `json:"order"`
	}{Order: order}

	if _, err = c.rc.R().SetBody(request).Post("/accounts/" + accountID + "/orders"); err != nil {
		return
	}

	// TODO: Check what the transactions are in the response -- perhaps the order was cancelled.

	return
}
