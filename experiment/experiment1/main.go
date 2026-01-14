package main

import (
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open(
		"whatever.db",
		0600,
		&bbolt.Options{Timeout: 5 * time.Second}, // Opening an already open Bolt database will cause it to hang until the other process closes it. To prevent an indefinite wait you can pass a timeout option to the Open() function
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	///// write

	err = db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("SomeBucketName"))
		if err != nil {
			return err
		}
		return bucket.Put([]byte("key"), []byte("value"))
	})
	if err != nil {
		log.Fatal(err)
	}

	///// read

	var value []byte

	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("SomeBucketName"))
		val := b.Get([]byte("key"))
		copy(value, val) // TODO: I think it will fail otherwise
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(value))
}
