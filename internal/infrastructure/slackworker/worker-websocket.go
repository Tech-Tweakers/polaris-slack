package slackworker

import (
	"context"
	"fmt"
	"slack-messages-api/internal/domain/appcontext"
	"slack-messages-api/internal/domain/slackmessagesapi"
	"slack-messages-api/internal/infrastructure/botengines"
	"slack-messages-api/internal/infrastructure/environment"
	"strings"

	"log"
	"os"

	"slack-messages-api/internal/infrastructure/logger"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

var payloadText string
var payloadTS string
var id int = 0
var replied bool = false
var r slackmessagesapi.MessageList
var setup slackmessagesapi.Messages
var toGo []slackmessagesapi.Messages

func StartPolling(appctx appcontext.Context, input Input) {
	logger, dispose := logger.New()
	defer dispose()
	env := environment.GetInstance()
	token := env.SLACK_AUTH_TOKEN
	appToken := env.SLACK_APP_TOKEN
	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r.Messages = make([]slackmessagesapi.Messages, 0)

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		for {
			select {
			case <-ctx.Done():
				logger.Error("Shutting down socketmode listener")
				return
			case event := <-socketClient.Events:
				switch event.Type {
				case socketmode.EventTypeEventsAPI:
					eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
					if !ok {
						logger.Error("Failed to parse EventsAPIEvent")
						continue
					}
					socketClient.Ack(*event.Request)
					payloadText, payloadTS = HandleEventMessage(eventsAPIEvent)
					polarisCheck := strings.Contains(payloadText, "Polaris")
					if polarisCheck {
						botengines.BotReply(payloadText, payloadTS)
					}
				}
			}
		}
	}(ctx, client, socketClient)
	socketClient.Run()
}

func HandleEventMessage(event slackevents.EventsAPIEvent) (string, string) {
	innerEvent := event.InnerEvent
	payloadText = innerEvent.Data.(*slackevents.MessageEvent).Text
	payloadTS = innerEvent.Data.(*slackevents.MessageEvent).TimeStamp
	return payloadText, payloadTS
}

func CheckNewMessages() []slackmessagesapi.Messages {
	logger, dispose := logger.New()
	defer dispose()
	if payloadTS == "" {
		logger.Info("No new messages")
	} else {
		setup = slackmessagesapi.Messages{ID: id, PayloadTS: payloadTS, PayloadText: payloadText, Replied: replied}
		if payloadText == "" {
			toGo = r.Messages
		} else {
			toGo = append(toGo, setup)
			r.Messages = append(r.Messages, setup)
			toGo = r.Messages
			id++
			payloadTS = ""
		}
	}
	return toGo
}

func FlagReplied(ID int) {
	logger, dispose := logger.New()
	defer dispose() // Dispose of the logger

	toGo[ID].Replied = true
	logger.Info("Replied to message with id: " + fmt.Sprint(id))
}
