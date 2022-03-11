package main

import "disc_bot/service/bot"

func main() {
	// Starts bot instance.
	b := bot.NewDiscordBot()

	// Runs bot service.
	if err := b.Run(); err != nil {
		panic(err)
	}
}
