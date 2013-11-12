/**
 * User: jackong
 * Date: 10/17/13
 * Time: 6:24 PM
 */
package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"os"
	. "morning-dairy/global"
	"morning-dairy/io"
)


func main() {
	Router.HandleFunc("/{what}/{when}", func(writer http.ResponseWriter, req * http.Request) {
			vars := mux.Vars(req)
			io.Required(req, "param")
			Access.Debug(1, req, "hello")
			io.Puts(writer, "data", vars)
		})
	err := http.ListenAndServe(Project.String("server", "addr"), Router)
	if	err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
