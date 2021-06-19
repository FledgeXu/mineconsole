build:
	mkdir -p out
	go build -o out ./...
clean:
	rm -rf out cache
