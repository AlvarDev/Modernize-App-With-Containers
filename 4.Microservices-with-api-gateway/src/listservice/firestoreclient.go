package main

import (
	"context"
	"log"

	pb "listservice/pb"

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

func ListRemainders(userUID string) ([]*pb.Remainder, error) {

	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var rms []*pb.Remainder
	remaindersIter := client.
		Collection("appmod/"+userUID+"/remainders").
		OrderBy("CreatedAt", firestore.Asc).
		Limit(30).
		Documents(ctx)

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
		remainder.RemainderID = doc.Ref.ID
		rms = append(rms, remainder)
	}

	return rms, nil
}
