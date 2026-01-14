package main

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

// func main() {
// 	db, err := bbolt.Open(
// 		"whatever.db",
// 		0600,
// 		&bbolt.Options{Timeout: 5 * time.Second}, // Opening an already open Bolt database will cause it to hang until the other process closes it. To prevent an indefinite wait you can pass a timeout option to the Open() function
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	///// write

// 	err = db.Update(func(tx *bbolt.Tx) error {
// 		bucket, err := tx.CreateBucketIfNotExists([]byte("SomeBucketName"))
// 		if err != nil {
// 			return err
// 		}
// 		return bucket.Put([]byte("key"), []byte("value"))
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	///// read

// 	var value []byte

// 	err = db.View(func(tx *bbolt.Tx) error {
// 		b := tx.Bucket([]byte("SomeBucketName"))
// 		val := b.Get([]byte("key"))
// 		copy(value, val) // TODO: I think it will fail otherwise
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println(string(value))
// }

func main() {
	// db, err := bbolt.Open(
	// 	"whatever.db",
	// 	0600,
	// 	&bbolt.Options{Timeout: 5 * time.Second}, // Opening an already open Bolt database will cause it to hang until the other process closes it. To prevent an indefinite wait you can pass a timeout option to the Open() function
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	ps := NewPersistentStorage()
	defer ps.Close()

	///// read

	// var value []byte

	// err := ps.db.View(func(tx *bbolt.Tx) error {
	// 	b := tx.Bucket([]byte("SomeBucketName"))
	// 	val := b.Get([]byte("key"))
	// 	copy(value, val) // TODO: I think it will fail otherwise
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(string(value))

	value, found := ps.Read([]byte("key"))
	if found {
		fmt.Println(string(value))
	} else {
		fmt.Println("Not Found")
	}

	///// write

	// err := ps.db.Update(func(tx *bbolt.Tx) error {
	// 	bucket, err := tx.CreateBucketIfNotExists([]byte("SomeBucketName"))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return bucket.Put([]byte("key"), []byte("valueeee"))
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ps.Write([]byte("key"), []byte("fggfgfdgfdgfdh"))
}

//////////
////////// persistent storage
//////////

type PersistentStorage struct {
	db *bbolt.DB
}

func NewPersistentStorage() *PersistentStorage {
	db, err := bbolt.Open(
		"whatever.db",
		0600,
		&bbolt.Options{Timeout: 5 * time.Second}, // Opening an already open Bolt database will cause it to hang until the other process closes it. To prevent an indefinite wait you can pass a timeout option to the Open() function
	)
	if err != nil {
		log.Fatal(err)
	}

	return &PersistentStorage{
		db: db,
	}
}

func (self *PersistentStorage) Close() {
	self.db.Close()
}

func (self *PersistentStorage) Read(key []byte) (_data []byte, _found bool) {
	var value []byte
	var found bool

	err := self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("SomeBucketName")) // TODO: use const
		val := b.Get(key)

		if val == nil {
			found = false
			return nil
		}

		// `val` needs to be copied
		// otherwise it the value in `value` "dissapears" after the function returns

		// value = make([]byte, len(val))
		// copy(value, val)

		value = append([]byte(nil), val...)
		found = true

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return value, found
}

func (self *PersistentStorage) Write(key []byte, value []byte) {
	err := self.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("SomeBucketName")) // TODO: use const
		if err != nil {
			return err
		}
		return bucket.Put(key, value)
	})
	if err != nil {
		log.Fatal(err)
	}
}
