package utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"<PROJECT_NAME>/types"
)

func Response(ctx *gin.Context, code int, msg string, data interface{}) {
	message := types.ErrorTable[code]
	if msg != "" {
		message = fmt.Sprintf("%s, %s", message, msg)
	}
	ctx.JSON(
		http.StatusOK,
		types.Resp{
			Code: code,
			Msg:  message,
			Data: data,
		})
}

// HandleError deal with error in handlers
func HandleError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	logrus.Error(err.Error())
	if err == io.EOF {
		Response(ctx, types.ErrEOF, "", nil)
		return true
	}

	Response(ctx, types.ErrGeneral, err.Error(), nil)
	return true
}

