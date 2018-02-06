package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"encoding/json"
)

var db, err = sql.Open("mysql", "root:123qwe@tcp(localhost:3306)/test")

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.Host, r.URL, r.Proto)
	fmt.Println(r.RemoteAddr, r.RequestURI, r.Referer(), r.ContentLength)
	headerList := r.Header
	for key, value := range headerList {
		fmt.Println(key, ":", strings.Join(value, " "))
	}
	fmt.Println(r.Cookies())
	r.ParseForm()
	name := r.FormValue("name")
	age := r.FormValue("age")
	ua := r.UserAgent()
	fmt.Println(name, age, ua)
	fmt.Fprintf(w, "hello world\n")
}

func ListHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	id := r.FormValue("id")
	pid, _ := strconv.Atoi(id)
	result, err := query(pid)
	if err != nil || len(result) == 0{
		fmt.Fprintf(w, "ID不存在")
		return
	}
	res, _ := json.Marshal(result)
	fmt.Fprintf(w, string(res))
}

func CreateHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	name := r.FormValue("name")
	addr := r.FormValue("addr")

	id, err := create(name, addr)
	if err != nil {
		fmt.Fprintf(w, "添加失败")
		return
	}
	fmt.Fprintf(w, "添加成功: " + strconv.Itoa(int(id)))
}

func DeleteHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	id := r.FormValue("id")
	pid, _ := strconv.Atoi(id)

	err := delete(pid)
	if err != nil {
		fmt.Fprintf(w, "删除失败")
		return
	}

	fmt.Fprintf(w, "成功")
}

func UpdateHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	name := r.FormValue("name")
	id := r.FormValue("id")
	pid, _ := strconv.Atoi(id)

	err := update(name, pid)

	if err != nil {
		fmt.Fprintf(w, "编辑失败")
		return
	}

	fmt.Fprintf(w, "成功")
}

func main() {

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/delete", DeleteHandler)
	http.HandleFunc("/update", UpdateHandler)
	http.HandleFunc("/list", ListHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func create(name, addr string) (lastInsertId int64, err error) {

	stmt , err := db.Prepare("insert into publisher(name, addr) values(?, ?)")

	if err != nil {
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, addr)
	if err != nil {
		return
	}

	lastInsertId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func delete(id int) (err error) {
	stmt, err := db.Prepare("delete from publisher where id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return
	}

	affectRows, err := result.RowsAffected()
	if err != nil || affectRows == 0 {
		return
	}
	return
}

func update(name string, id int) (err error) {
	stmt, err := db.Prepare("update publisher set name = ? where id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, id)
	if err != nil {
		return
	}
	affectRows, err := result.RowsAffected()
	if err != nil || affectRows == 0 {
		return
	}
	return
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

func query(id int) (res map[string]string, err error){

	stmt, err := db.Prepare("select * from publisher where id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	res = make(map[string]string)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			continue
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			res[columns[i]] = value
		}
	}

	if err = rows.Err(); err != nil {
		return
	}
	return
}
