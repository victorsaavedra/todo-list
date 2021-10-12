.PHONY: docker-build docker-run-air local

docker-build:
	docker compose -f .docker/dev/docker-compose.yaml build

local: docker-build
ifneq ("$(wildcard $(go.mod))", "")
	docker compose -f .docker/dev/docker-compose.yaml run todo-list go mod init github.com/victorsaavedra/todo-list
endif
ifneq ("$(wildcard $(.air.toml))", "")
	docker-compose -f .docker/dev/docker-compose.yaml run --rm todo-list air init
endif

docker-run-air:
	docker compose -f .docker/dev/docker-compose.yaml up --remove-orphans

#.PHONY: start
#start: go-run
#	docker compose -f .docker/dev/docker-compose.yaml up --remove-orphans

#.PHONY: go-run
#go-run:
#	docker compose -f .docker/dev/docker-compose.yaml run --rm todo-list go run main.go

#.PHONE: stop
#stop:
#	docker-compose -f .docker/dev/docker-compose.yaml stop