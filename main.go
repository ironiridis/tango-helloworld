package main

import (
	"os"

	"github.com/ironiridis/tango/rpc"
)

func main() {
	var err error
	h := rpc.NewHandle(os.Stdout, os.Stdin)

	hellomsg := rpc.UIDefineCategory{Label: "Hello"}
	err = h.Send(&hellomsg, false)
	if err != nil {
		panic(err)
	}

}
