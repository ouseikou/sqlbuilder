package controller

import (
	"github.com/forhsd/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ouseikou/sqlbuilder/common"
	"github.com/ouseikou/sqlbuilder/common/clause"
	"github.com/ouseikou/sqlbuilder/common/facade"
	pb "github.com/ouseikou/sqlbuilder/gen/proto"
	"github.com/ouseikou/sqlbuilder/service"
	"google.golang.org/protobuf/encoding/protojson"
)

// GenerateHBQL http json 协议的请求, 代码处理联合类型困难
func GenerateHBQL(c *gin.Context) {
	var req clause.BuilderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buildSql, err := service.BuildSqlByJson(req)

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

// GenerateHBQLByProto protobuf 协议处理联合类型方便
func GenerateHBQLByProto(c *gin.Context) {
	req := &pb.BuilderRequest{}
	bodyBuff, err := c.GetRawData()
	if err != nil {
		errorResponse(c, err)
		return
	}

	err = protojson.Unmarshal(bodyBuff, req)
	if err != nil {
		errorResponse(c, err)
		return
	}

	buildSql, err := service.BuildSqlByProto(req)

	if err != nil {
		errorResponse(c, err)
		return
	}

	result := map[string]string{"sql": buildSql}

	resp := common.Response{
		Code: http.StatusOK,
		Data: result,
	}

	c.JSON(http.StatusOK, resp)
}

func errorResponse(c *gin.Context, err error) {
	resp := common.Response{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
		Data: nil,
	}
	c.JSON(http.StatusInternalServerError, resp)
}

// AnalyzeTemplatesByJson : controller层, 根据http协议解析模板字符串, 提取子模板和主模板
// 参数:
//   - req: 解析请求体
//
// 返回值:
//   - 返回一个 模板上下文对象
//   - 返回一个 异常
func AnalyzeTemplatesByJson(c *gin.Context) {
	var req facade.AnalyzeTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	templateMap, err := service.AnalyzeTemplatesByJson(req)

	if err != nil {
		errorResponse(c, err)
		return
	}

	resp := common.Response{
		Code: http.StatusOK,
		Data: templateMap,
	}

	c.JSON(http.StatusOK, resp)
}

// AnalyzeTemplatesByProto : controller层, 根据proto协议解析模板字符串, 提取子模板和主模板
// 参数:
//   - req: 解析请求体
//
// 返回值:
//   - 返回一个 模板上下文对象
//   - 返回一个 异常
func AnalyzeTemplatesByProto(c *gin.Context) {
	req := &pb.AnalyzeTemplateRequest{}
	bodyBuff, err := c.GetRawData()
	if err != nil {
		errorResponse(c, err)
		return
	}

	err = protojson.Unmarshal(bodyBuff, req)
	if err != nil {
		errorResponse(c, err)
		return
	}

	templateMap, err := service.AnalyzeTemplatesByProto(req)

	if err != nil {
		errorResponse(c, err)
		return
	}

	resp := common.Response{
		Code: http.StatusOK,
		Data: templateMap,
	}

	c.JSON(http.StatusOK, resp)

}

func AnalyzeAdditionByProto(c *gin.Context) {
	req := &pb.AnalyzeAdditionRequest{}
	bodyBuff, err := c.GetRawData()
	if err != nil {
		errorResponse(c, err)
		return
	}

	err = protojson.Unmarshal(bodyBuff, req)
	if err != nil {
		errorResponse(c, err)
		return
	}

	additions, err := service.AnalyzeAdditionByProto(req)

	if err != nil {
		errorResponse(c, err)
		return
	}

	resp := common.Response{
		Code: http.StatusOK,
		Data: additions,
	}

	c.JSON(http.StatusOK, resp)

}

func SetLogger(log *logger.LocalLogger) {
	logger.SetDefaultLogger(log)
}
