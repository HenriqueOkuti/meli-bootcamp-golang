package extract

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractIDFromContext(c *gin.Context) (int, error) {
	id, exists := c.Params.Get("id")
	if !exists {
		return 0, errors.New("id not found")
	}

	int_id, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("id is not a number")
	}

	return int_id, nil
}
