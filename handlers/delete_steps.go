package costs

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/bou-cost/db"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteSteps(c *gin.Context) {
	idParam := c.Param("id")
	collection := db.MGI.Db.Collection("steps")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.D{{}})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": idParam, "status": "Steps deleted", "message": "Steps deleted successfully"})

}
