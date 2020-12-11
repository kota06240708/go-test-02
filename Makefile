# コンテナを起動
.PHONY: start
start:
	docker-compose up -d 

# コンテナを起動 (コンソールにデバックを表示させる)
.PHONY: start-build
start-build:
	docker-compose up -d --build

# ログを表示
.PHONY: logs
logs:
	docker-compose logs

# goのログを表示
.PHONY: logs-go
logs-go:
	docker logs -f golang-container

# cronのログを表示
.PHONY: logs-cron
logs-cron:
	docker logs -f go-test-api-01_cron_1

# 開発終了
.PHONY: kill
kill:
	docker-compose kill

# DBデータをdump
.PHONY: dump
dump:
	docker exec -it mysql-container mysqldump -u golang -pgolang -A > dump.sql

# volume毎削除
.PHONY: down
down:
	docker-compose down --volumes

# コンテナの状態を表示
.PHONY: ps
ps:
	docker-compose ps

# 全てのコンテナの状態を表示
.PHONY: ps-all
ps-all:
	docker ps -a

# mysqlのコンテナの中に入る
.PHONY: on-db
on-db:
	docker exec -it mysql-container bin/bash

# nodeのコンテナの中に入る
.PHONY: on-go
on-go:
	docker exec -i -t golang-container sh

# nodeのコンテナの中に入る
.PHONY: on-cron
on-cron:
	docker exec -i -t go-test-api-01_cron_1 sh

# nodeのコンテナの中に入る
.PHONY: on-redis
on-redis:
	docker exec -i -t go-test-api-01_redis_1 sh

# コンテナ、イメージを削除
.PHONY: clean
clean:
	@if [ "$(image)" != "" ] ; then \
			docker rmi -f $(image); \
	fi
	@if [ "$(contener)" != "" ] ; then \
			docker rm -f $(contener); \
	fi

# コンテナをリスタート
.PHONY: restart
restart: kill start
