package controller

import (
	"context"
	"fmt"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

var menuCol *mongo.Collection = database.OpenCollection(database.Client, "menu")
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order models.Order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the orders"})
		}
		c.JSON(http.StatusOK, order)
	}
}

// func CreateOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var table models.Table
// 		var order models.Order

// 		if err := c.BindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		validationErr := validate.Struct(order)

// 		if validationErr != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 			return
// 		}

// 		if order.Table_id != nil {
// 			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
// 			defer cancel()
// 			if err != nil {
// 				msg := fmt.Sprintf("message:Table was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 		}

// 		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

// 		order.ID = primitive.NewObjectID()
// 		order.Order_id = order.ID.Hex()

// 		result, insertErr := orderCollection.InsertOne(ctx, order)

// 		if insertErr != nil {
// 			msg := fmt.Sprintf("order item was not created")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		defer cancel()
// 		c.JSON(http.StatusOK, result)
// 	}
// }

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Menambahkan definisi tableCollection
		tableCollection := database.OpenCollection(database.Client, "table")

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": *order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"tabled_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{"menu", order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

// //////////////////////// delete order
func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")

		// Create a filter to match the Order_id
		filter := bson.M{"order_id": orderID}

		// Delete the order
		result, err := orderCollection.DeleteOne(ctx, filter)
		if err != nil {
			log.Println("Error deleting order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "terjadi kesalahan saat menghapus pesanan"})
			cancel()
			return
		}

		// Check the result to determine if the order was deleted
		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "pesanan tidak ditemukan"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "pesanan berhasil dihapus"})
	}
}

func OrderItemOrderCreator(order models.Order) string {

	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.Order_id
}

/////////////////////////////////////////////////////////////////// coba 1

// package controller

// import (
// 	"context"
// 	"fmt"
// 	"golang-restaurant-management/database"
// 	"golang-restaurant-management/models"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

// // var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table") // Tambahkan ini

// func GetOrders() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		result, err := orderCollection.Find(context.TODO(), bson.M{})
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing order items"}) // Ubah pesan error
// 			return
// 		}
// 		var allOrders []bson.M
// 		if err = result.All(ctx, &allOrders); err != nil {
// 			log.Fatal(err)
// 			return
// 		}
// 		c.JSON(http.StatusOK, allOrders)
// 	}
// }

// func GetOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		orderId := c.Param("order_id")
// 		var order models.Order

// 		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the orders"}) // Ubah pesan error
// 			return
// 		}
// 		c.JSON(http.StatusOK, order)
// 	}
// }

// func CreateOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var table models.Table
// 		var order models.Order

// 		if err := c.BindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		validationErr := validate.Struct(order)

// 		if validationErr != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 			return
// 		}

// 		if order.Table_id != " " { // Ubah dari order.Table_id != nil
// 			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
// 			defer cancel()
// 			if err != nil {
// 				msg := fmt.Sprintf("message:Table was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 		}

// 		order.Created_at = time.Now() // Ubah dari parsing string ke time.Now()
// 		order.Updated_at = time.Now() // Ubah dari parsing string ke time.Now()

// 		order.ID = primitive.NewObjectID()
// 		order.Order_id = order.ID.Hex()

// 		result, insertErr := orderCollection.InsertOne(ctx, order)

// 		if insertErr != nil {
// 			msg := fmt.Sprintf("order item was not created")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		defer cancel()
// 		c.JSON(http.StatusOK, result)
// 	}
// }

// func UpdateOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var table models.Table
// 		var order models.Order

// 		var updateObj primitive.D

// 		orderId := c.Param("order_id")
// 		if err := c.BindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if order.Table_id != "" { // Ubah dari order.Table_id != nil
// 			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
// 			defer cancel()
// 			if err != nil {
// 				msg := fmt.Sprintf("message:Table was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 			updateObj = append(updateObj, bson.E{"menu", order.Table_id})
// 		}

// 		order.Updated_at = time.Now() // Ubah dari parsing string ke time.Now()
// 		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

// 		upsert := true

// 		filter := bson.M{"order_id": orderId}
// 		opt := options.UpdateOptions{
// 			Upsert: &upsert,
// 		}

// 		result, err := orderCollection.UpdateOne(
// 			ctx,
// 			filter,
// 			bson.D{
// 				{"$set", updateObj}, // Ubah dari "$st" menjadi "$set"
// 			},
// 			&opt,
// 		)

// 		if err != nil {
// 			msg := fmt.Sprintf("order item update failed")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		defer cancel()
// 		c.JSON(http.StatusOK, result)
// 	}
// }

// func OrderItemOrderCreator(order models.Order) string {

// 	order.Created_at = time.Now() // Ubah dari parsing string ke time.Now()
// 	order.Updated_at = time.Now() // Ubah dari parsing string ke time.Now()
// 	order.ID = primitive.NewObjectID()
// 	order.Order_id = order.ID.Hex()

// 	orderCollection.InsertOne(ctx, order)
// 	defer cancel()

// 	return order.Order_id
// }
