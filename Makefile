DB_PORT	 := 13306
SQS_PORT := 19324

compose/database/up:
	docker compose -f ./tools/local/base.yml up -d

compose/database/down:
	docker compose -f ./tools/local/base.yml down

compose/base/init: run/sqs-local compose/database/init run/sqs-scripts

compose/database/init: compose/database/up sleep docker/mysql/create docker/mysql/migrate
	-mysql -h 127.0.0.1 -P $(DB_PORT) -uroot -proot test < seeds/seed.sql

run/sqs-local:
	docker run -d -p $(SQS_PORT):9324 softwaremill/elasticmq-native:latest

run/sqs-scripts:
	@sh ./scripts/init_sqs.sh http://localhost:$(SQS_PORT)
