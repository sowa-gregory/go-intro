package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id       *uint32 `json:"book_id" binding:"required"`
	BookName *string `json:"book_name" binding:"required"`
}

func POSTTest(ctx *gin.Context) {
	var inputData Book
	if err := ctx.BindJSON(&inputData); err != nil {
		ctx.String(http.StatusBadRequest, "bzdury")
		return
	}
	userName := ctx.Param("username")
	fmt.Println(userName, *inputData.Id, *inputData.BookName)
	ctx.String(http.StatusOK, "ok")
}

func main() {
	router := gin.New()
	router.POST("/buyBook/:username", POSTTest)
	router.Run(":5000")
}

/*
curl http://localhost:5000/buyBook/grzesiek -X POST -d '{"book_id":345, "book_name":"cos"}'
*/
