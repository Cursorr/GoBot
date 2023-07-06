package events

import (
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)


func findInvite(before map[string]*discordgo.Invite, after []*discordgo.Invite) (*discordgo.Invite){
	for _, invite := range after {
		if invite.Uses > before[invite.Code].Uses {
			return invite
		}
	}
	return nil
}

func OnReady(s *discordgo.Session, event *discordgo.Ready) {
	for _, guild := range s.State.Guilds {
		updateInvitesMap(s, guild.ID)		
	}
}

func OnMemberJoin(s *discordgo.Session, member *discordgo.GuildMemberAdd) {
	guild_id := member.GuildID
	
	if member.User.Bot { return }

	befores_invites := invitesMap[guild_id]
	actual_invites, _ := s.GuildInvites(guild_id)

	invite := findInvite(befores_invites, actual_invites)
	if invite != nil {
		updateUserData(guild_id, invite.Inviter.ID, bson.D{{
			Key: "$inc", Value: bson.D{{
				Key: "invites", Value: 1}}}})
	
		updateUserData(guild_id, member.User.ID, bson.D{{
			Key: "$set", Value: bson.D{{
				Key: "inviter_id", Value: invite.Inviter.ID}}}})
	}

	updateInvitesMap(s, guild_id)
}

func OnMemberRemove(s *discordgo.Session, member *discordgo.GuildMemberRemove) {
	guild_id := member.GuildID

	if member.User.Bot { return }

	data, err := getUserData(guild_id, member.User.ID)

	if err != nil {
		return
	}

	if data.InviterID != "" {
		updateUserData(guild_id, data.InviterID, bson.D{{
			Key: "$inc", Value: bson.D{
				{Key: "invites", Value: -1},
				{Key: "left", Value: 1}}}})
	}
}
