package main

import (
	pb "addservice/pb"
	"context"
	"log"

	ts "google.golang.org/protobuf/types/known/timestamppb"

	"cloud.google.com/go/firestore"
)

func createClient(ctx context.Context) *firestore.Client {

	projectID := ""
	mustMapEnv(&projectID, "PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func AddRemainder(remainder *pb.Remainder) (*pb.Remainder, error) {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	remainder.CreatedAt = ts.Now()
	doc, _, err := client.Collection("appmod/"+remainder.GetUserUID()+"/remainders").Add(ctx, remainder)
	if err != nil {
		log.Fatalf("Failed adding remainder: %v", err)
		return nil, err
	}

	remainder.RemainderID = doc.ID
	return remainder, nil
}
