# User management service

## Instructions

Assuming you are on the project directory and Docker is installed.
Also assuming that all of the information sent defined in the protobuf.

### Docker

For Apple M1 users :
`sudo docker buildx build -t user-management-service:latest --platform=linux/arm64 . --load`
For **amd64** we build the container image for our microservice
`DOCKER_BUILDKIT=0 docker build -t user-management-service:latest .`.
After docker building process is completed succesfully,
`docker compose up`
optionally (if you want to run it as background daemon).
`docker compose up -d` .
Then **user-management-service** should be ready to make CRUD operations on it.

### Local

If `export MONGO_URL=<mongo_url:mongo_port>`set, program will use it. Otherwise `mongodb://127.0.0.1:27017` is the default url.
Additionally **MUST** :
`export KAFKA_URL=<kafka_url>:<kafka_port>`
To run the application locally first make sure that mongo instance is running on the system.
Use `go build cmd/main.go -o user-management-service`.
Finally application is ready to be used with `./user-management-service`.

### Events

After `docker compose` is up and running. Service will be alive to respond client requests. When user send an operation request:
Events can be seen,

- Run `docker exec -it broker /bin/bash` to fall into bash of Kafka image.
- Change directory to `cd /bin`. Here we have the builtin tools to monitor messages provided by Kafka.
- Run to `kafka-console-consumer --bootstrap-server localhost:9092 --topic user --from-beginning`. To see which events have been published.

# Choises

I tried to keep OSS dependency low. Better frameworks for log management, parsing exists etc.. according to my findings.

### Why Hexagonal Architecture?

By implementing hexagonal architecture basic API functionality could easily be divided into _adapter_,_application_,_core_ layers. By doing so we could abstract each layer from another using interfaces like contracts. Since protobuf messages structures should not be sent directly to the database _model_ approach has been used to transform and manipulate data and vice a versa.

### Why MongoDB?

First of all great documentation. Since this is the first time I used MongoDB, documentation has huge effect on database selection. It is also NoSQL which makes it easy to generate data and play with it. However I had problem about keeping `created_at` field immutable. Could relational db would be better, I will think about it.

### Why Kafka?

Kafka is used to produce events and let other interested servers notified about the changes. Since kafka is well documented, easily built up I preferred Kafka to uses broker.

## Logging

Logs do not include user information because of the security concern.

## Improvements

### Security

Authorization could be implemented for secure communication.

### Scaling

From vertical scaling perspective increasing the machine specs would help which application is running. From the horizontal perspective a task queue which filled by client request and workers listening on that queue would help us on making concurrent jobs in the application.

### Deployment

An automated CI/CD pipeline to run tests and deployment could be added. After deployment monitoring tools could be added to collect, analyze and debug services.

## Testing

It would be nice if I had catch up to write integration tests.

## Screenshot

Server
![plot](resources/Screenshot%202023-03-22%20at%2009.18.25.png)
