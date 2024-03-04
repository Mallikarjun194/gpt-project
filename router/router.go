package router

import (
	"encoding/json"
	"fmt"
	"gpt-project/constants"
	"gpt-project/controller"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(corsMiddleware())

	addRoute(r)
	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		c.Next()
	}
}

func addRoute(r *gin.Engine) {
	fileData := readJsonFile()

	s := controller.Service{
		File: fileData,
	}
	r.POST(constants.File, s.SearchByText)

}

func readJsonFile() map[string]interface{} {
	directory, err := os.Getwd() //get the current directory using the built-in function
	if err != nil {
		fmt.Println(err) //print the error if obtained
	}
	fmt.Println("Current working directory:", directory) //print the required directory
	fileContent, err := os.Open(directory + "/internals/sample_mock.json")
	var res map[string]interface{}
	if err != nil {
		fmt.Println(err)
		return res
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	json.Unmarshal(byteResult, &res)

	fmt.Println(res["users"])
	return res
}
