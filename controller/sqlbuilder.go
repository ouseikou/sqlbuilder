package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sql-builder/common"
	"sql-builder/common/clause"
	"sql-builder/service"
)

func GenerateHBQL(c *gin.Context) {
	var req clause.BuilderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buildSql, err := service.BuildSql(req)

	if err != nil {
		resp := common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	result := map[string]string{"sql": buildSql}

	resp := common.Response{
		Code: http.StatusOK,
		Data: result,
	}

	c.JSON(http.StatusOK, resp)
}
