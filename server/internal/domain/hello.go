package domain

type HelloMessage struct {
	Text string `json:"text"`
}

type HelloRepository interface {
	Get() HelloMessage
}

