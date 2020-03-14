package goanda_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"gopkg.in/h2non/gock.v1"

	. "github.com/gelder-io/goanda"
)

var _ = Describe("Pricing", func() {
	var c *Client

	BeforeEach(func() {
		c = NewClient("super-secret-api-key")
	})

	JustBeforeEach(func() {
		gock.InterceptClient(c.GetHTTPClient())
	})

	Describe("GetPricing", func() {
		It("correctly returns prices of currency pairs", func() {
			gock.New(HostURL).
				Get("/accounts/first/pricing").
				Reply(200).
				JSON(getPricingResponse)

			prices, err := c.GetPricing("first", []string{"EUR_USD"})

			Expect(err).ToNot(HaveOccurred())
			Expect(prices).To(HaveLen(1))
			Expect(prices[0].Instrument).To(Equal("EUR_USD"))
			Expect(prices[0].Time).To(Equal("2016-06-22T18:41:36.201836422Z"))
			Expect(prices[0].Tradeable).To(Equal(true))
			Expect(prices[0].Bids).To(HaveLen(2))
			Expect(prices[0].Bids[0].Liquidity).To(Equal(10000000))
			Expect(prices[0].Bids[0].Price).To(Equal(decimal.NewFromFloat(1.13015)))
			Expect(prices[0].Bids[1].Liquidity).To(Equal(10000000))
			Expect(prices[0].Bids[1].Price).To(Equal(decimal.NewFromFloat(1.13013)))
			Expect(prices[0].Asks).To(HaveLen(2))
			Expect(prices[0].Asks[0].Liquidity).To(Equal(10000000))
			Expect(prices[0].Asks[0].Price).To(Equal(decimal.NewFromFloat(1.13028)))
			Expect(prices[0].Asks[1].Liquidity).To(Equal(10000000))
			Expect(prices[0].Asks[1].Price).To(Equal(decimal.NewFromFloat(1.13031)))
		})
	})
})

const getPricingResponse = `
{
  "prices": [
    {
      "asks": [
        {
          "liquidity": 10000000, 
          "price": "1.13028"
        }, 
        {
          "liquidity": 10000000, 
          "price": "1.13031"
        }
      ], 
      "bids": [
        {
          "liquidity": 10000000, 
          "price": "1.13015"
        }, 
        {
          "liquidity": 10000000, 
          "price": "1.13013"
        }
      ], 
      "closeoutAsk": "1.13032", 
      "closeoutBid": "1.13011", 
      "instrument": "EUR_USD", 
      "quoteHomeConversionFactors": {
        "negativeUnits": "0.95904000", 
        "positiveUnits": "0.95886000"
      }, 
      "tradeable": true, 
      "time": "2016-06-22T18:41:36.201836422Z", 
      "unitsAvailable": {
        "default": {
          "long": "2013434", 
          "short": "2014044"
        }, 
        "openOnly": {
          "long": "2013434", 
          "short": "2014044"
        }, 
        "reduceFirst": {
          "long": "2013434", 
          "short": "2014044"
        }, 
        "reduceOnly": {
          "long": "0", 
          "short": "0"
        }
      }
    }
  ]
}`
