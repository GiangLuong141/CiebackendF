package controllers

import (
	"context"
	"example/configs"
	"example/models"
	"example/responses"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var employeeCollection *mongo.Collection = configs.GetCollection(configs.DB, "employees")

func CreateEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var employee models.Employee
		defer cancel()

		if err := c.BindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		result, err := employeeCollection.InsertOne(ctx, employee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func UpdateEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := c.Param("id")

		var employee models.Employee
		if err := c.BindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		result, err := employeeCollection.UpdateOne(
			ctx,
			bson.M{"_id": objectID},
			bson.M{"$set": employee},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{
				"data": result,
			},
		})
	}
}

func GetListEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var employees []models.Employee
		var condition bson.M = bson.M{}
		keyword := c.Query("keyword")
		position := c.Query("position")
		gender := c.Query("gender")
		ethnic := c.Query("ethnic")

		if keyword != "" {
			condition["basic_infor.full_name"] = primitive.Regex{Pattern: keyword, Options: "i"}
		}

		if position != "" {
			condition["recruitment.position"] = position
		}

		if gender != "" {
			condition["basicInfo.gender"] = gender
		}

		if ethnic != "" {
			condition["basicInfo.ethnic"] = ethnic
		}

		defer cancel()
		results, err := employeeCollection.Find(ctx, condition)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.Employee
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			employees = append(employees, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": employees}},
		)
	}
}

func GetDetailEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("id")
		fmt.Print(userId)
		var employee models.Employee
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := employeeCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&employee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": employee}})
	}
}

func DeleteEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := employeeCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)
	}
}
