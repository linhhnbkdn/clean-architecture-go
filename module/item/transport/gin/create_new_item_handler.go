package ginItem

import (
	"cleanArchitechureGo/module/item/biz"
	"cleanArchitechureGo/module/item/model"
	"cleanArchitechureGo/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateNewItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var itemData model.ToDoItemDTOCreation
		// Parse request
		if err := ctx.ShouldBindJSON(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSQLStorage(db)
		business := biz.NewCreateNewItemBiz(store)

		if err := business.CreateNewItem(ctx.Request.Context(), &itemData); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
