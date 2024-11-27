package genders


import (

   "net/http"


   "github.com/TANTA1-KW/SE67/config"

   "github.com/TANTA1-KW/SE67/entity"

   "github.com/gin-gonic/gin"

)


func GetAll(c *gin.Context) {


   db := config.DB()


   var genders []entity.Genders

   db.Find(&genders)


   c.JSON(http.StatusOK, &genders)


}