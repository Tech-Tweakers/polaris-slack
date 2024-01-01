package slackmessagesapi

import "slack-messages-api/internal/domain/appcontext"

type UseCases interface {
	Reader
	Replier
}

type Input struct{}

type slackMessagesAPI struct{}

func New(ctx appcontext.Context, input *Input) UseCases {
	return &slackMessagesAPI{}
}
