migrate:
	SCYLLA_HOSTS=localhost SCYLLA_KEYSPACE=sessions go run cqlx/migration/main.go