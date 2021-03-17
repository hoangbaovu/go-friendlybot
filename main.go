package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoangbaovu/go-friendlybot/internal/commands"
	"github.com/hoangbaovu/go-friendlybot/internal/config"
	"github.com/hoangbaovu/go-friendlybot/internal/events"

	"github.com/bwmarrin/discordgo"
)

func main() {
	const fileName = "./config/config.json"

	cfg, err := config.ParseConfigFromJSONFile(fileName)
	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	s.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildMembers |
			discordgo.IntentsGuildMessages)

	registerEvents(s)
	registerCommands(s, cfg)

	if err = s.Open(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session
	s.Close()
}

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)
	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewMessageHandler().Handler)
}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := commands.NewCommandHandler(cfg.Prefix)
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Sessios.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
