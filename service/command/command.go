package command

import (
	"disc_bot/service/bitcoinMarket"
	"errors"
	"fmt"
	"strings"
)

type (
	CmdSvc interface {
		ListCoins(message string) (string, error)
	}

	cmdImp struct {
		bmSvc bitcoinMarket.BitcoinMarketSvc
	}
)

func NewCmd(bmSvc bitcoinMarket.BitcoinMarketSvc) CmdSvc {
	return &cmdImp{
		bmSvc: bmSvc,
	}
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
