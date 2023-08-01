package broker

type Connection interface {
	Send(message []byte, exchangeName, routingKey string) error
	Receive() error
}

type Service interface {
	Send(message []byte) error
	SendErr(message []byte) error
	Receive() error
}
