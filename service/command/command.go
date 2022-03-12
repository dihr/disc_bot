package command

import (
	"disc_bot/service/bitcoin_market"
	"errors"
	"fmt"
	"strings"
)

type (
	CmdSvc interface {
		ListCoins(message string) (string, error)
		ListCoinTicker(message string) (string, error)
	}

	cmdImp struct {
		bmSvc bitcoin_market.BitcoinMarketSvc
	}
)

func NewCmd(bmSvc bitcoin_market.BitcoinMarketSvc) CmdSvc {
	return &cmdImp{
		bmSvc: bmSvc,
	}
}

func (c *cmdImp) ListCoinTicker(message string) (string, error) {
	// Checks if user message contains
	// coin parameter.
	if len(strings.Fields(message)) <= 1 {
		return "", errors.New("invalid command, missing coin param")
	}

	coin := strings.ToUpper(strings.TrimPrefix(message,
		"ticker "))

	coinTicker, err := c.bmSvc.GetCoinTicker(coin)
	if err != nil {
		return "", err
	}

	textResponse := fmt.Sprintf(`
:coin:%s
:chart_with_upwards_trend: Highest trading in the last 24H __***%s***__
:chart_with_downwards_trend: Lowest trading in the last 24H __***%s***__
:1234: Amount traded in the last 24H __***%s***__
:dollar: Unit price of the last trade __***%s***__
:money_mouth: Highest bid price in the last 24 hours __***%s***__
:money_with_wings: lowest bid price in the last 24 hours __***%s***__
			`, coin, coinTicker.High, coinTicker.Low, coinTicker.Vol,
		coinTicker.Last, coinTicker.Buy, coinTicker.Sell)

	return textResponse, nil
}

func (c *cmdImp) ListCoins(message string) (string, error) {
	// Checks if user message contains
	// coin parameter.
	if len(strings.Fields(message)) <= 1 {
		return "", errors.New("invalid command, missing coin param")
	}

	coin := strings.TrimPrefix(message, "list ")

	coins := c.bmSvc.GetCoins(coin)

	if len(coins) == 0 {
		return "", errors.New("invalid coin")
	}

	var textResponse string
	for key, value := range coins {
		textResponse += fmt.Sprintf(":coin: %s   \t:arrow_right: %s\n", key, value)
	}
	return textResponse, nil
}
