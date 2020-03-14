package goanda

// GetAccounts returns the IDs of all accounts registered with the broker.
func (c *Client) GetAccounts() (ids []string, err error) {
	result := &struct {
		Accounts []AccountProperties
	}{}

	if _, err = c.rc.R().SetResult(result).Get("/accounts"); err != nil {
		return
	}

	for _, acc := range result.Accounts {
		ids = append(ids, acc.ID)
	}

	return
}

// GetAccount returns the IDs of all accounts registered with the broker.
func (c *Client) GetAccount(accountID string) (acc Account, err error) {
	result := &struct {
		Account Account
	}{}

	if _, err = c.rc.R().SetResult(result).Get("/accounts/" + accountID); err != nil {
		return
	}

	acc = result.Account

	return
}
