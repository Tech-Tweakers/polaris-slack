package slackworker

import (
	"slack-messages-api/internal/domain/appcontext"
	slackmessageapi "slack-messages-api/internal/domain/slackmessagesapi"
	"slack-messages-api/internal/infrastructure/logger/logwrapper"
)

type Input struct {
	Logger          logwrapper.LoggerWrapper
	SlackMessageAPI slackmessageapi.UseCases
}

func Start(ctx appcontext.Context, input Input) {
	appctx := appcontext.NewBackground()
	appctx.SetLogger(input.Logger)

	go StartPolling(appctx, input)
}
