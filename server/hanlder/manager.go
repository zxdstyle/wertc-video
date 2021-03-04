package hanlder

var (
	Clients   = make(map[string]*User, 100)
	ChatRooms = make([]*ChatRoom, 100)
)
