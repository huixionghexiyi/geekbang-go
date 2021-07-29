package handler

import (
	"database/sql"
	"fmt"
	"geekbang-go/webDemo/entity"
)

func GetStudent(db *sql.DB, id int) (*entity.Student, error) {

	stmtOut, err := db.Prepare("select id,username,age,gender from student where id = ?")
	if err != nil {
		print(err.Error())
	}
	defer stmtOut.Close()

	student := &entity.Student{
		Id: id,
	}
	stmtOut.QueryRow(2).Scan(student)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 1 is:%+v", student)
	return student, nil
}
