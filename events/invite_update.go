package events

import "github.com/bwmarrin/discordgo"


func OnInviteCreate(s *discordgo.Session, invite *discordgo.InviteCreate) {
	guildID := invite.GuildID

	if _, ok := invitesMap[guildID]; ok {
		invitesMap[guildID][invite.Code] = invite.Invite
	} else {
		updateInvitesMap(s, guildID)
	}
}

func OnInviteDelete(s *discordgo.Session, invite *discordgo.InviteDelete) {
	guildID := invite.GuildID
	if _, ok := invitesMap[guildID]; ok {
		delete(invitesMap[guildID], invite.Code)
	}
}
