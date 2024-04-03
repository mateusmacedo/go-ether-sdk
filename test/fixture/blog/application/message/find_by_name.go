package message

import (
	"encoding/json"

	"github.com/mateusmacedo/go-ether-sdk/application/message"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/repository"
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

type FindByNameMessageResult []struct {
	ID   int `json:"id"`
	Version int `json:"version"`
	Name string `json:"name"`
}

func (m FindByNameMessageResult) Content() message.Content {
	content, _ := json.Marshal(m)
	return message.Content(content)
}

func NewFindByNameMessageResult(authors repository.FindByNameResult) message.Message {
	result := make(FindByNameMessageResult, 0)
	for _, author := range authors {
		result = append(result, struct {
			ID   int `json:"id"`
			Version int `json:"version"`
			Name string `json:"name"`
		}{
			ID: author.ID(),
			Version: author.Version(),
			Name: author.Name(),
		})
	}
	return result
}