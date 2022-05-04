package main

import (
	"fmt"
	"os"
)

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// load secret via env
	//test
	awsToken := os.Getenv("AWS_TOKEN")
//new add
	func buildSql(email string) string {
  return fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email)
}

buildSql("oyetoketoby80@gmail.com")
	//test added
	
	f := "apple"
	fmt.Println(f)
}
