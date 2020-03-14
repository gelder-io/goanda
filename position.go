package goanda

// GetPositions returns positions for all instruments that have had a position during the account's lifetime.
func (c *Client) GetPositions(accountID string) (positions []Position, err error) {
	result := &struct {
		Positions []Position
	}{}

	if _, err = c.rc.R().SetResult(result).Get("/accounts/" + accountID + "/positions"); err != nil {
		return
	}

	positions = result.Positions

	return
}
