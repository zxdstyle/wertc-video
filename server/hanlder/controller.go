package hanlder

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func WsHandler(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}

	user := NewUser(ws)
	msg := &Message{
		Type:    "connect",
		Content: user.UserID,
	}
	user.SendMsg(msg)
	Clients[user.UserID] = user

	defer func() {
		fmt.Println("客户端链接断开")

		fmt.Printf("当前在线人数%d\r", len(Clients))
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}

		HandleData(ws, msg)
	}
}

func HandleData(ws *ghttp.WebSocket, msg []byte) {
	message := &Message{}

	err := json.Unmarshal(msg, message)
	if err != nil {
		fmt.Println("处理数据 json Unmarshal", err)
		return
	}

	switch message.Type {
	// 信令中转
	case "SIGN":
		Sign()
	}
}

func Sign() {

}

func CloseConnection() {
	fmt.Println("客户端链接断开")
}
