package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	BOT_ID    = os.Getenv("BOT_ID")
	BOT_TOKEN = os.Getenv("BOT_TOKEN")
)

type (
	DiscordBotSvc interface {
		Run() error
	}

	discordBotImp struct {
	}
)

func NewDiscordBot() DiscordBotSvc {
	return &discordBotImp{}
}

func (d *discordBotImp) Run() error {
	// Create a new Discord session using the provided bot token.
	fmt.Print(BOT_TOKEN)
	dg, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}
	defer dg.Close()

	// dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}

	select {}
}
