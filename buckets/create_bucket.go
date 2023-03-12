package buckets

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"object-storage.app/utils"
)

// Encode in your terminal with: echo -n "your string" | base64
// decode base64 string to json
// This is how we get around commiting our credentials to github.
func CovertStringToJSON(env_details string) []byte {
	decoded_json, err := base64.StdEncoding.DecodeString(env_details)
	if err != nil {
		panic(err)
	}

	return decoded_json

}

func clientConnection() *storage.Client {
	ctx := context.Background()

	// Set up credentials to authenticate with Google Cloud
	var bucketVar string

	if env := os.Getenv("ENV"); env != "PROD" {
		bucketVar = utils.GoDotEnvVariable("./.env", "BUCKET_CREATOR")
	} else {
		bucketVar = os.Getenv("BUCKET_CREATOR")
	}

	new_creds := CovertStringToJSON(bucketVar)
	creds := option.WithCredentialsJSON(new_creds)

	// Create a client to interact with the Google Cloud Storage API
	client, err := storage.NewClient(ctx, creds)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client

}

func CreateBucket(projectID string, bucketName string) {

	ctx := context.Background()

	client := clientConnection()

	// Create a new bucket in the given project with the given name if the bucket
	// already exists then it will just print (bucket already exists)
	if err := client.Bucket(bucketName).Create(ctx, projectID, nil); err != nil {
		fmt.Printf("Failed to create bucket %q: %v", bucketName, err)
	}

	fmt.Printf("Bucket %q created.\n", bucketName)
}
