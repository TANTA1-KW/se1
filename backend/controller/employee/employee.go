package employees


import (

   "net/http"


   "github.com/gin-gonic/gin"

   "github.com/TANTA1-KW/SE67/config"

   "github.com/TANTA1-KW/SE67/entity"

)


func GetAll(c *gin.Context) {
   var employees []entity.Employee
   db := config.DB()

   results := db.Preload("Gender").Find(&employees)

   if results.Error != nil {
       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
       return
   }
   c.JSON(http.StatusOK, employees)
}


func Get(c *gin.Context) {
   var employee entity.Employee
   ID := c.Param("id")
   db := config.DB()

   results := db.Preload("Gender").First(&employee, ID)

   if results.Error != nil {
       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
       return
   }
   if employee.ID == 0 {
       c.JSON(http.StatusNoContent, gin.H{})
       return
   }
   c.JSON(http.StatusOK, employee)
}


func Update(c *gin.Context) {
   var employee entity.Employee
   EmployeeID := c.Param("id")
   db := config.DB()

   result := db.First(&employee, EmployeeID)

   if result.Error != nil {
       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
       return
   }
   if err := c.ShouldBindJSON(&employee); err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
       return

   }

   result = db.Save(&employee)

   if result.Error != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
       return
   }
   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}


func Delete(c *gin.Context) {
   id := c.Param("id")
   db := config.DB()

   if tx := db.Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
       return
   }
   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})
}