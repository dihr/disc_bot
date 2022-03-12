package main

import (
	"disc_bot/service/bitcoinMarket"
	"disc_bot/service/bot"
	"disc_bot/service/command"
)

func main() {
	// Starts bitCoin Market service.
	btm := bitcoinMarket.NewBitcoinMarket()

	// Starts command service
	cmd := command.NewCmd(btm)

	// Starts bot service.
	b := bot.NewDiscordBot(cmd)

	// Runs bot service.
	if err := b.Run(); err != nil {
		panic(err)
	}
}
