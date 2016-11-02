package models

import (
	"gopkg.in/mgo.v2"
	"os"
	"log"
)

type CollectionMappable interface {
	getCollection() string
}

var databaseSession *mgo.Session

func BootDatabaseConn() {
	//databaseDial := &mgo.DialInfo{
	//	Addrs: []string{os.Getenv("MONGODB_DIAL_HOST")},
	//	Timeout: 60 * time.Second,
	//	Database: os.Getenv("MONGODB_DATABASE"),
	//	Username: os.Getenv("MONGODB_USER"),
	//	Password: os.Getenv("MONGODB_PASSWORD"),
	//}

	session, err := mgo.Dial(os.Getenv("MONGODB_HOST"))

	//databaseSession, err := mgo.DialWithInfo(databaseDial)

	if err != nil {
		log.Fatal("Error creating database session", err)
	}

	databaseSession = session.Copy()
	databaseSession.SetMode(mgo.Monotonic, true)
}

func getCollectionFromModel(collection CollectionMappable, callback func(*mgo.Collection)) {
	if databaseSession == nil {
		log.Fatal("Database session not initialized")
	}

	sessionCopy := databaseSession.Copy()
	defer sessionCopy.Close()
	callback(sessionCopy.DB(os.Getenv("MONGODB_DATABASE")).C(collection.getCollection()))
}