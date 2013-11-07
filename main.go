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
)


func main() {
	Router.HandleFunc("/{what}", func(writer http.ResponseWriter, req * http.Request) {
			vars := mux.Vars(req)
			Debug(1, req, "hello")
			fmt.Fprint(writer, vars["what"])
		})
	Log.Debug("debug")
	Log.Info("info")
	Log.Warn("warn")
	Log.Alert("alert")
	Log.Error("error")
	err := http.ListenAndServe(Project.String("server", "addr"), Router)
	if	err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}