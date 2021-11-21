package testSuites

import (
	"fmt"
	"os/exec"
	"strings"
)

func VerifyOSVersion(osVersion string) bool {
	var resultBool bool
	cmdOut, err := exec.Command("cat", "/etc/issue").Output()
	fmt.Printf("OS version: %s\n", cmdOut)
	if err != nil {
		fmt.Printf("%s", err)
	}
	cmdOutStr := string(cmdOut[:])
	if cmdOutStrContainsOSVer := strings.Contains(cmdOutStr, osVersion); cmdOutStrContainsOSVer {
		resultBool = true
	} else {
		resultBool = false
	}
	return resultBool
}

func VerifyRootPasswd() bool {
	var resultBool bool
	cmd := exec.Command("cat", "/etc/shadow")
	c, _ := cmd.Output()
	if cmd.Stderr != nil {
		fmt.Println("This is an error")
		fmt.Printf("%s\n", cmd.Stderr)
	}
	cmdOutStr := string(c[:])
	cmdOutStrSplit := strings.Split(cmdOutStr, "\n")
	for _, cmdOutStrLine := range cmdOutStrSplit {
		if cmdOutStrLineRoot := strings.Contains(cmdOutStrLine, "root"); cmdOutStrLineRoot {
			cmdOutStrLineRootSplit := strings.Split(string(cmdOutStrLine), ":")
			if cmdOutStrLineRootSplit[1] == "!" {
				resultBool = true
			} else {
				resultBool = false
			}
		}
	}
	return resultBool
}
