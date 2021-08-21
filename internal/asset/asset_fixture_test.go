package asset_test

import (
	c "github.com/achannarasappa/ticker/internal/common"
)

var fixtureAssetQuotes = []c.AssetQuote{
	{
		Name:     "ThoughtWorks",
		Symbol:   "TWKS",
		Class:    c.AssetClassStock,
		Currency: c.Currency{FromCurrencyCode: "USD"},
		QuotePrice: c.QuotePrice{
			Price:          110.0,
			PricePrevClose: 100.0,
			PriceOpen:      100.0,
			PriceDayHigh:   110.0,
			PriceDayLow:    90.0,
			Change:         10.0,
			ChangePercent:  10.0,
		},
	},
	{
		Name:     "Microsoft Inc",
		Symbol:   "MSFT",
		Class:    c.AssetClassStock,
		Currency: c.Currency{FromCurrencyCode: "USD"},
		QuotePrice: c.QuotePrice{
			Price:          220.0,
			PricePrevClose: 200.0,
			PriceOpen:      200.0,
			PriceDayHigh:   220.0,
			PriceDayLow:    180.0,
			Change:         20.0,
			ChangePercent:  10.0,
		},
	},
	{
		Name:     "Solana USD",
		Symbol:   "SOL1-USD",
		Class:    c.AssetClassCryptocurrency,
		Currency: c.Currency{FromCurrencyCode: "USD"},
		QuotePrice: c.QuotePrice{
			Price:          11.0,
			PricePrevClose: 10.0,
			PriceOpen:      10.0,
			PriceDayHigh:   11.0,
			PriceDayLow:    9.0,
			Change:         1.0,
			ChangePercent:  10.0,
		},
	},
}
var fixtureAssets = []c.Asset{
	{
		Name:     "ThoughtWorks",
		Symbol:   "TWKS",
		Class:    c.AssetClassStock,
		Currency: c.Currency{FromCurrencyCode: "USD", ToCurrencyCode: "USD"},
		QuotePrice: c.QuotePrice{
			Price:          110.0,
			PricePrevClose: 100.0,
			PriceOpen:      100.0,
			PriceDayHigh:   110.0,
			PriceDayLow:    90.0,
			Change:         10.0,
			ChangePercent:  10.0,
		},
		Meta: c.Meta{
			OrderIndex: 0,
		},
	},
	{
		Name:     "Microsoft Inc",
		Symbol:   "MSFT",
		Class:    c.AssetClassStock,
		Currency: c.Currency{FromCurrencyCode: "USD", ToCurrencyCode: "USD"},
		QuotePrice: c.QuotePrice{
			Price:          220.0,
			PricePrevClose: 200.0,
			PriceOpen:      200.0,
			PriceDayHigh:   220.0,
			PriceDayLow:    180.0,
			Change:         20.0,
			ChangePercent:  10.0,
		},
		Meta: c.Meta{
			OrderIndex: 1,
		},
	},
	{
		Name:     "Solana USD",
		Symbol:   "SOL1-USD",
		Class:    c.AssetClassCryptocurrency,
		Currency: c.Currency{FromCurrencyCode: "USD", ToCurrencyCode: "USD"},
		QuotePrice: c.QuotePrice{
			Price:          11.0,
			PricePrevClose: 10.0,
			PriceOpen:      10.0,
			PriceDayHigh:   11.0,
			PriceDayLow:    9.0,
			Change:         1.0,
			ChangePercent:  10.0,
		},
		Meta: c.Meta{
			OrderIndex: 2,
		},
	},
}
