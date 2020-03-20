package goanda_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"gopkg.in/h2non/gock.v1"

	. "github.com/gelder-io/goanda"
)

var _ = Describe("Account", func() {
	var c *Client

	BeforeEach(func() {
		c = NewClient("super-secret-api-key")
	})

	JustBeforeEach(func() {
		gock.InterceptClient(c.GetHTTPClient())
	})

	Describe("GetAccounts", func() {
		It("correctly returns account IDs", func() {
			gock.New(HostURL).
				Get("/accounts").
				Reply(200).
				JSON(getAccountsResponse)

			ids, err := c.GetAccounts()

			Expect(err).ToNot(HaveOccurred())
			Expect(ids).To(ConsistOf("first", "second"))
		})
	})

	Describe("GetAccount", func() {
		It("correctly returns account information", func() {
			gock.New(HostURL).
				Get("/accounts/first").
				Reply(200).
				JSON(getFirstAccountResponse)

			acc, err := c.GetAccount("first")

			Expect(err).ToNot(HaveOccurred())
			Expect(acc.ID).To(Equal("first"))
			Expect(acc.Alias).To(Equal("My Account #1"))
			Expect(acc.Currency).To(Equal("CHF"))

			Expect(acc.OpenTradeCount).To(Equal(0))
			Expect(acc.OpenPositionCount).To(Equal(0))
			Expect(acc.PendingOrderCount).To(Equal(0))

			Expect(acc.HedgingEnabled).To(Equal(false))

			Expect(acc.MarginRate).To(Equal(decimal.NewFromFloat(0.02)))
			Expect(acc.UnrealizedPL).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(acc.NAV).To(Equal(decimal.NewFromFloat(43650.78835)))
			Expect(acc.MarginUsed).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(acc.MarginAvailable).To(Equal(decimal.NewFromFloat(43650.78835)))
			Expect(acc.PositionValue).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(acc.MarginCloseoutUnrealizedPL).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(acc.MarginCloseoutNAV).To(Equal(decimal.NewFromFloat(43650.78835)))
			Expect(acc.MarginCloseoutMarginUsed).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(acc.MarginCloseoutPercent).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(acc.WithdrawalLimit).To(Equal(decimal.NewFromFloat(43650.78835)))
			Expect(acc.Balance).To(Equal(decimal.NewFromFloat(43650.78835)))
			Expect(acc.PL).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(acc.ResettablePL).To(Equal(decimal.NewFromFloat(0.00001)))

			Expect(acc.Trades).To(BeEmpty())
			Expect(acc.Positions).To(BeEmpty())
			// Expect(acc.Orders).To(BeEmpty())
		})
	})

	Describe("GetAccountInstruments", func() {
		It("correctly returns account instruments", func() {
			gock.New(HostURL).
				Get("/accounts/first/instruments").
				Reply(200).
				JSON(getFirstAccountInstrumentsResponse)

			instruments, err := c.GetAccountInstruments("first")

			Expect(err).ToNot(HaveOccurred())
			Expect(instruments).To(HaveLen(2))

			Expect(instruments[0].Name).To(Equal("EUR_CHF"))
			Expect(instruments[0].DisplayName).To(Equal("EUR/CHF"))

			Expect(instruments[1].Name).To(Equal("EUR_USD"))
			Expect(instruments[1].DisplayName).To(Equal("EUR/USD"))
		})
	})
})

const getAccountsResponse = `
{
	"accounts": [
		{
			"id": "first",
			"tags": []
		}, 
		{
			"id": "second",
			"tags": []
		}
	]
}`

const getFirstAccountResponse = `
{
  "account": {
    "id": "first", 
    "alias": "My Account #1", 
    "currency": "CHF", 
    "openTradeCount": 0, 
    "openPositionCount": 0, 
    "pendingOrderCount": 0, 
    "hedgingEnabled": false, 
    "marginRate": "0.02", 
    "unrealizedPL": "0.00001", 
    "NAV": "43650.78835", 
    "marginUsed": "0.00001", 
    "marginAvailable": "43650.78835", 
    "positionValue": "0.00001", 
    "marginCloseoutUnrealizedPL": "0.00001", 
    "marginCloseoutNAV": "43650.78835", 
    "marginCloseoutMarginUsed": "0.00001", 
    "marginCloseoutPercent": "0.00001", 
    "marginCloseoutPositionValue": "0.00001", 
    "withdrawalLimit": "43650.78835",
    "balance": "43650.78835", 
    "pl": "0.00001", 
    "resettablePL": "0.00001", 
    "trades": [], 
    "positions": [], 
    "orders": []
  } 
}`

const getFirstAccountInstrumentsResponse = `
{
  "instruments": [
    {
      "displayName": "EUR/CHF", 
      "displayPrecision": 5, 
      "marginRate": "0.05", 
      "maximumOrderUnits": "100000000", 
      "maximumPositionSize": "0", 
      "maximumTrailingStopDistance": "1.00000", 
      "minimumTradeSize": "1", 
      "minimumTrailingStopDistance": "0.00050", 
      "name": "EUR_CHF", 
      "pipLocation": -4, 
      "tradeUnitsPrecision": 0, 
      "type": "CURRENCY"
    }, 
    {
      "displayName": "EUR/USD", 
      "displayPrecision": 5, 
      "marginRate": "0.02", 
      "maximumOrderUnits": "100000000", 
      "maximumPositionSize": "0", 
      "maximumTrailingStopDistance": "1.00000", 
      "minimumTradeSize": "1", 
      "minimumTrailingStopDistance": "0.00050", 
      "name": "EUR_USD", 
      "pipLocation": -4, 
      "tradeUnitsPrecision": 0, 
      "type": "CURRENCY"
    } 
	]
}`
