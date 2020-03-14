package goanda

// GetOpenTrades returns all open trades of an account.
func (c *Client) GetOpenTrades(accountID string) (trades []Trade, err error) {
	result := &struct {
		Trades []Trade
	}{}

	if _, err = c.rc.R().SetResult(result).Get("/accounts/" + accountID + "/openTrades"); err != nil {
		return
	}

	trades = result.Trades

	return
}
