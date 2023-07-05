package events

import "github.com/bwmarrin/discordgo"


func OnGuildJoin(s *discordgo.Session, guild *discordgo.GuildCreate) {
	updateInvitesMap(s, guild.ID)
}

func OnGuildRemove(s *discordgo.Session, guild *discordgo.GuildDelete) {
	if _, ok := invitesMap[guild.ID]; ok {
		delete(invitesMap, guild.ID)
	}
}
