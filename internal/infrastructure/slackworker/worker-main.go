package slackworker

import (
	"polaris-slack/internal/domain/appcontext"
	polarisslack "polaris-slack/internal/domain/polarisslack"
	"polaris-slack/internal/infrastructure/logger/logwrapper"
)

type Input struct {
	Logger          logwrapper.LoggerWrapper
	SlackMessageAPI polarisslack.UseCases
}

func Start(ctx appcontext.Context, input Input) {
	appctx := appcontext.NewBackground()
	appctx.SetLogger(input.Logger)

	go StartPolling(appctx, input)
}
