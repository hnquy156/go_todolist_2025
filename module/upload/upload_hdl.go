package upload

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"todolist/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Upload(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInternal(err))
		}

		dst := fmt.Sprintf("static/%d.%s", time.Now().UnixNano(), fileHeader.Filename)

		if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
			panic(common.ErrInternal(err))
		}

		img := common.Image{
			Id: 0, Url: dst, Width: 100, Height: 100, CloudName: "local", Extension: "",
		}

		img.Fullfill(os.Getenv("APP_HOST"))

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
