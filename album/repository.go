package album

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// Repository ...
type Repository struct{}

// SERVER is the DB server
const SERVER = "localhost:27017"

// DBNAME is the name of the DB instance
const DBNAME = "musicstore"

// DOCNAME is the name of the collection to use for album docs
const DOCNAME = "albums"

// GetAlbums returns the list of Albums
func (r Repository) GetAlbums() Albums {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish a connection to Mongo server:", err)
	}
	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddAlbum inserts an Album to the DB
func (r Repository) AddAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(album)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateAlbum updates an Album in the DB (not used for now)
func (r Repository) UpdateAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	session.DB(DBNAME).C(DOCNAME).UpdateId(album.ID, album)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteAlbum deletes an Album
func (r Repository) DeleteAlbum(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	if err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	// verify id is ObjectId, otherwise halt
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}

	// grab id
	oid := bson.ObjectIdHex(id)

	// remove user
	if err := session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	return "OK"
}
