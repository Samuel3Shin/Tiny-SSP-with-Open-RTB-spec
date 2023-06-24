package logserver

import (
	"context"
	"testing"
	"time"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestLogToDB(t *testing.T) {
	// Setup MongoDB client
	cfg := common.GetConfig()
	clientOptions := options.Client().ApplyURI(cfg.MONGODB_URL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Get the logs collection
	collection := client.Database("logDB").Collection("logs")

	// Create a test log entry
	testLog := "test log"
	entry := LogEntry{
		Time:    time.Now(),
		LogText: testLog,
	}

	// Insert the log entry into the database
	_, err = collection.InsertOne(context.Background(), entry)
	if err != nil {
		t.Fatalf("Failed to insert log entry: %v", err)
	}

	// Query the log entry back from the database
	var result LogEntry
	err = collection.FindOne(context.Background(), bson.M{"logText": testLog}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find log entry: %v", err)
	}

	// Check if the log entry content is correct
	if result.LogText != testLog {
		t.Errorf("Log content incorrect, got: %s, want: %s.", result.LogText, testLog)
	}
}
