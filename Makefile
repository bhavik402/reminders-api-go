test-v: 
	go test -v ./...

test: 
	go test ./...

air: 
	air

inttest:
	pushd integration && go run cmd/main/main.go