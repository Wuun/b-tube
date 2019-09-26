package api

import (
	"btube/serializer"
	"btube/service"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

//LikeOperations operate the like operations.
func LikeOperations(c *gin.Context) {
	var srv service.LikeService
	srv.UserID = c.Param("user_id")
	srv.VideoID = c.Param("video_id")
	code := c.Param("operate_code")
	icode, err := codeValid(code)
	if err != nil {
		c.JSON(200, &serializer.Response{
			StatusCode: 4001,
			Msg:        err.Error(),
		})
		return
	}
	srv.Code = icode
	c.JSON(200, srv.Operation())
}

func codeValid(code string) (int, error) {
	icode, err := strconv.Atoi(code)
	if err != nil {
		return -1, err
	}
	switch icode {
	case 0, 1, 2, 3:
		return icode, nil
	default:
		return -1, errors.New("invalid operation code")
	}
}
