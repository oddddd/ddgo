package interceptor

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"ddgo/core"
	"ddgo/core/code"
)


func ParameterInterceptor(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := r.URL.Query()
		getTest := query["test"][0]
		fmt.Println(getTest)

		sign := r.PostFormValue("sign")
		requestTime := r.PostFormValue("requestTime")

		if len(sign) == 0 || len(requestTime) == 0{
			log.Println("ParameterInterceptor : error no sign or requestTime ")

			c,m  := code.Success1()

			core.WriteErrorResponse(w,c,m)
		}

		h(w, r, ps)
	}
}
