package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/FerretDB/FerretDB/v2/ferretdb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// runExampleClient shows an example of running MongoDB client with embedded FerretDB.
func runExampleClient(uri string) {
	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
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
	f, err := ferretdb.New(&ferretdb.Config{
		PostgreSQLURL: "postgres://postgres:password@127.0.0.1:5432/postgres",
		ListenAddr:    "127.0.0.1:27017",
		StateDir:      ".",
		LogLevel:      slog.LevelInfo,
		LogOutput:     os.Stderr,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	done := make(chan struct{})

	go func() {
		f.Run(ctx)
		close(done)
	}()

	uri := f.MongoDBURI()
	log.Printf("Embedded FerretDB started, use %s to connect.", uri)

	runExampleClient(uri)

	<-done
}
