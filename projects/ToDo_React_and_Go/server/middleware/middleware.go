package middleware

import (
	"context"       // for context.WithValue
	"encoding/json" // for json.Marshal
	"fmt"           // for debugging
	"log"           // for logging
	"net/http"      // for http.StatusOK
	"os"            // for file handling

	"github.com/gorilla/mux"                                                            // for mux.Vars
	"github.com/iyashjayesh/Go-Lang-Resources/projects/ToDo_React_and_Go/server/models" // for models.Task

	// for routers.Router
	"github.com/joho/godotenv"                   // for godotenv.Load
	"go.mongodb.org/mongo-driver/bson"           // for bson.M
	"go.mongodb.org/mongo-driver/bson/primitive" // for primitive.ObjectID
	"go.mongodb.org/mongo-driver/mongo"          // for mongo.Client
	"go.mongodb.org/mongo-driver/mongo/options"  // for options.FindOne
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	fmt.Println("Loading the env file")

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func createDBInstance() {

	fmt.Println("Creating the DB instance")

	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")

	fmt.Println("Connection String: ", connectionString)
	fmt.Println("DB Name: ", dbName)
	fmt.Println("Collection Name: ", collectionName)

	// Connect to MongoDB
	clientOptions, err := options.Client().ApplyURI(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get the collection
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created!")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payLoad := getAllTasks()
	json.NewEncoder(w).Encode(payLoad)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList

	json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneTask(params["id"])
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTasks()

	json.NewEncoder(w).Encode(count)
}

// func DeleteAllCompletedTasks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	deleteAllCompletedTasks()
// 	json.NewEncoder(w).Encode("All completed tasks deleted")
// }

func getAllTasks() []primitive.M {
	cur, err := collection.Find(context.Background, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

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

	cur.Close(context.Background())

	return results
}

//bson.M -> M : An unordered representation of a BSON document (map)
//bson.D -> D : An ordered representation of a BSON document (map)
//bson.A -> A : An ordered representation of a BSON array
//bson.E -> E : A BSON element

func taskComplete(task_id string) {

	id, _ := primitive.ObjectIDFromHex(task_id)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": true}}
	result, err := collection.FindOneAndUpdate(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated Task: ", result)
	fmt.Println("Task Completed Count!", result.modifiedCount)
}

func undoTask(task string) {

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": false}}
	result, err := collection.FindOneAndUpdate(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated Task: ", result)
	fmt.Println("Task Completed Count!", result.modifiedCount)
}

func insertOneTask(task models.ToDoList) {
	insertresult, err := collection.insertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertresult.InsertedID)

}

func deleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Record", d.DeletedCount)
}

func deleteAllTasks() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Record", d.DeletedCount)
	return d.DeletedCount
}
