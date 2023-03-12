package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"minibolg/internal/pkg/errno"
)

// ErrResponse 定义了发生错误时的返回消息.
type ErrResponse struct {
	// Code 指定了业务错误码.
	Code string `json:"code"`

	// Message 包含了可以直接对外展示的错误信息.
	Message string `json:"message"`
}

// WriteResponse 将错误或响应数据写入 HTTP 响应主体。
// WriteResponse 使用 errno.Decode 方法，根据错误类型，尝试从 err 中提取业务错误码和错误信息.
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		decode := errno.Decode(err)
		c.JSON(decode.HTTP, ErrResponse{
			Code:    decode.Code,
			Message: decode.Message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
