package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func POSTTest(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	println(string(body))
	ctx.String(http.StatusOK, "ok")

}

func main() {
	router := gin.New()
	router.POST("/test", POSTTest)
	router.Run(":5000")
}

/*
curl http://localhost:5000/test -X POST -d "blekota"
*/
