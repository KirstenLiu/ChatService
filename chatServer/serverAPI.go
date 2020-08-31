package main

import (
	"fmt"
	"time"
)

func initDataStore() DbManager {
	var dataStorageInterface DbManager
	var inMem = InMemory{
		Users:                     make(map[UserId]User),
		GlobalChatRooms:           []ChatRoomId{},
		HistoryMessagePerChatRoom: []ChatRoom{},
	}

	dataStorageInterface = inMem

	return dataStorageInterface
}

var globalDataStore = initDataStore()

func getDataStore() DbManager {
	return globalDataStore
}

type User struct {
	Uid            UserId
	Uname          string
	JoinedChatRoom []ChatRoomId
}

type ChatRoom struct {
	Cid             ChatRoomId
	Cname           string
	HistoryMessages []string
}

type UserId struct {
	Id int
}

type ChatRoomId struct {
	Id int
}

type CreateUserRequest struct {
	UserName string
}

type CreateUserResponse struct {
	Uid UserId
}

//CreateUser API create an user with id when user name is provided.
func CreateUser(req CreateUserRequest) CreateUserResponse {
	userName := req.UserName
	var userCreated User
	fmt.Println(userName)

	//TODO: do sth to userName to create user id...
	userCreated.Uid.Id = -1
	var res CreateUserResponse
	res.Uid = userCreated.Uid
	return res
}

type GetChatRoomListRequest struct {
}

type GetChatRoomListResponse struct {
	ChatRoomList []ChatRoomId
}

//GetChatRoomList provide the global list of all chat rooms on the server.
func GetChatRoomList(req GetChatRoomListRequest) GetChatRoomListResponse {
	var chatRoomList []ChatRoomId
	chatRoomList = nil
	var res GetChatRoomListResponse
	res.ChatRoomList = chatRoomList
	return res
}

type JoinChatRoomRequest struct {
	Uid              UserId
	Cid              ChatRoomId
	CreateIfNotExist bool
}

type JoinChatRoomResponse struct {
	Success         bool
	JoinedChatRooms []ChatRoomId
	HistoryMessages []string
}

//JoinChatRoom put specific user into requested chat room,
//provide the history chat in that requested chat room.
func JoinChatRoom(req JoinChatRoomRequest) JoinChatRoomResponse {
	uid := req.Uid
	cid := req.Cid
	createIfNotExist := req.CreateIfNotExist
	fmt.Println(uid, cid, createIfNotExist)

	var res JoinChatRoomResponse
	res.Success = false
	res.JoinedChatRooms = nil
	res.HistoryMessages = nil

	return res
}

type LeaveChatRoomRequest struct {
	Uid UserId
	Cid ChatRoomId
}

type LeaveChatRoomResponse struct {
	Success         bool
	JoinedChatRooms []ChatRoomId
	LeavingTime     time.Time
}

//LeaveChatRoom removes requested user from the specific chat room.
func LeaveChatRoom(req LeaveChatRoomRequest) LeaveChatRoomResponse {
	uid := req.Uid
	cid := req.Cid

	fmt.Println(uid, cid)

	var res LeaveChatRoomResponse
	res.Success = false
	res.LeavingTime = time.Now()
	res.JoinedChatRooms = nil

	return res
}

type SendMessageRequest struct {
	SenderId   UserId
	ReceiverId ChatRoomId
	Message    string
}

type SendMessageResponse struct {
	Success  bool
	SentTime int64
}

//SendMessage send defined message from sender to receiver. Receiver can be either an user or chat room.
func SendMessage(req SendMessageRequest) SendMessageResponse {
	senderId := req.SenderId
	receiverId := req.ReceiverId
	msg := req.Message
	fmt.Println(senderId, receiverId, msg)

	var res SendMessageResponse
	res.Success = false
	res.SentTime = time.Now().Unix()
	return res
}

type LoginRequest struct {
	UserName string
}

type LoginResponse struct {
	Success        bool
	Uid            UserId
	JoinedChatRoom []ChatRoom
}

//Login lets specific user login into the chat service.
func Login(req LoginRequest) LoginResponse {
	userName := req.UserName
	loginTime := time.Now()
	fmt.Println(userName, loginTime)

	dataStore := getDataStore()

	uid := dataStore.GetUser(userName)
	if uid.Id == 0 {
		uid = dataStore.SaveUser(userName)
	}

	//joinedRoomIDs := dataStore.GetUserChatRoom(uid)
	chatRooms := dataStore.GetHistoryMessage(uid)

	var chatRoomInfo []ChatRoom
	for _, chatRoom := range chatRooms {
		var chatRoomItem ChatRoom
		chatRoomItem.Cid = chatRoom.Cid
		chatRoomItem.Cname = chatRoom.Cname
		chatRoomItem.HistoryMessages = chatRoom.HistoryMessages

		chatRoomInfo = append(chatRoomInfo, chatRoomItem)
	}

	var res LoginResponse
	res.Success = true
	res.Uid = uid
	res.JoinedChatRoom = chatRoomInfo

	return res
}

type LogoutRequest struct {
	Uid        UserId
	LogoutTime time.Time
}

type LogoutResponse struct {
	Success bool
}

//Logout lets specific user logout from the chat service.
func Logout(req LogoutRequest) LogoutResponse {
	uid := req.Uid
	logoutTime := req.LogoutTime
	fmt.Println(uid, logoutTime)
	var res LogoutResponse
	res.Success = false
	return res
}
