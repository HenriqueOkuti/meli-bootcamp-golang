package pingpong

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Are you a Ping or a Pong type of person?
*/
func PingPong(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusTeapot, gin.H{
				"message": "No Ping or Pong for you!",
			})
		}
	}()

	upperCaseProbability := rand.Intn(2)%2 == 0
	var pp IPingPong

	if upperCaseProbability {
		pp = &PingPongPingPong{}
		c.JSON(http.StatusOK, pp.PingPong("ping"))
	} else {
		pp = &PingPongPingPong{}
		c.JSON(http.StatusOK, pp.PingPong("pong"))
	}
}

type PingPongPingPong struct{}

func (pp *PingPongPingPong) PingPong(plinplinplon string) PingPongMessage {
	var message string
	numOfPlinPlinPlon := rand.Intn(10000)

	if numOfPlinPlinPlon > 9000 || numOfPlinPlinPlon < 1000 {
		panic("Too many plin plin plons!")
	}

	for i := 0; i < numOfPlinPlinPlon; i++ {
		message += plinplinplon
	}

	return PingPongMessage{
		Message: message,
	}

}

type PingPongMessage struct {
	Message string `json:"message"`
}

type IPingPong interface {
	PingPong(message string) PingPongMessage
}
