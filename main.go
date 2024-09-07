package main

import (
	"github.com/HyunWoo-Jo/Go_Server_Pr3/server_pr3"
)

const PORT = ":8555"

func main() {
	server_pr3.OpenServer(PORT)
}
