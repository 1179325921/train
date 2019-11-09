package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	//GetMongoDBClient()
	GetMongoDBClientS()
}

func GetMongoDBClient() {
	clientOptions := options.Client()
	clientOptions.Auth = &options.Credential{
		Username:"root",
		Password:"TCho8gLF38@oJ^G2" }

	//clientOptions.Hosts = []string{"dds-m5ed508dd35f28441.mongodb.rds.aliyuncs.com:3717","dds-m5ed508dd35f28442.mongodb.rds.aliyuncs.com:3717"}
	clientOptions.Hosts = []string{"dds-2ze548712360a2441.mongodb.rds.aliyuncs.com:3717",
		"dds-2ze548712360a2442.mongodb.rds.aliyuncs.com:3717"}
	replSet := "mgset-15868189"
	clientOptions.ReplicaSet = &replSet
	client, err := mongo.NewClient(clientOptions)
	if nil != err {
		fmt.Println("New Client failed:" + err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if nil != err {
		fmt.Println("Connect failed:" + err.Error())
	}
	err = client.Ping(ctx, readpref.Primary())
	fmt.Println("Ping " + err.Error())
}

func GetMongoDBClientS()  {
	url := "mongodb://root:1dMkheavJLEKxoKU@dds-m5ed508dd35f28441.mongodb.rds.aliyuncs.com:3717,dds-m5ed508dd35f28442.mongodb.rds.aliyuncs.com:3717/admin?replicaSet=mgset-15868189"
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if nil != err {
		fmt.Println("New Client failed:" + err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if nil != err {
		fmt.Println("Connect failed:" + err.Error())
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	fmt.Println("Ping ")
	fmt.Println(err)
	collection := client.Database("notifycentert").Collection("testing")
	timeString := time.Now().Format("2006-01-02 15:04:05")
	result, err := collection.InsertOne(ctx, TestPo{InsertTime: timeString})
	fmt.Println("result")
	fmt.Println(result)
	fmt.Println("err")
	fmt.Println(err)

	qResult := collection.FindOne(ctx, bson.M{"insertTime": timeString})
	var testPo TestPo
	err = qResult.Decode(&testPo)
	fmt.Println("qResult")
	fmt.Println(testPo)
	fmt.Println("qerr")
	fmt.Println(err)
}


type TestPo struct {
	InsertTime string `bson:"insertTime"`
}