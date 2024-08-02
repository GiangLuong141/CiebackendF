package controllers

import (
	"context"
	"example/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func StatisticDataInYear() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// Lấy năm hiện tại
		currentYear := time.Now().Year()
		startDate := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(currentYear, time.December, 31, 23, 59, 59, 0, time.UTC)

		countEmployees, err := employeeCollection.CountDocuments(ctx, bson.M{
			"basic_infor.issuance_date": bson.M{
				"$gte": startDate,
				"$lte": endDate,
			},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		countTopics, err := topicCollection.CountDocuments(ctx, bson.M{
			"created_at": bson.M{
				"$gte": startDate,
				"$lte": endDate,
			},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		countStatements, err := statementCollection.CountDocuments(ctx, bson.M{
			"created_at": bson.M{
				"$gte": startDate,
				"$lte": endDate,
			},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{
			"countTopics":     countTopics,
			"countStatements": countStatements,
			"countEmployees":  countEmployees,
		}})
	}
}
