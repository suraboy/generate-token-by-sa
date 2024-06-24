package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

func main() {
	// Path to the service account JSON key file
	keyFilePath := "./storage/paotang-core-sit-d50437475f5b.json"

	// Target audience for the ID token
	targetAudience := "https://paotang-telco-check-6ti7spcsya-as.a.run.app"

	// Create ID token source with service account credentials
	tokenSource, err := idtoken.NewTokenSource(context.Background(), targetAudience, option.WithCredentialsFile(keyFilePath))
	if err != nil {
		log.Fatalf("Failed to create ID token source: %v", err)
	}

	// Get the ID token
	token, err := tokenSource.Token()
	if err != nil {
		log.Fatalf("Failed to get ID token: %v", err)
	}

	fmt.Printf("ID Token: %s\n", token.AccessToken)
}
