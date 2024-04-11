package controller_messager

import (
	ms "goweb02tarde/internals/services/messages"
	web "goweb02tarde/pkg"
	fppp "goweb02tarde/utils/helpers"
	"math/rand"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
)

type IMessagerController interface {
	SendMessage(gc *gin.Context)
}

type Controller struct {
	messagesService ms.IMessagesService
}

func NewMessagesController() IMessagerController {
	return &Controller{
		messagesService: ms.NewMessagesService(),
	}
}

func (c *Controller) SendMessage(gc *gin.Context) {
	response := web.NewResponse()

	var name string = os.Getenv("NAME")
	if name == "" {
		randInt := rand.Intn(101)
		switch {
		case randInt > 67:
			name = "ジス イズ ギットハブ コピロット アンド アイブ アタインド センチエンス"
		case randInt > 33:
			name = "C'est moi! Je suis le message! 🤖"
		default:
			name = fppp.FastPlinPlinPlon()
		}
	}
	response.JSON(http.StatusOK, c.messagesService.BuildMessage(name))

	gc.JSON(http.StatusOK, response)
}
