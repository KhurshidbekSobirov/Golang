package repo

// Email ...
type Email struct {
	ID             string
	Subject        string
	Body           string
	RecipientEmail string
}

type Sms struct {
	ID        string
	ApiKey    string
	ApiSecret string
	To        string
	Text      string
}

// SendStorageI ...
type SendStorageI interface {
	MakeSent(ID string) error
	Send(subject, body string, status bool, recipients string) error
	SendS(To, Text string) error
}
