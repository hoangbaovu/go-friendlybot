package commands

import "github.com/bwmarrin/discordgo"

// [prefix][invoke/alias] [1st arg] [2nd arg] [3rd arg]
// ;;kick

type Context struct {
	Sessios *discordgo.Session
	Message *discordgo.Message
	Args    []string
	Handler *CommandHandler
}
