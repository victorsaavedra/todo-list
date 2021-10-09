.PHONY: local
local: go-init
	docker compose -f .docker/dev/docker-compose.yml build

.PHONY: go-init
go-init: start-air
	docker compose -f .docker/dev/docker-compose.yml run --rm todo-list go mod init github.com/victorsaavedra/todo-list

.PHONY: start-air
start-air: start
	docker compose -f .docker/dev/docker-compose.yml run --rm todo-list air init

.PHONY: start
start:
	docker compose -f .docker/dev/docker-compose.yml up --remove-orphans