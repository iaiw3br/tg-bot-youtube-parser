package models

type RestResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id ItemInfo `json:"id"`
}

type ItemInfo struct {
	VideoId string `json:"videoId"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

type Response struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type URL struct {
	URLs []string
}
