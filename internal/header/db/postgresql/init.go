package postgresql

var DB *PostgreSQL

func init() {
	DB = NewPostgreSQL()
}
