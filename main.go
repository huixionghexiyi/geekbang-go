package main

import (
	"encoding/json"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

/**
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

func main() {

	//server := server.NewHttpServer("hello-server")
	//server.Route("/hello", service.Hello)
	//server.Route("/hello", service.GetUser)
	//server.Start("localhost:8080")

}

//http.HandleFunc("/", handler)
//http.HandleFunc("/readBodyOnce", readBodyOnce)
//http.HandleFunc("/getBodyIsNil", getBodyIsNil)
//http.HandleFunc("/queryParams", queryParams)
//http.HandleFunc("/wholeUrl", wholeUrl)
//http.ListenAndServe(":8080", nil)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi welcome,%s!", r.URL.Path[1:])
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "the body is %s \n", string(body))

	body, err = io.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "read the data is %s , the len is %d \n", string(body), len(body))

}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprintf(w, "GetBody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody is not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "quert is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}
