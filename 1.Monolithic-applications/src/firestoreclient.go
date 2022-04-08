package main

import (
	"context"
	"log"
	pb "monolithicapp/pb"

	ts "google.golang.org/protobuf/types/known/timestamppb"

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

func DeleteRemainder(remainder *pb.Remainder) error {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.Doc("appmod/" + remainder.GetUserUID() + "/remainders/" + remainder.GetRemainderID()).Delete(ctx)
	return err
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
