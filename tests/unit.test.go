package tests

import (
	"fmt"
	"go-api/services"
)

func UnitTest() {
	fmt.Println("Start Test==========================")
	conn := services.ExecuteCmd("docker ps", "0.0.0.0", "22", "user", "pass")
	// result := conn.Cmd("docker ps")
	fmt.Println("Out : %s", conn)
	fmt.Println("Test End==========================")
}
