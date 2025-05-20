package ginitem

import (
	"net/http"
	"todolist/common"
	"todolist/module/item/biz"
	"todolist/module/item/model"
	"todolist/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		go func() {
		}()
		// var a []int
		// fmt.Print(a[0])
		var query struct {
			common.Paging
			model.Filter
		}
		if err := c.ShouldBind(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		query.Paging.Process()

		store := storage.NewSQLStore(db)
		business := biz.NewGetItemsBiz(store)

		data, err := business.GetItems(c.Request.Context(), &query.Paging, &query.Filter)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, query.Paging, query.Filter))
	}
}
