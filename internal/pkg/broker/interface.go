package broker

type Connection interface {
	Send(message []byte, exchangeName, queueName string) error
	Receive() error
}

type Service interface {
	Send(message []byte) error
	Receive() error
}
