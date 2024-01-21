dev:
	air
start:
	GIN_MODE=release go run .
run-test:
	go test -v ./test/
