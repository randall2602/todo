// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
        "fmt"
        "log"

        // Imports the Google Cloud Datastore client package.
        "cloud.google.com/go/datastore"
        "golang.org/x/net/context"
)

type database struct {
    projectID string
    ctx context.Context
    client datastore.Client
}

func (db *database) Init(projectID string) {
    db.projectID = projectID
    db.ctx = context.Background()
    db.client, err := datastore.NewClient(db.ctx, db.projectID)
    if err != nil {
            log.Fatalf("Failed to create client: %v", err)
    }
}

type Task struct {
    Description string
}

func (db *database) StoreTask(kind, name string, task Task) string {
    key := datastore.NameKey(kind, name, nill)
    if _, err := db.client.Put(db.ctx, key, &task); err != nill {
        log.Fatalf("Failed to save task: %v", err)
    }
    return key
}

func main() {
    var db database
    db.Init("alert-height-150418")
    kind := "Task"
    name := "task001"
    task := Task{
            Description: "Buy milk",
    }
    fmt.Printf("Saved %v\n", db.StoreTask(kind, name, task))
}