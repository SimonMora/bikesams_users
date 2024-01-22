package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/SimonMora/bikesams_users/aws_go"
	"github.com/SimonMora/bikesams_users/database"
	"github.com/SimonMora/bikesams_users/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	aws_go.InitAws()

	if !EnvironmentVariableValidation() {
		log.Default().Print("Error when parsing environment variable 'SecretName'")
		err := errors.New("The environment variable SecretName must be provided.")
		return event, err
	}

	var userCredentials models.SignUp
	eventCredentials := event.Request.UserAttributes

	for cred, att := range eventCredentials {
		switch cred {
		case "email":
			userCredentials.UserEmail = att
			log.Default().Println("User email registered in the SignUp class..")
		case "sub":
			userCredentials.UserUUID = att
			log.Default().Println("User uiid registered in the SignUp class..")
		}
	}

	err := database.ReadSecrets()
	if err != nil {
		log.Default().Println("Error reading the secret: " + err.Error())
		return event, err
	}

	err = database.SignUp(userCredentials)
	return event, err

}

func EnvironmentVariableValidation() bool {
	var isParam bool
	_, isParam = os.LookupEnv("SecretName")
	return isParam
}
