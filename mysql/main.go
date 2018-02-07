package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db, err = sql.Open("mysql", "root:123qwe@tcp(localhost:3306)/test")

func main() {

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	//create()
	//delete()
	//update()
	//updateMany()
	query()
}

func create()  {
	stmt , err := db.Prepare("insert into publisher(name, addr) values(?, ?)")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec("工信出版社", "北京大兴")
	if err != nil {
		fmt.Println(err)
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get last insert id failed: ", err)
		return
	}
	fmt.Println(lastId)
}

func delete()  {
	stmt, err := db.Prepare("delete from publisher where id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	id := 7
	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	affectRows, _ := result.RowsAffected()
	fmt.Println(affectRows)
}

func update()  {
	stmt, err := db.Prepare("update publisher set name = ? where id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	id := 6
	name := "工信出版社"
	result, err := stmt.Exec(name, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	affectRows, _ := result.RowsAffected()
	fmt.Println(affectRows)
}

func updateMany() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	id := 5
	name := "工信出版社"

	result, err := tx.Exec("update publisher set name = ? where id = ?", name, id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	affectRows, err := result.RowsAffected()
	if affectRows == 0 {
		tx.Rollback()
		return
	}

	id = 4
	name = "经济学"
	result, err = tx.Exec("update category set name = ? where id = ?", name, id)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	affectRows, err = result.RowsAffected()
	if affectRows == 0 {
		tx.Rollback()
		return
	}

	tx.Commit()
}

func query()  {

	db.Query("select * from publisher")
	return
	stmt, err := db.Prepare("select * from publisher")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(columns)
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("---------------------------------")
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
}