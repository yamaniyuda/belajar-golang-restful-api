package simple

type Database struct {
	Name string
}

type DatabasePostgresSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgresSQL() *DatabasePostgresSQL {
	return (*DatabasePostgresSQL)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgresSQL *DatabasePostgresSQL
	DatabaseMongoDB     *DatabaseMongoDB
}

func NewDatabaseRepository(databasePostgresSQL *DatabasePostgresSQL, databaseMongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresSQL: databasePostgresSQL, DatabaseMongoDB: databaseMongoDB}
}
