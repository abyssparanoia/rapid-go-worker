package push

// Message ... プッシュ通知メッセージ
type Message struct {
	Title   string            `json:"title"`
	Body    string            `json:"body"`
	Data    map[string]string `json:"data"`
	IOS     *MessageIOS       `json:"ios"`
	Android *MessageAndroid   `json:"android"`
	Web     *MessageWeb       `json:"web"`
}

// MessageIOS ... プッシュ通知メッセージ(iOS独自部分)
type MessageIOS struct {
	Sound string `json:"sound,omitempty"`
	Badge int    `json:"badge,omitempty"`
}

// MessageAndroid ... プッシュ通知メッセージ(Android独自部分)
type MessageAndroid struct {
	ClickAction string `json:"click_action,omitempty"`
	Sound       string `json:"sound,omitempty"`
	Tag         string `json:"badge,omitempty"`
}

// MessageWeb ... プッシュ通知メッセージ(Web独自部分)
type MessageWeb struct {
	Icon string `json:"icon,omitempty"`
}
