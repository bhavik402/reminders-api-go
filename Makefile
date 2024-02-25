test-v: 
	go test -v ./...

test: 
	go test ./...

air: 
	cd api-rest && air

inttest:
	pushd integration && go run cmd/main/main.go