package bitcoin_market

import (
	"disc_bot/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinMarketImp_GetOrderBook(t *testing.T) {
	bmSvc := NewBitcoinMarket("https://www.mercadobitcoin.net/api")

	orderBook, err := bmSvc.GetOrderBook("ETH")
	assert.Nil(t, err)
	assert.NotEqual(t, (model.OrderBook{}), orderBook)
	fmt.Println(orderBook)
}

func TestBitcoinMarketImp_GetCoinTicker(t *testing.T) {
	bmSvc := NewBitcoinMarket("https://www.mercadobitcoin.net/api")

	tickerBTC, err := bmSvc.GetCoinTicker("BTC")
	assert.Nil(t, err)
	assert.NotEqual(t, (model.Ticker{}), tickerBTC)
	fmt.Println(tickerBTC)
}

func TestBitcoinMarketImp_GetCoins(t *testing.T) {
	bmSvc := NewBitcoinMarket("https://www.mercadobitcoin.net/api")

	mapCoins := bmSvc.GetCoins("NFT")

	expectedResult := map[string]string{
		"NFT00":    "Vale do Outback de 100 reais",
		"NFT10":    "Iasy Tata",
		"NFT11":    "NFT Feirante Abaetetubense",
		"NFT12":    "NFT Facas Feitas",
		"NFT13":    "NFT Mandala Yawanawa - Mariri a roda 2",
		"NFT14":    "Dodge Dart Sedan 1970 Verde Imperial",
		"NFT15":    "Dodge Dart Coupe 1971 Vermelho Etrusco",
		"NFT16":    "Dodge Charger LS 1974 White",
		"NFT17":    "Dodge Charger LS 1974 Black",
		"NFT18":    "Combo de Dodges",
		"NFT2":     "NFT Protetores da Floresta",
		"NFT3":     "NFT Protetores da Floresta - Peixe Mandy",
		"NFT4":     "NFT Error",
		"NFT5":     "NFT Simulation",
		"NFT6":     "NFT Cosmovisao Tupinamba da Amazonia",
		"NFT7":     "NFT Barbara Parawara",
		"NFT8":     "NFT Liberdade de sentir",
		"NFT9":     "NFT Pescaria",
		"NFTOKN01": "Cesta de NFTs",
	}
	assert.Equal(t, expectedResult, mapCoins)
}
