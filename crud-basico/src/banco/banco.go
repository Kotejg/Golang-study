package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver de conexao com mysql
)

// conectar abre a conexao com o banco de dado
func Conectar() (*sql.DB, error) {
	stringConexao := "root:root@/db_golang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return db, nil
	}

	return db, nil
}
