package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
	"net/http"
)

type WebSocket struct {
	beego.Controller
}

var clients = make(map[*websocket.Conn]bool)
var client = make(map[string]*websocket.Conn)

type SocketCon struct {
	clientId string
	client   *websocket.Conn
	msg      string
}

var cls = make(chan *SocketCon)

// var testMsg = make(chan string)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *WebSocket) Get() {
	go revMsg(ws)
	go sendMsg(cls)
}

func revMsg(ws *WebSocket) {
	web, err := upgrader.Upgrade(ws.Ctx.ResponseWriter, ws.Ctx.Request, nil)
	if err != nil {
		return
	}
	defer func(web *websocket.Conn) {
		err := web.Close()
		if err != nil {
			fmt.Println("socket关闭失败")
			return
		}
	}(web)
	clients[web] = true
	clientId := uuid.New().String()
	client[clientId] = web
	for {
		_, i, err := web.ReadMessage()
		if err != nil {
			log.Debugf("receive message err ==> ", err)
			delete(clients, web)
			break
		}
		cls <- &SocketCon{
			clientId: clientId,
			client:   web,
			msg:      string(i),
		}
	}
}

func sendMsg(ch chan *SocketCon) {
	for {
		select {
		case cls := <-ch:
			cls1 := cls.client
			//fmt.Println("clientId======>", cls.clientId)
			err := cls1.WriteJSON(cls.clientId)
			if err != nil {
				delete(clients, cls1)
				return
			}
		}
	}
}
func sendClient(id string) {
	client1 := client[id]
	err := client1.WriteJSON("aaaaa")
	if err != nil {
		return
	}

}
