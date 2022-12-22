package costs

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/bou-cost/db"
	co "github.com/tusharhow/bou-cost/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateSteps(c *gin.Context) {
	var steps co.Steps
	c.BindJSON(&steps)
	collection := db.MGI.Db.Collection("steps")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, bson.D{{}}, bson.D{{
		"$set", bson.D{

			{"steps", steps.Steps},
			{"stepsType", steps.StepsType},
			{"stepsDate", steps.StepsDate},
			{"stepsDescription", steps.StepsDescription},
		},
	}})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "Steps updated", "message": "Steps updated successfully"})
}
