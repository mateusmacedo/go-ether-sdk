package message

type Content []byte

type Message interface {
	Content() Content
}

type VoidMessage struct{}

func (m VoidMessage) Content() Content {
	return nil
}