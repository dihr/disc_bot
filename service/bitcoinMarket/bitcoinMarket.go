package bitcoinMarket

import (
	"fmt"
	"regexp"
)

type (
	BitcoinMarketSvc interface {
		GetCoins(coin string) map[string]string
	}

	bitcoinMarketImp struct {
	}
)

func NewBitcoinMarket() BitcoinMarketSvc {
	return &bitcoinMarketImp{}
}

func (b *bitcoinMarketImp) GetCoins(coin string) map[string]string {
	result := make(map[string]string)
	for key, value := range coins {
		parameter := fmt.Sprintf("^(?i)%s", coin)
		if ok, _ := regexp.MatchString(parameter, value); ok {
			result[key] = value
		}
	}
	return result
}
