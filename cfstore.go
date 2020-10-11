package main


import ( 
	"os"
	"context"
	"log"
	firebase "firebase.google.com/go/v4" 
 	"github.com/joho/godotenv"  
)

func Init() {

	//initialize env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//get projectID from env
	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatal("Error finding environment variable")
	}

	ctx := context.Background()

	conf := &firebase.Config{ProjectID: projectID}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return	
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("Firestore init: %v", err)
		return	
	}

	defer client.Close()
}