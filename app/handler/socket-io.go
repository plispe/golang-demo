package handler

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Users struct {
	Username string `json:"username"`
	NumUsers int    `json:"numUsers"`
}

type NumUsers struct {
	NumUsers int `json:"numUsers"`
}

type Username struct {
	Username string `json:"username"`
}

func NewSocketIOHandler() http.Handler {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	users := make(map[string]string)

	numUsers := 0
	server.On("connection", func(so socketio.Socket) {
		so.Join("chat")
		so.On("new message", func(message string) {
			so.BroadcastTo("chat", "new message", Message{users[so.Id()], message})
		})

		so.On("add user", func(nickname string) {
			numUsers++
			users[so.Id()] = nickname
			so.Emit("login", NumUsers{numUsers})
			so.BroadcastTo("chat", "user joined", Users{users[so.Id()], numUsers})
		})

		// server.On("typing", func() {
		// 	so.BroadcastTo("chat", "typing", Username{users[so.Id()]})
		// })
		//
		// server.On("stop typing", func() {
		// 	so.BroadcastTo("chat", "stop typing", Username{users[so.Id()]})
		// })

		server.On("disconnect", func() {
			_, ok := users[so.Id()]
			if ok {
				numUsers = numUsers - 1
				so.BroadcastTo("chat", "user left", Users{users[so.Id()], numUsers})
			}
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}
