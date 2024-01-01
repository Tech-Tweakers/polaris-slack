package slackclient

import (
	"polaris-slack/internal/infrastructure/environment"
	"polaris-slack/internal/infrastructure/logger"
	"polaris-slack/internal/infrastructure/slackworker"

	"github.com/slack-go/slack"
)

func ReplyMessage(PayloadTS string, PayloadText string, Replied bool, ID int) {

	logger, dispose := logger.New()
	defer dispose()

	env := environment.GetInstance()
	token := env.SLACK_AUTH_TOKEN
	channel := env.SLACK_CHANNEL_ID
	api := slack.New(token)
	attachment := slack.Attachment{
		Color: "green",
		Text:  PayloadText,
	}

	logger.Info("Sending Message")
	channelID, timestamp, err := api.PostMessage(
		channel,
		// slack.MsgOptionText(PayloadText, false),
		slack.MsgOptionTS(PayloadTS),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	slackworker.FlagReplied(ID)
	logger.Info("Message Successfully Sent to" + channelID + " at " + timestamp)
}
