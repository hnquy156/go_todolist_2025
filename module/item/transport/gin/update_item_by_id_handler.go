package ginitem

import (
	"net/http"
	"strconv"
	"todolist/common"
	"todolist/module/item/biz"
	"todolist/module/item/model"
	"todolist/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
