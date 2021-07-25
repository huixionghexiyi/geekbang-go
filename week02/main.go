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

/**
	如果查询的时候返回 ErrNoRow 时，不应该抛出错误，应该内部做降级处理，我这里是将查询不到的信息放到message 中，并返回一个空的 student
 */
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
	db, err := sql.Open("mysql", "root:123456@/geekbang-go")
	if err != nil {
		return nil, xerrors.Wrap(err, "打开数据库失败")
	}

	queryStr := "select id,username,age,gender from student where id = ?"
	stmt, err := db.Prepare(queryStr)
	if err != nil {
		return nil, xerrors.Wrap(err, "stmt 创建失败")
	}
	defer stmt.Close()

	id := 2
	var username string
	var age int
	var gender int

	result := &Result{}

	err = stmt.QueryRow(2).Scan(&id, &username, &age, &gender)
	if xerrors.Is(err, sql.ErrNoRows) {
		result.code = 404
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
