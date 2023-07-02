package email

type Emailer interface {
	Send(email string, data string) error
}
