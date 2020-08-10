package domain

type Message struct {
	ID        string `json:"id,required"`
	Text      string `json:"text,omitempty"`
	Submitter string `json:"submitter,omitempty"`
}

// Create an interface to expose the CRUD operations for messages
type MessageStore interface {
	Create(Message) error
	Update(string, Message) error
	Delete(string) error
	GetById(string) (Message, error)
	GetAll() ([]Message, error)
}
