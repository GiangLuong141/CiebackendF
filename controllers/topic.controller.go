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

var topicCollection *mongo.Collection = configs.GetCollection(configs.DB, "topics")

func CreateTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var topic models.Topic
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&topic); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		result, err := topicCollection.InsertOne(ctx, topic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func UpdateTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := c.Param("id")

		var topic models.Topic
		if err := c.BindJSON(&topic); err != nil {
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

		result, err := topicCollection.UpdateOne(
			ctx,
			bson.M{"_id": objectID},
			bson.M{"$set": topic},
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

func GetListTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var topics []models.Topic
		var condition bson.M = bson.M{}
		keyword := c.Query("keyword")
		typeA := c.Query("type")
		managementLevel := c.Query("managementLevel")
		if keyword != "" {
			condition["title"] = primitive.Regex{Pattern: keyword, Options: "i"}
		}
		if typeA != "" {
			condition["type"] = typeA
		}

		if managementLevel != "" {
			condition["managementLevel"] = managementLevel
		}

		defer cancel()
		results, err := topicCollection.Find(ctx, condition)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTopic models.Topic
			if err = results.Decode(&singleTopic); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			topics = append(topics, singleTopic)
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": topics}},
		)
	}
}

func GetListTopicPending() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var topics []models.Topic
		var condition bson.M = bson.M{}
		keyword := c.Query("keyword")
		typeA := c.Query("type")
		managementLevel := c.Query("managementLevel")
		if keyword != "" {
			condition["title"] = primitive.Regex{Pattern: keyword, Options: "i"}
		}
		if typeA != "" {
			condition["type"] = typeA
		}

		if managementLevel != "" {
			condition["managementLevel"] = managementLevel
		}

		defer cancel()
		results, err := topicCollection.Find(ctx, condition)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTopic models.Topic
			if err = results.Decode(&singleTopic); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			topics = append(topics, singleTopic)
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": topics}},
		)
	}
}

func GetDetailTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		topicId := c.Param("id")
		var topic models.Topic
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(topicId)

		err := topicCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&topic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": topic}})
	}
}

func DeleteTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		result, err := topicCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Topic with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Topic successfully deleted!"}},
		)
	}
}
