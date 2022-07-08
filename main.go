package main

import (
	"context"
	"log"
	"os"

	"github.com/FerretDB/FerretDB/ferretdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// runExampleClient shows an example of running MongoDB client with embedded FerretDB.
func runExampleClient(uri string) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// insert data
	collection := client.Database("test").Collection("test")
	if _, err = collection.InsertOne(ctx, bson.D{{"hello", "world"}}); err != nil {
		log.Fatal(err)
	}

	// show a list of databases
	databases, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Databases: %v", databases)

	// show inserted data
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var res []bson.D
	if err = cursor.All(ctx, &res); err != nil {
		log.Fatal(err)
	}
	log.Printf("Data in test.test: %v", res)
}

func main() {
	log.SetOutput(os.Stderr)

	f, err := ferretdb.New(&ferretdb.Config{
		Handler:       "pg",
		PostgreSQLURL: "postgres://postgres:password@127.0.0.1:5432/postgres",
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	done := make(chan error)
	go func() {
		done <- f.Run(ctx)
	}()

	uri := f.MongoDBURI()
	log.Printf("Embedded FerretDB started, use %s to connect.\n", uri)

	runExampleClient(uri)

	log.Fatal(<-done)
}
