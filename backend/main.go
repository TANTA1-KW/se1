package main


import (

   "net/http"


   "github.com/gin-gonic/gin"


   "github.com/TANTA1-KW/SE67/config"
   "github.com/TANTA1-KW/SE67/controller/genders"
   "github.com/TANTA1-KW/SE67/controller/users"


   "github.com/TANTA1-KW/SE67/controller/employee"
   "github.com/TANTA1-KW/SE67/controller/roles"
   "github.com/TANTA1-KW/SE67/controller/status"


   "github.com/TANTA1-KW/SE67/middlewares"

)


const PORT = "8000"


func main() {
   // open connection database
   config.ConnectionDB()
   // Generate databases
   config.SetupDatabase()

   r := gin.Default()
   r.Use(CORSMiddleware())

   // Auth Route
   r.POST("/signup", users.SignUp)
   r.POST("/signin", users.SignIn)

   router := r.Group("/")

   {
       router.Use(middlewares.Authorizes())
       // User Route
       router.PUT("/user/:id", users.Update)
       router.GET("/users", users.GetAll)
       router.GET("/user/:id", users.Get)
       router.DELETE("/user/:id", users.Delete)
       // employee Route
       router.PUT("/employee/:id", employees.Update)
       router.GET("/employees", employees.GetAll)
       router.GET("/employee/:id", employees.Get)
       router.DELETE("/employee/:id", employees.Delete)
   }
   r.GET("/genders", genders.GetAll)
   r.GET("/roles", roles.GetAll)
   r.GET("/status", status.GetAll)

   r.GET("/", func(c *gin.Context) {

       c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)

   })
   // Run the server
   r.Run("localhost:" + PORT)
}


func CORSMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

       c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
       c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
       c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
       c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

       if c.Request.Method == "OPTIONS" {
           c.AbortWithStatus(204)
           return
       }
       c.Next()
   }
}