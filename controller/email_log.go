package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"send-email/constants"
	"send-email/model/repository"
	"strconv"
)

func GetEmailLogs(c *gin.Context)  {
	serverName, _ := c.Get(constants.ClientDomain)
	emailLogs, err := repository.GetEmailLogsByServer(serverName.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, emailLogs)
}

func GetEmailLogById(c *gin.Context) {
	pk, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"detail": err.Error()})
		c.Abort()
		return
	}
	emailLog, err := repository.GetEmailLogByID(pk, "test")
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"detail": err.Error()})
		c.Abort()
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, emailLog)
}