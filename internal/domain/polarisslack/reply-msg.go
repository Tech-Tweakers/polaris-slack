package polarisslack

import "polaris-slack/internal/domain/appcontext"

type Replier interface {
	ReplyNewMessage(ctx appcontext.Context, messageReply string) error
}

func (s *slackMessagesAPI) ReplyNewMessage(ctx appcontext.Context, messageReply string) error {
	return nil
}
