package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 业务状态码定义
const (
	CodeSuccess         = 0     // 成功
	CodeBadRequest      = 10001 // 请求参数错误
	CodeUnauthorized    = 10002 // 未授权
	CodeForbidden       = 10003 // 禁止访问
	CodeNotFound        = 10004 // 资源不存在
	CodeServerError     = 10005 // 服务器内部错误
	CodeDBError         = 10006 // 数据库错误
	CodeValidationError = 10007 // 数据验证错误
)

// 分页信息结构
type Pagination struct {
	Page       int `json:"page"`       // 当前页码
	PageSize   int `json:"pageSize"`   // 每页数量
	TotalCount int `json:"totalCount"` // 总记录数
	TotalPages int `json:"totalPages"` // 总页数
}

// 分页数据响应
type PageResponse struct {
	Items      interface{} `json:"items"`      // 数据列表
	Pagination Pagination  `json:"pagination"` // 分页信息
}

// 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    CodeSuccess,
		"message": "操作成功",
		"data":    data,
	})
}

// 带消息的成功响应
func SuccessResponseWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    CodeSuccess,
		"message": message,
		"data":    data,
	})
}

// 分页响应
func PageResponseSuccess(c *gin.Context, items interface{}, pagination Pagination) {
	pageResponse := PageResponse{
		Items:      items,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    CodeSuccess,
		"message": "查询成功",
		"data":    pageResponse,
	})
}

// 错误响应
func ErrorResponse(c *gin.Context, code int, message string, errorDetails ...interface{}) {
	response := gin.H{
		"code":    code,
		"message": message,
	}

	// 添加错误详情
	if len(errorDetails) > 0 {
		if err, ok := errorDetails[0].(error); ok {
			response["error"] = err.Error()
		} else if str, ok := errorDetails[0].(string); ok {
			response["error"] = str
		}

		// 添加字段级错误详情
		if len(errorDetails) > 1 {
			if details, ok := errorDetails[1].(map[string]string); ok {
				var detailList []gin.H
				for field, msg := range details {
					detailList = append(detailList, gin.H{"field": field, "message": msg})
				}
				response["details"] = detailList
			}
		}
	}

	//c.JSON(getStatusCode(code), response)
	//c.Abort()
	c.AbortWithStatusJSON(getStatusCode(code), response)
}

// 验证错误响应
func ValidationErrorResponse(c *gin.Context, errors map[string]string) {
	ErrorResponse(c, CodeValidationError, "数据验证失败", "请检查输入字段", errors)
}

// 根据业务状态码获取HTTP状态码
func getStatusCode(bizCode int) int {
	if bizCode < 10000 {
		return bizCode
	}
	switch bizCode {
	case CodeBadRequest, CodeValidationError:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
