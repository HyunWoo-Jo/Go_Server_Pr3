# Go_Server_Pr3
---
Room Join 서버

---
### Main Logic Flow
<img width="881" alt="MainLogic" src="https://github.com/user-attachments/assets/8cc34674-7a98-419e-bd20-e461437b1849">

---

#### server 구동
```go
/// server.go
// 서버 Open / tcp 연결
func OpenServer(port string) {

	cm = NewConnManager() // conn manager 생성
	rm = NewRoomManager() // room manager 생성

	lnsten, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Server err :", err)
		return
	}
	defer lnsten.Close()

	go OnAccept(lnsten)

	// Server Off 처리
	s := ""
	for {
		fmt.Println("Server Off command \"server -off\"")
		fmt.Scanln(&s)
		if s == "server -off" {
			break
		}
	}
	fmt.Println("close")
}
```

---

#### 연결 처리
```go
/// server.go
// 연결 수락
func OnAccept(ln net.Listener) {
	fmt.Println("On Accept")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept err :", err)
			continue
		}
		fmt.Println(conn.RemoteAddr().String())
		cm.AddConn(conn.RemoteAddr().String(), conn) // conn Manager에 등록

		go OnReceiveMessage(conn)
	}
}
/// message.go
// 수신
func OnReceiveMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil { // 에러
			fmt.Println("Read Err", err)
			cancel := MessageData{"err:cancel", conn}
			go OnKernel(cancel)
			return
		}
		utills.TrimNewline(&message)
		fmt.Printf("Read : %s Size: %d\n", message, len(message))
		go OnKernel(MessageData{message, conn})
	}
}

```

---

#### 분배
```go
/// kernel.go
func OnKernel(msgData MessageData) {
	msgs := utills.Decoposit(msgData.Msg)
	switch msgs[1] {
	case "createRoom":
		CreateRoom(msgData, msgs)
	case "requestRoom":
		RequestRoom(msgData)
	case "joinRoom":
		joinRoom(msgData, msgs)
	case "cancel":
		Cancel(msgData)
	}
}
```

---


