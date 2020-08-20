package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", reqBody)

	var reqData LoginRequest
	if err := json.Unmarshal(reqBody, &reqData); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", reqData)
	fmt.Printf("Data: %s\n", reqData.UserName)

	result := Login(reqData)
	if !result.Success {
		fmt.Errorf("login failed on server with user %v", reqData.UserName)
	}
	writer.Header().Set("Content-Type", "application/json")

	/*var userR User
	userR.Uid = result.Uid
	userR.Uname = reqData.UserName
	var chatRoomIds []ChatRoomId
	for _, chatRoomItem := range result.JoinedChatRoom {
		var cid ChatRoomId
		cid = chatRoomItem.Cid
		chatRoomIds = append(chatRoomIds, cid)
	}
	userR.JoinedChatRoom = chatRoomIds*/

	json.NewEncoder(writer).Encode(result)
	fmt.Println("welcome to kiki's chat service")

}

func main() {
	http.HandleFunc("/login", httpHandler)
	http.ListenAndServe(":8080", nil)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
