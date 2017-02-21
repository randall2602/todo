// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
        "fmt"
        "log"
        "buffio"
        "os"

        // Imports the Google Cloud Datastore client package.
        "cloud.google.com/go/datastore"
        "golang.org/x/net/context"
)

type database struct {
    projectID string
    ctx context.Context
    client *datastore.Client
}

func (db *database) Init(projectID string) {
    db.projectID = projectID
    db.ctx = context.Background()
    var err error
    db.client, err = datastore.NewClient(db.ctx, db.projectID)
    if err != nil {
            log.Fatalf("Failed to create client: %v", err)
    }
}

type Task struct {
    Description string
}

func (db *database) StoreTask(kind, name string, task Task) *datastore.Key {
    key := datastore.NameKey(kind, name, nil)
    if _, err := db.client.Put(db.ctx, key, &task); err != nil {
        log.Fatalf("Failed to save task: %v", err)
    }
    return key
}

func main() {
    var db database
    db.Init("alert-height-150418")
    scanner := buffio.NewScanner(os.Stdin)
    kind := "Task"
    taskCounter := 0
    fmt.Print("q to quit; enter first task: ")
    for scanner.Scan() {
        if scanner.Text() == "quit" {
            break
        }
        description := scanner.Text()
        taskCounter++
        name := "task" + toString(taskCounter)
        task := Task{
            Description: description,
        }
        fmt.Printf("Saved %v\n", db.StoreTask(kind, name, task))
        fmt.Print("enter next task: ")
    }
}