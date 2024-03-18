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
	if strings.ToLower(inputSlice[0]) == "hi" || strings.ToLower(inputSlice[0]) == "hey" {
		m := model.Msg{
			Msg: "Hello! How can I assist you today?",
		}
		c.JSON(http.StatusOK, m)
		return
	}
	if strings.ToLower(inputSlice[0]) == "hello" {
		m := model.Msg{
			Msg: "Hi there! Welcome to Victoria's Secret! 😊 " +
				"How can I assist you today? Whether you're looking for lingerie, " +
				"swimwear, or beauty products, I'm here to help you find the perfect fit!",
		}
		c.JSON(http.StatusOK, m)
		return

	}

	reqString := checkInputSlice(inputSlice)

	v, ok := s.File[reqString]
	if ok {
		jsonArray, _ := json.Marshal(v)
		json.Unmarshal(jsonArray, &res)
		c.JSON(http.StatusOK, res)
	} else {
		m := model.Msg{
			Msg: "Please modify your query, I couldn't understand!! " +
				"Also visit https://www.victoriassecret.com to continue your shopping experience",
		}
		c.JSON(http.StatusOK, m)
	}
}

func checkInputSlice(inputSlice []string) string {
	var res string
	for _, v := range inputSlice {
		if strings.ToLower(v) == "lip" {
			res = "show me some lip essentials for valentine's day"
		} else if strings.ToLower(v) == "beauty" || strings.ToLower(v) == "perfume" || strings.ToLower(v) == "perfumes" || strings.ToLower(v) == "gift" || strings.ToLower(v) == "wife's" || strings.Contains(strings.ToLower(v), "wife's") ||
			strings.Contains(strings.ToLower(v), "birthday") || strings.Contains(strings.ToLower(v), "gift") {
			res = "show me some beauty products"
		} else if strings.ToLower(v) == "sleep" || strings.ToLower(v) == "sleepwear" || strings.ToLower(v) == "sleepwears" {
			res = "show me some trending regular sleep wear"
		} else if strings.ToLower(v) == "sportswear" || strings.ToLower(v) == "sports" {
			res = "show me some comfortable sportswear for my regular performance activity"
		} else if strings.ToLower(v) == "night" || strings.ToLower(v) == "nightwear" || strings.ToLower(v) == "nightwears" {
			res = "suggest me some night wear for my girls night out"
		} else if strings.ToLower(v) == "looking" || strings.ToLower(v) == "specific" {
			res = "I'm looking for a bra but have specific needs"
		} else if strings.ToLower(v) == "bra" || strings.ToLower(v) == "bras" {
			res = "I'm looking for a bra but have specific needs"
		} else if strings.ToLower(v) == "lingerie" {
			res = "show me some lingerie"
		} else if strings.ToLower(v) == "panties" || strings.ToLower(v) == "panty" {
			res = "show me some panties"
		} else if strings.ToLower(v) == "swim" || strings.ToLower(v) == "beach" || strings.ToLower(v) == "miami" || strings.Contains(strings.ToLower(v), "beach") ||
			strings.Contains(strings.ToLower(v), "trip") {
			res = "show me some swim suits"
		} else if strings.ToLower(v) == "cancer" || strings.ToLower(v) == "survivor" || strings.ToLower(v) == "survivor" {
			res = "cancer survivor"
		} else if strings.ToLower(v) == "check" || strings.ToLower(v) == "checkout" || strings.ToLower(v) == "help" {
			res = "I will check it out, thanks for your help"
		}
	}
	return res
}
