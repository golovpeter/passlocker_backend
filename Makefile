migrate_up:
	migrate -database ${POSTGRESQL_URL} -path db/migration up
migrate_down:
	migrate -database ${POSTGRESQL_URL} -path db/migration down