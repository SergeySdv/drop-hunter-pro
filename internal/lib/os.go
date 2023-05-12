package lib

import (
	"fmt"
	"os"
	"syscall"
)

func OSSignalsHandler(signal os.Signal) {
	if signal == syscall.SIGTERM {
		fmt.Println(" user interrupted script execution")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		fmt.Println(" user interrupted script execution")
		os.Exit(0)
	}
}
