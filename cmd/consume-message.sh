#!/bin/bash

# RabbitMQ connection details
rabbitmq_host="localhost"
rabbitmq_port="15672"  # Use the management port
username="myuser"
password="mypassword"
vhost="%2f"
queue_name="logs-queue-error"
routing_key="logs-error"
exchange_name="logs-exchange"

# Define the URL for rabbitmqadmin
rabbitmqadmin_url="http://${rabbitmq_host}:${rabbitmq_port}/api"

# Function to consume messages
consume_messages() {
    #while true; do
        message=$(curl --user ${username}:${password} --data '{"count":10,"ackmode":"ack_requeue_true","encoding":"application/json","truncate":50000}' -H "content-type:application/json" --request POST ${rabbitmqadmin_url}/queues/${vhost}/${queue_name}/get)
        if [ "$message" != "null" ]; then
            echo "Received message: ${message}"
        fi
    #done
}

# Start consuming messages
consume_messages