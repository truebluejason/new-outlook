build:
	dep ensure
	mkdir -p bin
	env GOOS=linux go build -ldflags="-s -w" -o bin/new-outlook .

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
