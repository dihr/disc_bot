package command

import (
	"disc_bot/mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandImp_ListOrderBook(t *testing.T) {
	btcoinApiMock := mock.NewBitcoinMarketMock()
	commandSvc := NewCmd(btcoinApiMock)

	// API CALL NOK
	btcoinApiMock.SetError(true)
	expectedError := errors.New("fail to get orderbook")
	resultNok, err := commandSvc.ListOrderBook("orderbook BTC")
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, "", resultNok)

	// API CALL OK
	expectedResult := ":moneybag: sale offers\n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n:Price: __***10420.000000***__ \tQuantity: __***0.009970***__ \n:Price: __***10488.999990***__ \tQuantity: __***0.466349***__ \n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n:Price: __***10420.000000***__ \tQuantity: __***0.009970***__ \n:Price: __***10488.999990***__ \tQuantity: __***0.466349***__ \n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n:Price: __***10420.000000***__ \tQuantity: __***0.009970***__ \n:Price: __***10488.999990***__ \tQuantity: __***0.466349***__ \n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n\n\n:money_with_wings: shopping offers\n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n:Price: __***10420.000000***__ \tQuantity: __***0.009970***__ \n:Price: __***10488.999990***__ \tQuantity: __***0.466349***__ \n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n:Price: __***10420.000000***__ \tQuantity: __***0.009970***__ \n:Price: __***10488.999990***__ \tQuantity: __***0.466349***__ \n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n:Price: __***10420.000000***__ \tQuantity: __***0.009970***__ \n:Price: __***10488.999990***__ \tQuantity: __***0.466349***__ \n:Price: __***10410.000060***__ \tQuantity: __***2.091900***__ \n"
	btcoinApiMock.SetError(false)
	resultOk, err := commandSvc.ListOrderBook("orderboock BTC")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resultOk)

	// INVALID COMMAND
	expectedError = errors.New("invalid command, missing coin param")
	btcoinApiMock.SetError(false)
	resultInvalidCommand, err := commandSvc.ListOrderBook("BTC")
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedError)
	assert.Equal(t, "", resultInvalidCommand)

}

func TestCommandImp_ListCoinTicker(t *testing.T) {
	btcoinApiMock := mock.NewBitcoinMarketMock()
	commandSvc := NewCmd(btcoinApiMock)

	// API CALL NOK
	btcoinApiMock.SetError(true)
	expectedError := errors.New("fail to get coin ticker")
	resultNok, err := commandSvc.ListCoinTicker("ticker BTC")
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, "", resultNok)

	// API CALL OK
	btcoinApiMock.SetError(false)
	expectedResult := "\n:coin:BTC\n:chart_with_upwards_trend: Highest trading in the last 24H __***14481.47000000***__\n:chart_with_downwards_trend: Lowest trading in the last 24H __***13706.00002000***__\n:1234: Amount traded in the last 24H __***443.73564488***__\n:dollar: Unit price of the last trade __***14447.01000000***__\n:money_mouth: Highest bid price in the last 24 hours __***14447.00100000***__\n:money_with_wings: lowest bid price in the last 24 hours __***14447.01000000***__\n\t\t\t"
	resultOk, err := commandSvc.ListCoinTicker("ticker BTC")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resultOk)

	// INVALID COMMAND
	expectedError = errors.New("invalid command, missing coin param")
	btcoinApiMock.SetError(false)
	resultInvalidCommand, err := commandSvc.ListCoinTicker("BTC")
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedError)
	assert.Equal(t, "", resultInvalidCommand)
}

func TestCommandImp_ListCoins(t *testing.T) {
	btcoinApiMock := mock.NewBitcoinMarketMock()
	commandSvc := NewCmd(btcoinApiMock)

	// INVALID COIN
	btcoinApiMock.SetError(true)
	expectedError := errors.New("invalid coin")
	result, err := commandSvc.ListCoins("list XXX")
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, "", result)

	// COIN LIST OK
	btcoinApiMock.SetError(false)
	expectedResult := ":coin: BTC   \t:arrow_right: Bitcoin\n"
	result, err = commandSvc.ListCoins("list BTC")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}
