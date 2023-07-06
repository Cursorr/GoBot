package utils

import (
	"github.com/Cursorr/gobot/events"
	"github.com/bwmarrin/discordgo"
)


var eventHandlers = []interface{}{
	events.OnReady,
	events.OnGuildJoin,
	events.OnGuildRemove,
	events.OnInviteCreate,
	events.OnInviteDelete,
	events.OnMemberJoin,
	events.OnMemberRemove,
}

func RegisterEvents(s *discordgo.Session) {
	for _, event := range eventHandlers {
		s.AddHandler(event)
	}
}
