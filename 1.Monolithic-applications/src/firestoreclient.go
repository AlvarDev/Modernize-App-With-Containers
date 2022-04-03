package main

import (
	"context"
	"log"
	pb "monolithicapp/pb"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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

func ListRemainders(userUID string) (*pb.ListRemaindersResponse, error) {

	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var rms []*pb.Remainder
	remaindersIter := client.Collection("remainders").Documents(ctx)

	for {

		doc, err := remaindersIter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		remainder := &pb.Remainder{}
		doc.DataTo(&remainder)
		rms = append(rms, remainder)
	}

	return &pb.ListRemaindersResponse{Remainders: rms}, nil
}
