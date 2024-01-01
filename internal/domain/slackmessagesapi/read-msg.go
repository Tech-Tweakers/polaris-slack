package slackmessagesapi

import (
	"slack-messages-api/internal/domain/appcontext"
)

type Reader interface {
	ReadNewMessage(ctx appcontext.Context, channelID string) ([]Messages, error)
}

func (s *slackMessagesAPI) ReadNewMessage(ctx appcontext.Context, channelID string) ([]Messages, error) {
	return nil, nil
}
