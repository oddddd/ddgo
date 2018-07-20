package crawler

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"ddgo/core"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}



func PostTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	boot := "123"
	core.WriteOKResponse(w,boot)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
