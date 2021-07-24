package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
	"os"
)

type Student struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
}

type Result struct {
	data interface{} `json:"data"`
	code int         `json:"code"`
	msg  string      `json:"msg"`
}

func main() {

	student, err := GetStudent()
	if err != nil {
		fmt.Printf("original error: %T \n %v\n", xerrors.Cause(err), xerrors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("The student is:%v", student)

}

// GetStudent 获取学生
func GetStudent() (*Result, error) {
	db, err := sql.Open("mysql", "root:12345@/geekbang-go")
	if err != nil {
		return nil, xerrors.Wrap(err, "打开数据库失败")
	}

	queryStr := "select id,username,age,gender from student where id = ?"
	stmt, err := db.Prepare(queryStr)
	if err != nil {
		return nil, xerrors.Wrap(err, "stmt 创建失败")
	}
	defer stmt.Close()

	var id int
	var username string
	var age int
	var gender int

	result := &Result{}

	err = stmt.QueryRow(1).Scan(&id, &username, &age, &gender)
	if xerrors.Is(err, sql.ErrNoRows) {
		result.code = 400
		result.msg = fmt.Sprintf("not found student,id:%d", id)
		return result, nil
	}
	student := &Student{
		Id:       id,
		Username: username,
		Age:      age,
		Gender:   gender,
	}
	result.data = student
	result.code = 200

	return result, nil
}
