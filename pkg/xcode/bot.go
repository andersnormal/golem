package xcode

import (
	"github.com/dghubble/sling"
)

// Bot ...
type Bot struct {
	sling *sling.Sling
}

func newBot(sling *sling.Sling) *Bot {
	b := new(Bot)
	b.sling = sling

	return b
}

// CreateBotRequest ...
type CreateBotRequest struct {
	// Name ...
	Name string ``
	// Configuration ...
	Configuration *Configuration
}

// CreateBotResponse ...
type CreateBotResponse struct {
	ID      string `json:"_id"`
	Rev     string `json:"_rev"`
	DocType string `json:"doc_type"`

	*CreateBotRequest
}

// Configuration ...
type Configuration struct {
}

// Error ...
type Error struct {
	Status  int
	Message string
}

// New is adding a new job in chronos
func (b *Bot) Create(req *CreateBotRequest) (*CreateBotResponse, *Error, error) {
	failure := new(Error)
	success := new(CreateBotResponse)

	_, err := b.sling.New().Post(defaultAPIBots).BodyJSON(req).Receive(success, failure)
	if err != nil {
		return nil, nil, err
	}

	return success, failure, nil
}
