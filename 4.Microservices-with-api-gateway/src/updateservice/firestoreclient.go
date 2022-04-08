package main

import (
	"context"
	"log"
	pb "updateservice/pb"

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

func UpdateRemainder(remainder *pb.Remainder) (*pb.Remainder, error) {

	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.Doc("appmod/"+remainder.GetUserUID()+"/remainders/"+remainder.GetRemainderID()).Update(ctx, []firestore.Update{
		{
			Path:  "Remainder",
			Value: remainder.GetRemainder(),
		},
	})

	return remainder, err
}
