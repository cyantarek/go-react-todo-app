gen-mock:
	mockgen -source=internal/domain/hello.go -destination=internal/mocks/hello_mocks.go -package=mocks HelloRepository

build-docker:
	docker build -t tarek5/backend -f build/docker/Dockerfile .

test-hello:
	go test -v backend/internal/hello

test-hello-cover:
	go test -coverprofile=assets/coverage/hello.out -coverpkg=backend/internal/hello backend/internal/hello

build-hello:
	cd cmd/hello && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor -o hello