package commands

import "github.com/bwmarrin/discordgo"

type MwPermissions struct {
}

func (mw *MwPermissions) Exec(ctx *Context, cmd Command) (next bool, err error) {
	if !cmd.AdminRequired() {
		next = true
		return
	}

	defer func() {
		if !next && err == nil {
			_, err = ctx.Sessios.ChannelMessageSend(ctx.Message.ChannelID,
				"You dont have the permission to execute this command!")
		}
	}()

	guild, err := ctx.Sessios.Guild(ctx.Message.GuildID)
	if err != nil {
		return
	}

	if guild.OwnerID == ctx.Message.Author.ID {
		next = true
		return
	}

	roleMap := make(map[string]*discordgo.Role)

	for _, role := range guild.Roles {
		roleMap[role.ID] = role
	}

	// 1000 = 8
	// 0100 = 4
	// 1100 = 12

	for _, rID := range ctx.Message.Member.Roles {
		if role, ok := roleMap[rID]; ok && role.Permissions&discordgo.PermissionAdministrator > 0 {
			next = true
			break
		}
	}

	return
}
