package utils

import (
	"github.com/Cursorr/gobot/commands"
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
	commands.TestCommand,
}

func RegisterEvents(s *discordgo.Session) {
	for _, event := range eventHandlers {
		s.AddHandler(event)
	}
}
