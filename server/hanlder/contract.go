package hanlder

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guid"
	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	RoomId   string
	Name     string
	Password string
	Creator  *User
	Viewers  []*User
}

func NewChatRoom(name string, password string, user *User) *ChatRoom {
	return &ChatRoom{
		RoomId:   guid.S(),
		Name:     name,
		Password: password,
		Creator:  user,
	}
}

// 向房主发送消息
func (room *ChatRoom) SendMsgToCreator(message interface{}) {
	msg, _ := gjson.Encode(message)
	room.Creator.Socket.WriteMessage(websocket.TextMessage, msg)
}

// 向所有观众发送消息
func (room *ChatRoom) SendMsgToAllViewer(message interface{}) {
	msg, _ := gjson.Encode(message)
	for _, viewer := range room.Viewers {
		viewer.Socket.WriteMessage(websocket.TextMessage, msg)
	}
}

type User struct {
	UserID string
	Socket *ghttp.WebSocket
}

func NewUser(ws *ghttp.WebSocket) *User {
	return &User{
		Socket: ws,
		UserID: guid.S(),
	}
}

func (user *User) SendMsg(msg interface{}) {
	message, _ := gjson.Encode(msg)

	user.Socket.WriteMessage(websocket.TextMessage, message)
}

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}
