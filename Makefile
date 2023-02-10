
.DEFAULT_GOAL := help

# go
GO           := go run
GO_SRC_DIR   := ./src

# docker
DOCKER_COMPOSE_DIR                 := ./docker
DOCKER_COMPOSE_FILE                := $(DOCKER_COMPOSE_DIR)/docker-compose.yml

DOCKER_EXEC                        := docker exec -it
DOCKER_COMPOSE_PJ                  := go-simple-echo-server
SERVER_CONTAINER                   := echo-server

# rm
RM := rm -rf

.PHONY: up
up: ## docker環境を立ち上げる
	docker-compose \
	-f $(DOCKER_COMPOSE_FILE) \
	-p $(DOCKER_COMPOSE_PJ) \
	up -d

.PHONY: down
down: ## dockerイメージを削除し、docker環境を閉じる
	docker-compose \
	-f $(DOCKER_COMPOSE_FILE) \
	-p $(DOCKER_COMPOSE_PJ) \
	down \
	--rmi all --volumes --remove-orphans

.PHONY: fclean
fclean:down del-volumes ## マウントしたデータを削除、またdockerイメージも削除する

.PHONY: re
re:fclean up ## 完全に初期化した状態でdocker環境を立ち上げる

.PHONY: del-volumes
del-volumes:

.PHONY: attach-server
attach-server: ## dockerのserverコンテナにアクセスする
	$(DOCKER_EXEC) $(SERVER_CONTAINER) bash

.PHONY: go-lint
go-lint: ## src配下のgoのコードを整形する
	cd $(GO_SRC_DIR) && gofmt -l -w .

.PHONY: go-test
go-test: ## goのテストを実行する
	cd $(GO_SRC_DIR) && go test -v ./...

.PHONY: log
log: ## docker compose環境のログを確認する
	docker-compose \
	-f $(DOCKER_COMPOSE_FILE) \
	-p $(DOCKER_COMPOSE_PJ) \
	logs -f

.PHONY: help
help: ## コマンドの一覧を標示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
