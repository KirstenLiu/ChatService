package main

type DbManager interface {
	GetUser(uname string) UserId
	SaveUser(string) UserId
	GetUserChatRoom(UserId) []ChatRoomId
	SaveUserChatRoom(string) []ChatRoomId
	GetHistoryMessage(UserId) []ChatRoom
	SaveHistoryMessage(ChatRoomId, string) bool

	GetGobalChatRoom() []ChatRoomId
	SaveGobalChatRoom(string) []ChatRoomId
}

type InMemory struct {
	Users                     map[UserId]User
	GlobalChatRooms           []ChatRoomId
	HistoryMessagePerChatRoom []ChatRoom
}

type InDataBase struct {
}

func (m InMemory) GetUser(uname string) UserId {
	var uid UserId
	uid.Id = 0
	for _, user := range m.Users {
		name := user.Uname
		if name == uname {
			uid = user.Uid
			break
		}
	}
	return uid
}

func (m InMemory) SaveUser(uname string) UserId {
	id := len(m.Users) + 1
	var uid UserId
	uid.Id = id

	var newUser User
	newUser.Uid = uid
	newUser.Uname = uname
	newUser.JoinedChatRoom = nil
	m.Users[uid] = newUser
	return uid
}

func (m InMemory) GetUserChatRoom(uid UserId) []ChatRoomId {
	return m.Users[uid].JoinedChatRoom
}

func (m InMemory) SaveUserChatRoom(string) []ChatRoomId {
	return nil
}

func (m InMemory) GetHistoryMessage(uid UserId) []ChatRoom {
	joinedRoomIDs := m.GetUserChatRoom(uid)
	var joinedChatRooms []ChatRoom

	for _, cid := range joinedRoomIDs {
		joinedChatRooms = append(joinedChatRooms, m.HistoryMessagePerChatRoom[cid.Id])
	}
	return joinedChatRooms
}

func (m InMemory) SaveHistoryMessage(ChatRoomId, string) bool {
	return false
}

func (m InMemory) GetGobalChatRoom() []ChatRoomId {
	return nil
}
func (m InMemory) SaveGobalChatRoom(string) []ChatRoomId {
	return nil
}
