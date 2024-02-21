package controller

import (
	"encoding/json"
	"errors"
	"gpt-project/model"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	File map[string]interface{}
}

//func (s Service) CreateFile(c *gin.Context) {
//	var input model.Input
//
//	err := c.ShouldBindJSON(&input)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, model.Error{Message: err.Error()})
//		return
//	}
//	// TODO call openAI
//
//	jsonArray, _ := json.Marshal(s.File)
//
//	err = ioutil.WriteFile(input.FileName+".json", jsonArray, 0644)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, model.Error{Message: err.Error()})
//		return
//	}
//
//}

// SearchByText : Service to fetch data
func (s Service) SearchByText(c *gin.Context) {
	var res model.BadgesResponse
	var req model.Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: err.Error()})
		return
	}
	if len(req) == 0 {
		c.JSON(http.StatusBadRequest, errors.New("bad-request"))
		return
	}
	// TODO call openAI

	inputSlice := strings.Split(req[0], " ")
	reqString := checkInputSlice(inputSlice)

	if v, ok := s.File[reqString]; ok {
		jsonArray, _ := json.Marshal(v)
		json.Unmarshal(jsonArray, &res)
		c.JSON(http.StatusCreated, res)
	}

}

func checkInputSlice(inputSlice []string) string {
	var res string
	for _, v := range inputSlice {
		if strings.ToLower(v) == "lip" {
			res = "show me some lip essentials for valentine's day"
		} else if strings.ToLower(v) == "beauty" || strings.ToLower(v) == "perfumes" {
			res = "show me some beauty products"
		} else if strings.ToLower(v) == "sleep" || strings.ToLower(v) == "sleepwear" {
			res = "show me some trending regular sleep wear"
		} else if strings.ToLower(v) == "sportswear" || strings.ToLower(v) == "sports" {
			res = "show me some comfortable sportswear for my regular performance activity"
		} else if strings.ToLower(v) == "night" || strings.ToLower(v) == "nightwear" {
			res = "suggest me some night wear for my girls night out"
		}
	}
	return res
}
