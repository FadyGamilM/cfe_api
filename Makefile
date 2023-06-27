# the container for the database
DB_DOCKER_CONTAINER=cfe_db

# when we need to deploy our api to for example an EC2 instance, we will need to build a binary for example to 
# linux and run this binary, so we will name this binary from here 
BINARY_NAME=cfeapi

# to create a running container instance for postgres using the name of the container we specified above
postgres:
	docker run -d --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=fady -e POSTGRES_PASSWORD=fady postgres:14

# now lets create an actual datbase inside this container 
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=fady --owner=fady cfedb

run:
	go run cmd/server/main.go
