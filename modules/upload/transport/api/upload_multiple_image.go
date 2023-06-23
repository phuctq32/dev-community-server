package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *uploadHandler) UploadMultipleImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["images"]

		var dataBytesArr [][]byte
		for _, f := range files {
			file, err := f.Open()
			if err != nil {
				panic(common.NewCustomBadRequestError("invalid request"))
			}

			defer func() {
				_ = file.Close()
			}()

			dataBytes := make([]byte, f.Size)
			if _, err := file.Read(dataBytes); err != nil {
				panic(err)
			}

			dataBytesArr = append(dataBytesArr, dataBytes)
		}

		urls, err := hdl.business.UploadImages(c.Request.Context(), dataBytesArr)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Upload successfully", urls))
	}
}
