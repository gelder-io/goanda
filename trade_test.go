package goanda_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"gopkg.in/h2non/gock.v1"

	. "github.com/gelder-io/goanda"
)

var _ = Describe("Trade", func() {
	var c *Client

	BeforeEach(func() {
		c = New("super-secret-api-key")
	})

	JustBeforeEach(func() {
		gock.InterceptClient(c.GetHTTPClient())
	})

	Describe("GetOpenTrades", func() {
		It("returns an empty list if there are no trades", func() {
			gock.New(HostURL).
				Get("/accounts/first/openTrades").
				Reply(200).
				JSON(getTradesEmptyResponse)

			trades, err := c.GetOpenTrades("first")

			Expect(err).ToNot(HaveOccurred())
			Expect(trades).To(BeEmpty())
		})

		It("correctly returns trades", func() {
			gock.New(HostURL).
				Get("/accounts/first/openTrades").
				Reply(200).
				JSON(getTradesResponse)

			trades, err := c.GetOpenTrades("first")

			Expect(err).ToNot(HaveOccurred())
			Expect(trades).To(HaveLen(1))
			Expect(trades[0].ID).To(Equal("6395"))
			Expect(trades[0].Instrument).To(Equal("EUR_USD"))
			Expect(trades[0].State).To(Equal(Open))
			Expect(trades[0].Price).To(Equal(decimal.NewFromFloat(1.13033)))
			Expect(trades[0].InitialUnits).To(Equal(decimal.NewFromFloat(100.1)))
			Expect(trades[0].CurrentUnits).To(Equal(decimal.NewFromFloat(100.1)))
			Expect(trades[0].RealizedPL).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(trades[0].UnrealizedPL).To(Equal(decimal.NewFromFloat(-0.01438)))
			Expect(trades[0].OpenTime).To(Equal("2016-06-22T18:41:48.258142231Z"))
		})
	})
})

const getTradesEmptyResponse = `
{
  "trades": []
}`

const getTradesResponse = `
{
  "trades": [
    {
      "currentUnits": "100.1",
      "financing": "0.00000",
      "id": "6395",
      "initialUnits": "100.1",
      "instrument": "EUR_USD",
      "openTime": "2016-06-22T18:41:48.258142231Z",
      "price": "1.13033",
      "realizedPL": "0.00001",
      "state": "OPEN",
      "unrealizedPL": "-0.01438"
    }
  ]
}`
