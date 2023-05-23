package simple

type Database struct {
	Name string
}

type DatabaseMongoDB Database
type DatabasePostgreSQL Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

// Constructor for DatabaseRepository
func NewDatabaseRepository(mongoDB *DatabaseMongoDB, postgresSQL *DatabasePostgreSQL) *DatabaseRepository {
	return &DatabaseRepository{
		DatabaseMongoDB:    mongoDB,
		DatabasePostgreSQL: postgresSQL,
	}
}
