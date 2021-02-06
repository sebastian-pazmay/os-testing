package main

import (
	"fmt"
    "os-testing/testSuites"
)

var expectedVersion = "Ubuntu 18.04.5"

func main() {
	fmt.Println("## TEST CASES ##")
	fmt.Println("1. Verify software OS version")
	fmt.Println(testSuites.VerifyOSVersion(expectedVersion))
	fmt.Println("2. Verify root password is not empty")
	fmt.Println(testSuites.VerifyRootPasswd())
}
