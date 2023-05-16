// package main

// import (
// 	"os"

// 	"golang-restaurant-management/database"

// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"

// 	"github.com/gin-gonic/gin"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// func main() {
// 	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
// 	// url := os.Getenv("restaurant")
// 	// clientOptions := options.Client().ApplyURI(url)
// 	// client, err := mongo.Connect(context.TODO(), clientOptions)
// 	// port := os.Getenv(clientOptions)

// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "8000"
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)
// 	router.Use(middleware.Authentication())

// 	routes.FoodRoutes(router)
// 	routes.MenuRoutes(router)
// 	routes.TableRoutes(router)
// 	routes.OrderRoutes(router)
// 	routes.OrderItemRoutes(router)
// 	routes.InvoiceRoutes(router)

// 	router.Run(":" + port)
// }

// //////////////////////////////////////////////////////////////////////////////////////////		1
package main

import (
	"fmt"
	"os"

	middleware "golang-restaurant-management/middleware"
	routes "golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	fmt.Println("Server is running...")
	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
	// url := os.Getenv("restaurant")
	// clientOptions := options.Client().ApplyURI(url)
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// port := os.Getenv(clientOptions)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}

////////////////////////////////////////////////////////////////////////////////////////////		2 uji coba

// package main

// import (
// 	"fmt"
// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"
// 	"net/http"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// // func Handler(w http.ResponseWriter, r *http.Request) {
// // 	server := New()

// // 	server.get("/",func(context *Context) {
// // 		context.JSON(200,H{
// // 			"message":"Hello From Vercel !!!"})
// // 	})

// // 	server.get("/users",func(context *Context) {
// // 		name := context.Query("name")
// // 		if name = "" {
// // 			context.JSON(400,H{
// // 				"message":"Hello From Vercel !!!"})
// // 		}else{
// // 			context.JSON(400,H{
// // 				"message":fmt.Sprintf("hai %s", name)})
// // 		}
// // 	})
// // }

// func main(w http.ResponseWriter, r *http.Request) {
// 	server := database()

// 	server.get("/", func(context *Context) {
// 		context.JSON(200, H{
// 			"message": "Hello From Vercel !!!"})
// 	})

// 	fmt.Println("Server is running...")
// 	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/test")
// 	// url := os.Getenv("restaurant")
// 	// clientOptions := options.Client().ApplyURI(url)
// 	// client, err := mongo.Connect(context.TODO(), clientOptions)
// 	// port := os.Getenv(clientOptions)

// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "8000"
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)
// 	router.Use(middleware.Authentication())

// 	routes.FoodRoutes(router)
// 	routes.MenuRoutes(router)
// 	routes.TableRoutes(router)
// 	routes.OrderRoutes(router)
// 	routes.OrderItemRoutes(router)
// 	routes.InvoiceRoutes(router)

// 	router.Run(":" + port)
// }

// //////////////////////////////////////////////////////////////////////////////////////////
// import (
// 	"os"

// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)
// 	router.Use(middleware.Authentication())

// 	routes.FoodRoutes(router)
// 	routes.MenuRoutes(router)
// 	routes.TableRoutes(router)
// 	routes.OrderRoutes(router)
// 	routes.OrderItemRoutes(router)
// 	routes.InvoiceRoutes(router)

// }

// //////////////////////////////////////////////////////////////////////////////////////////
// func main() {
// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "8000"
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)

// 	router.Use(middleware.Authentication())

// 	// API-2
// 	router.GET("/api-1", func(c *gin.Context) {

// 		c.JSON(200, gin.H{"success": "Access granted for api-1"})

// 	})

// 	// API-1
// 	router.GET("/api-2", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"success": "Access granted for api-2"})
// 	})

// 	router.Run(":" + port)
// }

////////////////////////////////////////////////////////////////////////////////////////////

// package main

// import (
// 	"os"

// 	"golang-restaurant-management/database"

// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"

// 	"github.com/gin-gonic/gin"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// func main() {
// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "8000"
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)
// 	router.Use(middleware.Authentication())

// 	routes.FoodRoutes(router)
// 	routes.MenuRoutes(router)
// 	routes.TableRoutes(router)
// 	routes.OrderRoutes(router)
// 	routes.OrderItemRoutes(router)
// 	routes.InvoiceRoutes(router)

// 	router.Run(":" + port)
// }
