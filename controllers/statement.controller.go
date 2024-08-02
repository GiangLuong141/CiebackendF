package controllers

import (
	"context"
	"example/configs"
	"example/models"
	"example/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SatementBody struct {
}

var statementCollection *mongo.Collection = configs.GetCollection(configs.DB, "statements")

func CreateStatement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var statement models.Statement
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&statement); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		result, err := statementCollection.InsertOne(ctx, statement)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func UpdateStatement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		var statement models.Statement

		// validate the request body
		if err := c.BindJSON(&statement); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
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

		result, err := statementCollection.UpdateOne(
			ctx,
			bson.M{"_id": objectID},
			bson.M{"$set": statement},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetListStatement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var statements []models.Statement
		var condition bson.M = bson.M{}
		keyword := c.Query("keyword")
		signer := c.Query("signer")
		status := c.Query("status")

		if keyword != "" {
			condition["title"] = primitive.Regex{Pattern: keyword, Options: "i"}
		}

		if signer != "" {
			condition["signer"] = signer
		}

		if status != "" {
			condition["status"] = status
		}

		defer cancel()
		results, err := statementCollection.Find(ctx, condition)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleStatement models.Statement
			if err = results.Decode(&singleStatement); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			statements = append(statements, singleStatement)
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": statements}},
		)
	}
}

func GetListStatementPending() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var statements []models.Statement
		defer cancel()
		results, err := statementCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleStatement models.Statement
			if err = results.Decode(&singleStatement); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			statements = append(statements, singleStatement)
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": statements}},
		)
	}
}

func GetDetailStatement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		statementId := c.Param("id")
		var statement models.Statement
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(statementId)

		err := statementCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&statement)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": statement}})
	}
}

func DeleteStatement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		result, err := statementCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Statement with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Statement successfully deleted!"}},
		)
	}
}
