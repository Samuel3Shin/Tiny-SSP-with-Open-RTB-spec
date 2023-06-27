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
	testAdID := "ad1"
	entry := LogEntry{
		Time: time.Now(),
		AdID: testAdID,
	}

	// Insert the log entry into the database
	_, err = collection.InsertOne(context.Background(), entry)
	if err != nil {
		t.Fatalf("Failed to insert log entry: %v", err)
	}

	// Query the log entry back from the database
	var result LogEntry
	err = collection.FindOne(context.Background(), bson.M{"adID": testAdID}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find log entry: %v", err)
	}

	// Check if the ad ID is correct
	if result.AdID != testAdID {
		t.Errorf("Ad ID incorrect, got: %s, want: %s.", result.AdID, testAdID)
	}
}
