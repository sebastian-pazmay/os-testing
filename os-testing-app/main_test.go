package main

import (
	"fmt"
	"os-testing-app/testSuites"
	"testing"
)

var expectedVersion = "Ubuntu 18.04.5"

func TestSuite1(t *testing.T) {
	fmt.Println("This is the testing function")
	fmt.Println(testSuites.VerifyOSVersion(expectedVersion))
	fmt.Println(testSuites.VerifyRootPasswd())
}

func TestSuite2(t *testing.T) {
	fmt.Println("This is the testing function")
	testSuites.VerifyOSVersion(expectedVersion)
	testSuites.VerifyRootPasswd()
}

func main() {
	fmt.Println("This is the main function")
}
