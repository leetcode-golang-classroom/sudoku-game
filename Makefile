.PHONY=build

build-game:
	@CGO_ENABLED=1 GOOS=linux go build -o bin/sudoku-game cmd/main.go

run-game: build-game
	@./bin/sudoku-game


coverage:
	@go test -v -cover ./internal/game/...

test:
	@go test -v ./internal/game/...