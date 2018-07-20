package httprouter
import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"ddgo/crawler"
	"ddgo/interceptor"
)

func UrlPath()  {
	router := httprouter.New()
	router.POST("/crawler/postTest/",interceptor.ParameterInterceptor(crawler.PostTest))
	router.POST("/crawler/hello/", interceptor.ParameterInterceptor(crawler.Hello))
	log.Fatal(http.ListenAndServe(":8081", router))
}

