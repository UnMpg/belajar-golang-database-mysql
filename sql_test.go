package golangdatabasemysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "insert into customer(name) value('fendy')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)

	}

	fmt.Println("sukses memasukkan data")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "select id,name from customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)

	}

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Print("id : ", id)
		fmt.Println("==>  name :", name)
	}
	defer rows.Close()
	fmt.Println("sukses memasukkan data")
}

func TestQuerySqlCOmplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "select id, name, email, balance, rating,birth_date, married, created_at from customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)

	}

	for rows.Next() {
		var id int
		var name string
		var email sql.NullString
		var balance int
		var rating float64
		var married int
		var birthDate, createdAt sql.NullTime

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Print("id : ", id)
		fmt.Println(" ==>  name :", name)

		if email.Valid {
			fmt.Println(" ==>  email :", email.String)
		}
		fmt.Println(" ==>  balace :", balance)
		fmt.Println(" ==>  rating :", rating)

		if birthDate.Valid {
			fmt.Println(" ==>  birth _date :", birthDate.Time)
		}
		fmt.Println(" ==>  create at :", createdAt)
	}
	defer rows.Close()
	fmt.Println("sukses memasukkan data")
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "adminss"

	script := "select username from user where username = '" + username + "' and password='" + password + "'LIMIT 1 "
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)

	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Print("sukses login ", username)
	} else {
		fmt.Println("gagal")
	}
	defer rows.Close()
}

func TestSqlQueryParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "select username from user where username =? and password=? LIMIT 1 "
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)

	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Print("sukses login ", username)
	} else {
		fmt.Println("gagal")
	}
	defer rows.Close()
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "fendy"
	password := "123456"

	script := "insert into user(username,password) value(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)

	}

	fmt.Println("sukses memasukkan data")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "fendy ay"
	password := "123456 aa"

	script := "insert into user(username,password) value(?,?)"
	result, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)

	}

	lastID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("sukses menginset data dengan ID ", lastID)

	fmt.Println("sukses memasukkan data")
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "insert into comment(email,comment) value(?,?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "fendy" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)

		hasil, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastID, _ := hasil.LastInsertId()

		fmt.Println("id terakhir ", lastID)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "insert into comment(email,comment) value(?,?)"
	for i := 0; i < 10; i++ {
		email := "arda" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)

		hasil, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		lastID, _ := hasil.LastInsertId()

		fmt.Println("id terakhir ", lastID)
	}

	// transaksinya

	err = tx.Commit()
	// err = tx.Rollback()

	if err != nil {
		panic(err)
	}
}
