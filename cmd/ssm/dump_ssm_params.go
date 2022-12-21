package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/devlibx/gox-aws/cmd/ssm/impl"
)

func main() {
	path := flag.String("path", "", "provide the path to fetch")
	printMode := flag.String("print", "export", "print export logs")
	flag.Parse()

	// Create a new AWS newSession
	newSession, err := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-south-1"),
		})
	if err != nil {
		panic(err)
	}

	result, err := impl.ReadSsmFromPath(*path, newSession)
	if err != nil {
		panic(err)
	}

	for _, p := range result {
		if *printMode == "export" {
			fmt.Printf("export %s=%s\n", p.Key, p.Value)
		} else {
			fmt.Printf("%s=%s\n", p.Key, p.Value)
		}
	}
	fmt.Println("")
}
