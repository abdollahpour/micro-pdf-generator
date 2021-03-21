compile:
	for i in darwin linux windows ; do \
		GOOS="$${i}" GOARCH=amd64 go build -o bin/mpg-"$${i}"-amd64 cmd/mpg/main.go; \
	done

archive:
	rm -f bin/*.zip
	for i in darwin linux windows ; do \
		zip -j "bin/mpg-$${i}-amd64.zip" "bin/mpg-$${i}-amd64" -x "*.DS_Store"; \
		zip "bin/mpg-$${i}-amd64.zip" -r templates; \
	done

run:
	go run cmd/mpg/main.go

get:
	go get -d -u ./...

image:
	docker build -t $(name) -f docker/Dockerfile . 

test:
	go test -coverprofile=coverage.out -cover ./...

goveralls:
	$GOPATH/bin/goveralls -service=travis-ci -coverprofile=coverage.out

spec:
	go test ./...
