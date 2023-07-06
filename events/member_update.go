package events

import (
	"github.com/Cursorr/gobot/database"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func findInvite(before map[string]*discordgo.Invite, after []*discordgo.Invite) *discordgo.Invite {
	for _, invite := range after {
		if invite.Uses > before[invite.Code].Uses {
			return invite
		}
	}
	return nil
}

func OnReady(s *discordgo.Session, _ *discordgo.Ready) {
	for _, guild := range s.State.Guilds {
		updateInvitesMap(s, guild.ID)
	}
}

func OnMemberJoin(s *discordgo.Session, member *discordgo.GuildMemberAdd) {
	guildId := member.GuildID

	if member.User.Bot {
		return
	}

	beforeInvites := invitesMap[guildId]
	actualInvites, _ := s.GuildInvites(guildId)

	invite := findInvite(beforeInvites, actualInvites)
	if invite != nil {

		database.Instance.UpdateUserData(guildId, invite.Inviter.ID, bson.D{{
			Key: "$inc", Value: bson.D{{
				Key: "invites", Value: 1}},
		}})

		database.Instance.UpdateUserData(guildId, member.User.ID, bson.D{{
			Key: "$set", Value: bson.D{{
				Key: "inviter_id", Value: invite.Inviter.ID}},
		}})
	}

	updateInvitesMap(s, guildId)
}

func OnMemberRemove(_ *discordgo.Session, member *discordgo.GuildMemberRemove) {
	guildId := member.GuildID

	if member.User.Bot {
		return
	}

	data, err := database.Instance.GetUserData(guildId, member.User.ID)

	if err != nil {
		return
	}

	if data.InviterID != "" {
		database.Instance.UpdateUserData(guildId, data.InviterID, bson.D{{
			Key: "$inc", Value: bson.D{
				{Key: "invites", Value: -1},
				{Key: "left", Value: 1}},
		}})
	}
}
