package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context) {
	// w.Header().Set("Content-Type", "application/json")
	// result := services.GetBooks()
	// json.NewEncoder(w).Encode(result)
	// return
	c.JSON(http.StatusOK, gin.H{"Message": "Get All Books"})
}
