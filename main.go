package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var (
	functionname string
	region       string
	filepath     string
	publish      bool
	zipfile      []byte
)

func main() {
	flag.StringVar(&functionname, "functionname", "", "The name of the function to be updated")
	flag.StringVar(&region, "region", "us-east-1", "The region of the Lambda function")
	flag.StringVar(&filepath, "filepath", "", "The path of the zipfile to be uploaded")
	flag.BoolVar(&publish, "publish", true, "Should the updated code be published immediately?")
	flag.Parse()
	fmt.Printf("Reading file at %s\n", filepath)
	zipfile = readFile()
	fmt.Println("Successfully read the file")
	fmt.Println("Uploading the file to Lambda")
	runUpdate()
	fmt.Println("Successfully updated the Lambda function")
}

func readFile() []byte {
	functionData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("An error occurred while trying to read the zipfile")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return functionData
}

func runUpdate() {
	svc := lambda.New(session.New(), &aws.Config{Region: aws.String(region)})

	params := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(functionname),
		Publish:      aws.Bool(publish),
		ZipFile:      zipfile,
	}
	_, err := svc.UpdateFunctionCode(params)

	if err != nil {
		fmt.Println("An error occurred while trying to update the Lambda function.")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
