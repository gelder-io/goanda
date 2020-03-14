package goanda_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"gopkg.in/h2non/gock.v1"

	. "github.com/gelder-io/goanda"
)

var _ = Describe("Position", func() {
	var c *Client

	BeforeEach(func() {
		c = NewClient("super-secret-api-key")
	})

	JustBeforeEach(func() {
		gock.InterceptClient(c.GetHTTPClient())
	})

	Describe("GetPositions", func() {
		It("returns an empty list if there are no positions", func() {
			gock.New(HostURL).
				Get("/accounts/first/positions").
				Reply(200).
				JSON(getPositionsEmptyResponse)

			positions, err := c.GetPositions("first")

			Expect(err).ToNot(HaveOccurred())
			Expect(positions).To(BeEmpty())
		})

		It("correctly returns the positions of an account", func() {
			gock.New(HostURL).
				Get("/accounts/first/positions").
				Reply(200).
				JSON(getPositionsResponse)

			positions, err := c.GetPositions("first")

			Expect(err).ToNot(HaveOccurred())
			Expect(positions).To(HaveLen(1))
			Expect(positions[0].Instrument).To(Equal("EUR_CHF"))
			Expect(positions[0].PL).To(Equal(decimal.NewFromFloat(-486.16662)))
			Expect(positions[0].ResettablePL).To(Equal(decimal.NewFromFloat(-486.16662)))
			Expect(positions[0].UnrealizedPL).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(positions[0].Long.PL).To(Equal(decimal.NewFromFloat(-483.91941)))
			Expect(positions[0].Long.ResettablePL).To(Equal(decimal.NewFromFloat(-483.91941)))
			Expect(positions[0].Long.Units).To(Equal(decimal.NewFromFloat(0)))
			Expect(positions[0].Long.UnrealizedPL).To(Equal(decimal.NewFromFloat(0.00001)))
			Expect(positions[0].Short.PL).To(Equal(decimal.NewFromFloat(-2.24721)))
			Expect(positions[0].Short.ResettablePL).To(Equal(decimal.NewFromFloat(-2.24721)))
			Expect(positions[0].Short.Units).To(Equal(decimal.NewFromFloat(0)))
			Expect(positions[0].Short.UnrealizedPL).To(Equal(decimal.NewFromFloat(0.00001)))
		})
	})
})

const getPositionsEmptyResponse = `
{
  "positions": []
}
`

const getPositionsResponse = `
{
	"positions": [
		{
			"instrument": "EUR_CHF", 
			"long": {
				"pl": "-483.91941", 
				"resettablePL": "-483.91941", 
				"units": "0", 
				"unrealizedPL": "0.00001"
			}, 
			"pl": "-486.16662", 
			"resettablePL": "-486.16662", 
			"short": {
				"pl": "-2.24721", 
				"resettablePL": "-2.24721", 
				"units": "0", 
				"unrealizedPL": "0.00001"
			}, 
			"unrealizedPL": "0.00001"
		}
	]
}
`
