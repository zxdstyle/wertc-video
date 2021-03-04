package hanlder

import (
	"github.com/gogf/gf/net/ghttp"
)

type CreateRoomRequest struct {
	Name     string `v:"required|length:2,30#请输入账号"`
	Password string
	UserID   string `v:"required#请输入用户ID"`
}

func CreateRoom(r *ghttp.Request) {
	var req *CreateRoomRequest
	if err := r.Parse(&req); err != nil {
		Failed(r, err)
	}

	user, ok := Clients[req.UserID]
	if !ok {
		Failed(r, "客户端未链接")
	}

	room := NewChatRoom(req.Name, req.Password, user)
	ChatRooms = append(ChatRooms, room)

	Success(r, room.RoomId)
}
