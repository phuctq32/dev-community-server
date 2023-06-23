package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *uploadHandler) UploadSingleImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("image")

		if err != nil {
			panic(common.NewCustomBadRequestError("invalid request"))
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.NewCustomBadRequestError("invalid request"))
		}

		defer func() {
			_ = file.Close()
		}()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(err)
		}

		url, err := hdl.business.UploadImage(c.Request.Context(), dataBytes)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Upload success", url))
	}
}
