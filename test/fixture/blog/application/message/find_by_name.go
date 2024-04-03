package message

import (
	"encoding/json"

	"github.com/mateusmacedo/go-ether-sdk/application/message"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/model"
)

type FindByNameMessage struct {
	name string
}

func (m FindByNameMessage) Content() message.Content {
	return []byte(m.name)
}

func NewFindByNameMessage(name string) message.Message {
	return &FindByNameMessage{name: name}
}

type FindByNameMessageResult struct {
	ID   int `json:"id"`
	Version int `json:"version"`
	Name string `json:"name"`
}

func (m FindByNameMessageResult) Content() message.Content {
	content, _ := json.Marshal(m)
	return message.Content(content)
}

func NewFindByNameMessageResult(author *model.Author) message.Message {
	return &FindByNameMessageResult{
		ID: author.ID(),
		Version: author.Version(),
		Name: author.Name(),
	}
}