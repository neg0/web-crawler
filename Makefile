.PHONY: help     # Generate list of targets with descriptions
help:
	@echo "\n"
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20


.PHONY: up       # Creates container for PHP
up:
	docker-compose -f build/docker/docker-compose.yml up -d

.PHONY: down     # It shuts down the running PHP container
down:
	docker-compose -f build/docker/docker-compose.yml down

.PHONY: build # compiles the Golang to executable file and installs it to bin
build:
	docker-compose -f build/docker/docker-compose.yml exec golang sh -c "go build cmd/crawler.go && mv /var/apps/cuvva/crawler /usr/bin"

.PHONY: run # Runs the app inside the Golang container  e.g. make run url=https://cuvva.com
run:
	@docker-compose -f build/docker/docker-compose.yml exec golang sh -c "crawler $(url)"

.PHONY: ssh      # Enters the Golang Container
ssh:
	docker-compose -f build/docker/docker-compose.yml exec golang sh

.PHONY: test # Runs the tests inside the Golang container
test:
	docker-compose -f build/docker/docker-compose.yml exec golang sh -c "go test ./pkg/... -coverprofile cover.out"

.PHONY: lint # Runs pipeline lints inside the Golang Container
lint:
	docker-compose -f build/docker/docker-compose.yml exec golang sh -c 'gofmt -w "./pkg/.." && gofmt -w "./internal/.." && golangci-lint run \
                                                                          --enable=golint \
                                                                          --enable=gofmt \
                                                                          --enable=gosec  \
                                                                          --enable=misspell \
                                                                          --enable=maligned \
                                                                          --enable=interfacer \
                                                                          --enable=stylecheck  \
                                                                          --enable=unconvert \
                                                                          --enable=maligned \
                                                                          --enable=goconst \
                                                                          --enable=whitespace \
                                                                          --enable=unparam \
                                                                          --enable=scopelint \
                                                                          --enable=rowserrcheck \
                                                                          --enable=lll \
                                                                          --enable=dupl \
                                                                          --enable=funlen \
                                                                          --enable=dogsled \
                                                                          --enable=depguard \
                                                                          --enable=bodyclose \
                                                                          --enable=nakedret'
