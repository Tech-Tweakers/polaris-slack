package slackmessagesapi

type PayloadConversationHistory struct {
	Messages []struct {
		ClientMsgID string `json:"client_msg_id,omitempty"`
		Type        string `json:"type"`
		Text        string `json:"text"`
		User        string `json:"user"`
		Ts          string `json:"ts"`
		Team        string `json:"team,omitempty"`
		Blocks      []struct {
			Type     string `json:"type"`
			BlockID  string `json:"block_id"`
			Elements []struct {
				Type     string `json:"type"`
				Elements []struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"elements"`
			} `json:"elements"`
		} `json:"blocks,omitempty"`
		ThreadTs        string   `json:"thread_ts,omitempty"`
		ReplyCount      int      `json:"reply_count,omitempty"`
		ReplyUsersCount int      `json:"reply_users_count,omitempty"`
		LatestReply     string   `json:"latest_reply,omitempty"`
		ReplyUsers      []string `json:"reply_users,omitempty"`
		IsLocked        bool     `json:"is_locked,omitempty"`
		Subscribed      bool     `json:"subscribed,omitempty"`
		Subtype         string   `json:"subtype,omitempty"`
		Inviter         string   `json:"inviter,omitempty"`
		Purpose         string   `json:"purpose,omitempty"`
	} `json:"messages"`
	HasMore             bool        `json:"has_more"`
	PinCount            int         `json:"pin_count"`
	ChannelActionsTs    interface{} `json:"channel_actions_ts"`
	ChannelActionsCount int         `json:"channel_actions_count"`
}

type HistoryMessages struct {
	ClientMsgID string `json:"client_msg_id,omitempty"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	User        string `json:"user"`
	Ts          string `json:"ts"`
	Team        string `json:"team,omitempty"`
	Blocks      []struct {
		Type     string `json:"type"`
		BlockID  string `json:"block_id"`
		Elements []struct {
			Type     string `json:"type"`
			Elements []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"elements"`
		} `json:"elements"`
	} `json:"blocks,omitempty"`
	ThreadTs            string      `json:"thread_ts,omitempty"`
	ReplyCount          int         `json:"reply_count,omitempty"`
	ReplyUsersCount     int         `json:"reply_users_count,omitempty"`
	LatestReply         string      `json:"latest_reply,omitempty"`
	ReplyUsers          []string    `json:"reply_users,omitempty"`
	IsLocked            bool        `json:"is_locked,omitempty"`
	Subscribed          bool        `json:"subscribed,omitempty"`
	Subtype             string      `json:"subtype,omitempty"`
	Inviter             string      `json:"inviter,omitempty"`
	Purpose             string      `json:"purpose,omitempty"`
	HasMore             bool        `json:"has_more"`
	PinCount            int         `json:"pin_count"`
	ChannelActionsTs    interface{} `json:"channel_actions_ts"`
	ChannelActionsCount int         `json:"channel_actions_count"`
}

type MessageList struct {
	Messages []Messages
}

type Messages struct {
	ID          int    `json:"id"`
	PayloadTS   string `json:"payloadts"`
	PayloadText string `json:"payloadtext"`
	Replied     bool   `json:"replied"`
}
