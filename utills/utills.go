package utills

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

func Decoposit(input string) []string {
	return strings.Split(input, ":")
}

func NetConnSplitIp(conn net.Conn) string {
	return strings.Split(conn.RemoteAddr().String(), ":")[0]
}

const (
	FMT_RED   = "\033[31m"
	FMT_GREEN = "\033[32m"
	FMT_RESET = "\033[0m"
)

func ColorPrintlnRed(strs ...string) {
	colorPrintln(FMT_RED, strs...)
}

func ColorPrintlnGreen(strs ...string) {
	colorPrintln(FMT_GREEN, strs...)
}

func colorPrintln(color string, strs ...string) {
	fmt.Print(color)
	for _, str := range strs {
		fmt.Print(str)
	}
	fmt.Println(FMT_RESET)
}

// sha 256 암호화
func CryptoSha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashByte := hash.Sum(nil)

	return hex.EncodeToString(hashByte)
}
