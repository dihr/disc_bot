package bot

import (
	"disc_bot/service/command"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BOT_ID            = os.Getenv("BOT_ID")
	BOT_TOKEN         = os.Getenv("BOT_TOKEN")
	DIRECT_CHANNEL_ID = os.Getenv("DIRECT_CHANNEL_ID")
)

type (
	DiscordBotSvc interface {
		Run() error
	}

	discordBotImp struct {
		cmd       command.CmdSvc
		functions map[string]func(string) (string, error)
	}
)

func NewDiscordBot(cmd command.CmdSvc) DiscordBotSvc {
	return &discordBotImp{
		cmd:       cmd,
		functions: map[string]func(string) (string, error){},
	}
}

func (d *discordBotImp) Run() error {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}
	defer dg.Close()

	// Set valid functions
	d.setFunctions()

	// Set handler function
	dg.AddHandler(d.Handler)

	// Open a websocket connection to Discord and begin listening.
	if err := dg.Open(); err != nil {
		return err
	}

	// Blocks forever with.
	select {}
}

func (d *discordBotImp) setFunctions() {
	// map with regex that triggers the command.
	d.functions["(?i)list"] = d.cmd.ListCoins
}

func (d *discordBotImp) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// removes bot prefix from message
	m.Content = strings.TrimPrefix(m.Content,
		fmt.Sprintf("<@!%s> ", BOT_ID))

	// Checks if bot has been mentioned or
	// has received direct message.
	if !isValidMessage(m.ChannelID, m.Mentions, m.Author.ID) {
		return
	}

	for key, value := range d.functions {
		go func(parameter string, fn func(string) (string, error)) {
			if ok, _ := regexp.MatchString(parameter, m.Content); ok {
				response, err := fn(m.Content)
				if err != nil {
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
				}
				_, _ = s.ChannelMessageSend(m.ChannelID, response)
			}
		}(key, value)
	}
}

func isValidMessage(channelID string, mentions []*discordgo.User, authorID string) bool {
	if authorID == BOT_ID {
		return false
	}

	if channelID == DIRECT_CHANNEL_ID {
		return true
	}

	if len(mentions) == 1 && mentions[0].ID == BOT_ID {
		return true
	}
	return false
}
