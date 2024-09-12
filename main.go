package main

import (
	"Go_Server_Pr3/server_pr3"
	server_pr3_testing "Go_Server_Pr3/testing"
)

const PORT = ":43233"
const TEST = false

func main() {
	if TEST {
		go server_pr3_testing.OnTest_Server(PORT)
	}
	server_pr3.OpenServer(PORT)
}
