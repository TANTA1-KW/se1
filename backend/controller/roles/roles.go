package roles


import (
   "net/http"
   "github.com/TANTA1-KW/SE67/config"
   "github.com/TANTA1-KW/SE67/entity"
   "github.com/gin-gonic/gin"
)


func GetAll(c *gin.Context) {
   db := config.DB()
   var role []entity.Roles

   db.Find(&role)
   c.JSON(http.StatusOK, &role)
}