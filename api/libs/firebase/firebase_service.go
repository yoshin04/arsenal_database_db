package libs

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type IFirebaseService interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

type firebaseService struct {
	AuthClient *auth.Client
}

func NewFirebaseService(ctx context.Context) (*firebaseService, error) {
	privateKey := strings.ReplaceAll(os.Getenv("FIREBASE_PRIVATE_KEY"), "\\n", "\n")
	firebaseConfig := map[string]string{
		"type":                        os.Getenv("FIREBASE_TYPE"),
		"project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
		"private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		"private_key":                 privateKey,
		"client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
		"auth_uri":                    os.Getenv("FIREBASE_AUTH_URI"),
		"token_uri":                   os.Getenv("FIREBASE_TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
	}

	configJSON, err := json.Marshal(firebaseConfig)
	if err != nil {
		log.Fatalf("Error marshalling firebase config: %v", err)
		return nil, err
	}

	opt := option.WithCredentialsJSON(configJSON)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseService{
		AuthClient: authClient,
	}, nil
}

func (fs *firebaseService) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	log.Println("Running firebaseService.VerifyIDToken")
	token, err := fs.AuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("Error verifying ID token: %v", err)
		return nil, err
	}
	return token, nil
}
