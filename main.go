package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sales struct {
	Id       int64              `bson:"_id,omitempty"`
	Date     primitive.DateTime `bson:"date,omitempty"`
	Item     string             `bson:"item,omitempty"`
	Price    float64            `bson:"price,omitempty"`
	Quantity int64              `bson:"quantity,omitempty"`
}

func (s *Sales) String() string {
	return fmt.Sprintf("Id : %v, Date : %v, Item : %v, Price :%v, Quantity :%v", s.Id, s.Date, s.Item, s.Price, s.Quantity)
}

func HandleHistory(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "전적!\n")
	fmt.Fprintf(w, "1게임 승\n")
	fmt.Fprintf(w, "2게임 승\n")
	fmt.Fprintf(w, "3게임 패 \n")
	fmt.Fprintf(w, "4게임 승\n")
	fmt.Fprintf(w, "5게임 승\n")
}

func HandleFriends(w http.ResponseWriter, r *http.Request) {

	// 데이터는 db에서 가져와야되고
	// 접속 확인
	var client *mongo.Client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://3dmp:VTwAnWPBJhwaZEWe@cluster0.vkgcv.mongodb.net"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	time.Sleep(2 * time.Second)

	// 접속
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	collection := client.Database("mongodbVSCodePlaygroundDB").Collection("sales")

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var sales []Sales
	if err = cur.All(ctx, &sales); err != nil {
		panic(err)
	}
	// fmt.Println(sales)

	for _, s := range sales {
		fmt.Fprintf(w, s.String())
	}

	// fmt.Fprintf(w, "친구 목록\n")
	// fmt.Fprintf(w, "제갈량\n")
	// fmt.Fprintf(w, "하우돈\n")
	// fmt.Fprintf(w, "장비\n")
	// fmt.Fprintf(w, "여포\n")
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

	http.ListenAndServe(":4000", nil)
}

// mongo drive install
