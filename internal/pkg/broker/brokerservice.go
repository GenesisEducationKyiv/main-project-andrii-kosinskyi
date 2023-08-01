package broker

import (
	"fmt"
)

var ErrUnknownName = fmt.Errorf("unknown broker name")

func LoggerServiceFactory(brokerName string, brokerMap map[string]Connection) (Service, error) {
	conn, ok := brokerMap[brokerName]
	if !ok {
		return nil, ErrUnknownName
	}
	switch brokerName {
	case RabbitmqBrokerName:
		return NewRabbitMQService(conn, conn.(*RabbitMQ).LogsExchange, conn.(*RabbitMQ).LogsQueue,
			conn.(*RabbitMQ).RoutingKey)
	case KafkaBrokerName:
		return nil, ErrUnknownName
	default:
		return nil, ErrUnknownName
	}
}
