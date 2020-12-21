package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjrpereira/cnbv/models"
)

func (server *Server) CreateUser(c *gin.Context) {

	//clear previous error if any
	//errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		//errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}
	//user.Prepare()
	//errorMessages := user.Validate("")
	//if len(errorMessages) > 0 {
	//	errList = errorMessages
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"status": http.StatusUnprocessableEntity,
	//		"error":  errList,
	//	})
	//	return
	//}
	userCreated, err := user.CreateUser(server.DB)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		//errList = formattedError
		//errList["Create_user_errror"] = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userCreated,
	})
}
