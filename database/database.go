package database

import (
	"context"
	"log"
	"time"

	"github.com/doha-ms/go-graphql-mongodb-crud/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb://localhost:27017"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("MongoDB Ping failed: ", err)
	}

	log.Println("Successfully connected to MongoDB!")

	return &DB{
		client: client,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var joblisting model.JobListing
	err := jobCollec.FindOne(ctx, filter).Decode(&joblisting)
	if err != nil {

		log.Fatal(err)

	}
	return &joblisting

}

func (db *DB) GetJobs() []*model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var joblistings []*model.JobListing
	cursor, err := jobCollec.Find(ctx, bson.D{})
	if err != nil {

		log.Fatal(err)

	}
	if err = cursor.All(context.TODO(), &joblistings); err != nil {
		panic(err)
	}
	return joblistings

}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserted, err := jobCollec.InsertOne(ctx, bson.M{
		"title":       jobInfo.Title,
		"description": jobInfo.Description,
		"url":         jobInfo.URL,
		"company":     jobInfo.Company,
	})
	if err != nil {
		log.Fatal(err)
	}
	var insertedID string
	if oid, ok := inserted.InsertedID.(primitive.ObjectID); ok {
		insertedID = oid.Hex()
	}
	returnJobListing := model.JobListing{ID: insertedID, Title: jobInfo.Title,
		Description: jobInfo.Description,
		URL:         jobInfo.URL,
		Company:     jobInfo.Company}

	return &returnJobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title

	}

	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description

	}

	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL

	}
	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *DB) DeleteJobListing(jobID string) *model.DeleteJobResponse {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobID)
	filter := bson.M{"_id": _id}
	_, err := jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return &model.DeleteJobResponse{DeletedJobID: jobID}
}
