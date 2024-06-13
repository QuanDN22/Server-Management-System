package monitor

import (
	"fmt"
	"os/exec"
	"strings"
)

func pingToServer(address string) (bool, error) {
	out, err := exec.Command("ping", address).Output()

	if err != nil {
		return false, err
	}

	if strings.Contains(string(out), "bytes=") {
		fmt.Println(string(out))
		return true, nil
	}
	
	return false, err 
}	
