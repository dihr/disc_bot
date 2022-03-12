package main

import (
	"disc_bot/service/bitcoin_market"
	"disc_bot/service/bot"
	"disc_bot/service/command"
)

func main() {
	// Starts bitCoin Market service.
	btm := bitcoin_market.NewBitcoinMarket("https://www.mercadobitcoin.net/api")

	// Starts command service
	cmd := command.NewCmd(btm)

	// Starts bot service.
	b := bot.NewDiscordBot(cmd)

	// Runs bot service.
	if err := b.Run(); err != nil {
		panic(err)
	}
}
