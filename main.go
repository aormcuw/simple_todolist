package main

import (
	"time"

	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	colelctionName string = "todo"
	port           string = ":9000"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"compeleted"`
		CreatedAt time.Time     `bson:"createAt"`
	}
	todo struct {
		ID        bson.ObjectId `bson:"id"`
		Title     string        `bson:"title"`
		Completed string        `bson:"compeleted"`
		CreatedAt time.Time     `bson:"create_at"`
	}
)
