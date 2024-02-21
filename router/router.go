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

	addRoute(r)
	return r
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
