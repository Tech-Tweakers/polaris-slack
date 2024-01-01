package botengines

import (
	"polaris-slack/internal/infrastructure/environment"
	"polaris-slack/internal/infrastructure/logger"
	"polaris-slack/internal/infrastructure/polaris"

	"github.com/slack-go/slack"
)

func BotReply(PayloadText string, PayloadTS string) {

	logger, dispose := logger.New()
	defer dispose()

	env := environment.GetInstance()
	token := env.SLACK_AUTH_TOKEN
	channel := env.SLACK_CHANNEL_ID
	api := slack.New(token)
	polarisResponse := polaris.ReplyFromSlack(PayloadText, PayloadTS)
	aiReplyStruct := slack.Attachment{
		Color: "green",
		Text:  polarisResponse,
	}
	logger.Info("Sending Message")
	channelID, timestamp, err := api.PostMessage(
		channel,
		// slack.MsgOptionText(PayloadText, false),
		slack.MsgOptionTS(PayloadTS),
		slack.MsgOptionAttachments(aiReplyStruct),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("Message Successfully Sent to" + channelID + " at " + timestamp)
}
