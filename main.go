package main

import (
	"context"
	"fmt"
	"gamedata/db/model"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func HandleHistory(w http.ResponseWriter, r *http.Request) { in")

	fmt.Fprintf(w, "전적!\n")
	fmt.Fprintf(w, "1게임 승\n")
	fmt.Fprintf(w, "2게임 승\n")
	fmt.Fprintf(w, "3게임 패 \n")
	fmt.Fprintf(w, "4게임 승\n")
	fmt.Fprintf(w, "5게임 승\n")
}

// TODO : error handling
func HandleFriends(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HandleFriends(w http.ResponseWriter, r *http.Request) in")

	// 데이터는 db에서 가져와야되고
	// 접속 확인
	var client *mongo.Client
	var err error
	// client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://3dmp:VTwAnWPBJhwaZEWe@cluster0.vkgcv.mongodb.net")) //몽고DB 접속클라 만듬
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.9:27017")) //몽고DB 접속클라 만듬
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//time.Sleep(2 * time.Second)

	// 접속		//실제 접속
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("mongodbVSCodePlaygroundDB").Collection("sales")

	//bson
	// select * from sales;
	// select id from sales;
	// select price from sales;
	// bson.D{id}
	// bson.D{price}
	cur, currErr := collection.Find(ctx, bson.D{}) //base.D

	if cur.RemainingBatchLength() == 0 {
		fmt.Fprintf(w, "no data")
	}

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var sales []model.Sales
	if err = cur.All(ctx, &sales); err != nil {
		panic(err)
	}
	// fmt.Println(sales)

	for _, s := range sales {
		fmt.Fprintf(w, s.String())
	}
}

func HandleUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HandleUser(w http.ResponseWriter, r *http.Request) in")

	// 데이터는 db에서 가져와야되고
	// 접속 확인
	var client *mongo.Client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.9:27017")) //몽고DB 접속클라 만듬
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	user := client.Database("gamedata").Collection("user")

	cur, currErr := user.Find(ctx, bson.D{}) //base.D

	if cur.RemainingBatchLength() == 0 {
		fmt.Fprintf(w, "no data")
	}

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var users []model.Users
	if err = cur.All(ctx, &users); err != nil {
		panic(err)
	}
	// fmt.Println(sales)

	for _, s := range users {
		fmt.Fprintf(w, s.String()+"\n")
	}
}

func main() {

	// TODO:

	// authentication
	// login된 jwt 사용해서, 유저가 공개하기 싫은 데이터는 jwt인증된 사람에게만 공개
	// jwt는 json web token인데, 예전의 cookie session과 같은 인증을 token으로 처리하는 방법
	// cookie session은 서버가 1개일때 편함, jwt는 msa 서버에서 편함

	// 민감한 아이템 같은 경우 숨겨야함
	// http => https
	// non secure => TLS

	// middleware
	// 공동된 처리를 하는 로직
	// jwt를 공통적으로 처리하는 로직
	// log를 공통적으로 처리하는 로직

	// protocol change : http => protobuf
	// get data from database : use mongodb
	// mongodb로 먼저 간 후, 아이템이나 db 확장이 fix되는 시점에 rdbms로 넘어가는 것도 좋음 (안 넘어가도 좋음)
	// kafka로 pub sub 연결 : 실제 게임 캐릭터가 렙업을 했거나 아이템을 먹었을 때, 보이는 화면에서 표시하기 위함

	// 전적 검색
	http.HandleFunc("/history", HandleHistory)

	// 함께하는 친구 검색
	http.HandleFunc("/friends", HandleFriends)

	http.HandleFunc("/users", HandleUser)

	port := 4000
	fmt.Println("Gamedata server is running on :", port)
	http.ListenAndServe(":4000", nil)

	// err := http.ListenAndServeTLS(":4000", "testdata/x509/ca_cert.pem", "testdata/x509/ca_key.pem", nil)

	// err := http.ListenAndServeTLS(":4000", "testdata/x509/server_cert.pem", "testdata/x509/server_key.pem", nil)

	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

}

// mongo drive install
