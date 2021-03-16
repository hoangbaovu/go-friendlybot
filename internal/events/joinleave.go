package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type JoinLeaveHandle struct {
}

func NewJoinLeaveHandler() *JoinLeaveHandle {
	return &JoinLeaveHandle{}
}

func (h *JoinLeaveHandle) HandlerJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		fmt.Println("Failed getting guild object: ", err)
		return
	}

	fmt.Printf("Member %s joined the guild %s\n", guild.Name, e.Member.User.String())
}

func (h *JoinLeaveHandle) HandlerLeave(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		fmt.Println("Failed getting guild object: ", err)
		return
	}

	fmt.Printf("Member %s left the guild %s\n", guild.Name, e.Member.User.String())
}
