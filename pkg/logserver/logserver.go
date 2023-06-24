package logserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogEntry struct {
	Time    time.Time `bson:"time"`
	LogText string    `bson:"logText"`
}

var logCollection *mongo.Collection

func init() {
	// Initialize MongoDB client and collection
	cfg := common.GetConfig()
	clientOptions := options.Client().ApplyURI(cfg.MONGODB_URL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	logCollection = client.Database("logDB").Collection("logs")
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the data
	logToDB(string(body))
	// fmt.Print("Received log: ", string(body))
	// Respond with a 200 OK
	fmt.Fprint(w, "OK")
}

func logToDB(logText string) {
	// Create a log entry
	entry := LogEntry{
		Time:    time.Now(),
		LogText: logText,
	}

	// Insert the log entry into the database
	_, err := logCollection.InsertOne(context.Background(), entry)
	if err != nil {
		log.Printf("Failed to insert log entry: %v", err)
	}
}

func StartServer() {
	http.HandleFunc("/", LogHandler)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
