.PHONY: build
build: clean
	go build -o assembler.out main.go

.PHONY: run
run:
	go run main.go

.PHONY: clean
clean:
	go clean
