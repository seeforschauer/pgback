db:
up:
	docker-compose up -d
	./db/initdb

stop:
down:
	docker-compose down

.PHONY: web
run:
	cd web; go run main.go

critic:
	go get -v github.com/go-lintpack/lintpack/...
	go get -v github.com/go-critic/go-critic/...
	lintpack build -o gocritic -linter.version='v0.3.4' -linter.name='gocritic' github.com/go-critic/go-critic/checkers
	./gocritic check ./...

test:
	go test ./...