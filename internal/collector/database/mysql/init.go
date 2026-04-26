package mysql

var DB *MySQL

func init() {
	DB = NewMySQL()
}
