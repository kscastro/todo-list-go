package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/kscastro/todo-list-go/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	GetAllTask() []primitive.M
	InsertOneTask(task model.TodoList)
	TaskComplete(task string)
	UndoTask(task string)
	DeleteOneTask(task string)
	DeleteAllTask() int64
}

type DB struct {
	redis      *redis.Client
	collection *mongo.Collection
}

var taskRepository TaskRepository = DB{}

func NewDB() *DB {
	connectionString := os.Getenv("MONGO_URL")
	dbName := os.Getenv("MONGO_DB_NAME")
	collName := os.Getenv("MONGO_DB_COLLECTION_NAME")
	clientOptions := options.Client().ApplyURI(connectionString)

	clientMongo, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = clientMongo.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collectionMongo := clientMongo.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")

	opt := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	}

	clientRedis := redis.NewClient(opt)
	pong, err := clientRedis.Ping().Result()
	if err != nil {
		log.Panic(err, pong)
	}

	return &DB{redis: clientRedis, collection: collectionMongo}
}

func (d DB) GetAllTask() []primitive.M {

	val, err := d.redis.Get("allTasks").Result()
	if err == nil && val != "" {

		fmt.Println("Pegando do Redis")

		result := make([]primitive.M, 0)

		err := json.Unmarshal([]byte(val), &result)
		if err != nil {
			log.Fatal(err)
		}

		return result
	}

	fmt.Println("Pegando do Mongo")

	cur, err := d.collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	resultStr, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	err = d.redis.Set("allTasks", string(resultStr), 5*time.Second).Err()
	if err != nil {
		log.Panic(err)
	}
	return results
}

func (d DB) InsertOneTask(task model.TodoList) {
	insertResult, err := d.collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func (d DB) TaskComplete(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := d.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func (d DB) UndoTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := d.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func (d DB) DeleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	doc, err := d.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", doc.DeletedCount)
}

func (d DB) DeleteAllTask() int64 {
	doc, err := d.collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", doc.DeletedCount)
	return doc.DeletedCount
}
