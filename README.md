# Location service


It's just a simple implementation of locations service which handles HTTP requests and put them into MQ


## Configuration

Configuration is set using environment variables.

### ADDRESS

The address to bind the server to.

Defaults to `:8000`.

### KAFKA_BROKER

Kafka broker url.

Defaults to `localhost:9092`.

TODO: update to use a list of brokers

### KAFKA_PRODUCER_MAX_RETRY

The total number of times to retry sending a message.

Defaults to `5`.

### KAFKA_TOPIC

The Kafka topic for messages.

Defaults to `location`.