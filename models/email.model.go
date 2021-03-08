package models

type Email struct {
	ID      string
	Headers []string

	Sender string
	From   string

	To  string
	Cc  []string
	Bcc []string

	Subject string
	Message string
}
