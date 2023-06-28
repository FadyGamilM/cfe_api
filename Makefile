DSN="host=localhost port=5432 user=fady password=fady dbname=cfedb sslmode=disable timezone=UTC connect_timeout=5"
PORT=8080

# the container for the database
DB_DOCKER_CONTAINER=cfe_db

# when we need to deploy our api to for example to an EC2 instance, we will need to build a binary for example to 
# linux and run this binary, so we will name this binary from here 
BINARY_NAME=cfeapi

# to create a running container instance for postgres using the name of the container we specified above
postgres:
	docker run -d --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=fady -e POSTGRES_PASSWORD=fady postgres:14

# now lets create an actual datbase inside this container 
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=fady --owner=fady cfedb

migrate_up:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\cfe_api\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://fady:fady@127.0.0.1/cfedb?sslmode=disable" up

migrate_down:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\cfe_api\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://fady:fady@127.0.0.1/cfedb?sslmode=disable" down


build:
	@echo " + Building the binary of the backend ... "
	go build -o ${BINARY_NAME} cmd/server/main.go
	@echo " + The binary are ready !"

start_docker:
	@echo " + starting the db docker container"
	docker start ${DB_DOCKER_CONTAINER}

run: build start_docker
	@echo " + starting the api"
	@env PORT=${PORT} DSN=${DSN} ./${BINARY_NAME} &
	@echo " + api started"
	

# i have to run this with unix-based system 
stop:
	@echo " + stopping the running api service"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo " + api service is stopped successfully!"

restart: stop run



# this is how to handle the dirty schema_migrations table >    UPDATE schema_migrations SET dirty = false WHERE version = 1;