package app

type (
	MessageForm struct {
		Name    string `json:"name" form:"name"`
		Text    string `json:"text" form:"text"`
		File    string
		IsReply bool
		Board   string
	}
)
