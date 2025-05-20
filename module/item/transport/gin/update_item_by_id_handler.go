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
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
