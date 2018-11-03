package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Item e o nome do campo json (o tipo influencia para mandar para o dynamo)
type Item struct {
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, kinesisEvent events.KinesisEvent) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Printf("error on set the region of aws - testing again")
		os.Exit(1)
	}

	svc := dynamodb.New(sess)

	// Itera sobre cada registro que cai dentro do Kinesis do limit proposto
	for _, record := range kinesisEvent.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)

		fmt.Printf("%s Data: %s", record.EventName, dataText)

		item := Item{}
		json.Unmarshal([]byte(dataText), &item)
		item.createItem(svc)
	}

	return nil
}

// Cria o item dentro do Dynamo ou da erro
func (m *Item) createItem(svc *dynamodb.DynamoDB) error {
	marshedItem, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
	}

	fmt.Printf("MarshedItem: %s", marshedItem)

	input := &dynamodb.PutItemInput{
		Item:      marshedItem,
		TableName: aws.String("test_enjoei_dynamo"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("Successfully item create: { user_id: %i, created_at: %s }\n", m.UserID, m.CreatedAt)
	return nil
}
