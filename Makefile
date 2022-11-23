.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/users/get api/handler/users/get.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/organization/create api/handler/organization/create.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/project/create api/handler/project/create.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
