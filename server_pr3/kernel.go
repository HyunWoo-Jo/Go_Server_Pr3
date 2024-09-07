package server_pr3

var (
	createListener      = make(chan MessageData)
	requestRoomListener = make(chan MessageData)
	cancelListener      = make(chan MessageData)
)

func OnKernel(msgData MessageData) {
	msgData := <-readListener
	msg := Decoposit(msgData.Msg)
	switch msg[1] {
	case "create":
		createListener <- msgData
	case "requestRoom":
		requestRoomListener <- msgData
	case "cancel":
		cancelListener <- <-cancelListener
	}

}
