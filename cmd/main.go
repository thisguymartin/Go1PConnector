package main

import (
	pkg "Go1PConnector/pkg"
	"os"

	"github.com/fatih/color"
)

var OP_CONNECT_HOST string
var OP_CONNECT_TOKEN string

func init() {
	OP_CONNECT_HOST = os.Getenv("OP_CONNECT_HOST")
	OP_CONNECT_TOKEN = os.Getenv("OP_CONNECT_TOKEN")

	if OP_CONNECT_HOST != "nil" || OP_CONNECT_TOKEN != "nil" {
		color.Red("OP_CONNECT_HOST or OP_CONNECT_TOKEN is not set")
		os.Exit(0)
	}
}

func main() {
	connection := pkg.OnePasswordConnect(OP_CONNECT_HOST, OP_CONNECT_TOKEN)
	pkg.GetAllVaules(connection)
}
