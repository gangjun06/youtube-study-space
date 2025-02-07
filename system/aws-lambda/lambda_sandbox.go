package main

import (
	"app.modules/aws-lambda/lambdautils"
	"app.modules/core"
	"app.modules/core/utils"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"time"
)

type LambdaSandBoxResponseStruct struct {
	Result  string       `json:"result"`
	Message string       `json:"message"`
}

// LambdaSandBox api gatewayで使わないから、引数と返却値はなしでいいと思う
func LambdaSandBox(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("LambdaSandBox()")
	
	ctx := context.Background()
	clientOption, err := lambdautils.FirestoreClientOption()
	if err != nil {
		return lambdautils.ErrorResponse(err)
	}
	_system, err := core.NewSystem(ctx, clientOption)
	if err != nil {
		return lambdautils.ErrorResponse(err)
	}
	defer _system.CloseFirestoreClient()
	
	log.Println("time.Now(): " + time.Now().Format(time.RFC3339))
	log.Println("JstNow(): " + utils.JstNow().Format(time.RFC3339))
	
	return LambdaSandBoxResponse()
}

func LambdaSandBoxResponse() (events.APIGatewayProxyResponse, error) {
	var apiResp LambdaSandBoxResponseStruct
	apiResp.Result = lambdautils.OK
	jsonBytes, _ := json.Marshal(apiResp)
	return lambdautils.Response(jsonBytes)
}

func main() {
	lambda.Start(LambdaSandBox)
}
