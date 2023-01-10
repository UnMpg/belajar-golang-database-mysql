package golangdatabasemysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestXxx(t *testing.T) {
	fmt.Println("test aja")
}

func TestOpenconnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_go")
	if err != nil {
		panic(err)
	}

	// gunakan DB

	defer db.Close()

}
