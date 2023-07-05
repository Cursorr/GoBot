package events

import "github.com/bwmarrin/discordgo"


var invitesMap map[string]map[string]*discordgo.Invite

func init() { 
	invitesMap = make(map[string]map[string]*discordgo.Invite)
}

func updateInvitesMap(s *discordgo.Session, guildID string) {
	invites, _ := s.GuildInvites(guildID)

	if invitesMap[guildID] == nil {
		invitesMap[guildID] = make(map[string]*discordgo.Invite)
	}

	for _, invite := range invites {
		invitesMap[guildID][invite.Code] = invite
	}
}
