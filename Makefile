build:
	go build -o ./bin/blockchain-from-cratch
	
run: build
	./bin/blockchain-from-cratch

test:
	go test ./...