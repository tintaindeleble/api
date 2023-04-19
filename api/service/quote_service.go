package service

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func New() error {
	clientFireStore := createClient()
	defer clientFireStore.Close()

	_, _, err := clientFireStore.Collection("tinta-indeleble").Add(context.TODO(), map[string]interface{}{
		"id":      "123",
		"message": "Erase una vez lo que nunca fue.",
		"author":  "Anton Tiruriro",
		"work":    "1815",
	})
	if err != nil {
		return err
	}

	return nil
}

func createClient() *firestore.Client {
	projectID := "site-tinta-indeleble"

	client, err := firestore.NewClient(context.TODO(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
