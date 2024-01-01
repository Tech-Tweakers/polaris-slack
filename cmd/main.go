package main

import (
	"slack-messages-api/internal/domain/appcontext"
	"slack-messages-api/internal/domain/slackmessagesapi"
	"slack-messages-api/internal/infrastructure/api"
	"slack-messages-api/internal/infrastructure/environment"
	"slack-messages-api/internal/infrastructure/logger"
	"slack-messages-api/internal/infrastructure/logger/logwrapper"
	"slack-messages-api/internal/infrastructure/slackworker"

	"go.uber.org/zap"
)

func main() {

	ctx := appcontext.NewBackground()

	env := environment.GetInstance()
	zaplogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zaplogger}).SetVersion("1.0")
	logger.Info("Starting Slack Messages API")

	logger.Info("Initializing Slack Messages API",
		zap.String("Log Level:", env.LOG_LEVEL),
		zap.String("Environment:", env.ENVIRONMENT),
		zap.String("Slack API URL:", env.SLACK_API_URL),
		zap.String("Slack APP Token:", env.SLACK_APP_TOKEN),
		zap.String("Slack Channel ID:", env.SLACK_CHANNEL_ID),
		zap.String("Slack Auth Token:", env.SLACK_AUTH_TOKEN),
	)

	slackMessagesAPIUseCases, err := setupSlackMessagesAPI(ctx)

	if err != nil {
		logger.Error("failed to configure slack messages api", zap.Error(err))
	}

	setupWorker(ctx, logger, slackMessagesAPIUseCases)

	logger.Info("Starting API")
	api.MakeDefaultRoutes(ctx)

}

func setupSlackMessagesAPI(ctx appcontext.Context) (slackMessagesAPIUseCases slackmessagesapi.UseCases, err error) {
	slackMessageApiInput := &slackmessagesapi.Input{}
	slackMessagesAPIUseCases = slackmessagesapi.New(ctx, slackMessageApiInput)
	return slackMessagesAPIUseCases, err
}

func setupWorker(ctx appcontext.Context, logger logwrapper.LoggerWrapper, slackMessagesAPIUseCases slackmessagesapi.UseCases) {
	logger.Info("Configuring Worker")
	input := slackworker.Input{
		Logger:          logger,
		SlackMessageAPI: slackMessagesAPIUseCases,
	}
	slackworker.Start(ctx, input)
}
