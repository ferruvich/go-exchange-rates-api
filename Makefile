mocks:
	go generate ./...

test:
	go test ./... -cover

run:
	docker-compose up --build