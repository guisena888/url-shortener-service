build:
	go build -o bin/app
run: build
	docker-compose up -d
	./bin/app

test:
	go test -v ./... -count=1