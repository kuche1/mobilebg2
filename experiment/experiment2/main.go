package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
)

func main() {
	ps := NewPersistentStorage("db")

	///// write

	ps.Write([]byte("key"), []byte("valueasdeeefdsfsdfds"))

	///// Read

	value, found := ps.Read([]byte("key"))
	fmt.Println("key =", string(value))
	fmt.Println("key =", found)
}

type PersistentStorage struct {
	db *badger.DB
}

func NewPersistentStorage(dbFolder string) *PersistentStorage {
	opts := badger.DefaultOptions(dbFolder)

	// disable print
	opts.Logger = nil

	db, err := badger.Open(opts)
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

func (self *PersistentStorage) Write(key []byte, value []byte) {
	err := self.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (self *PersistentStorage) Read(key []byte) (_data []byte, _found bool) {
	found := false
	var value []byte

	err := self.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil
			}
			return err
		}

		value, err = item.ValueCopy(nil)

		if err != nil {
			return err
		}

		found = true

		fmt.Println("key =", string(value))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("key =", string(value))

	return value, found
}

//////////
////////// old code
//////////

// func main() {
// 	opts := badger.DefaultOptions("badger_db")

// 	// disable print
// 	opts.Logger = nil

// 	db, err := badger.Open(opts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	///// write

// 	err = db.Update(func(txn *badger.Txn) error {
// 		return txn.Set([]byte("key"), []byte("valueeeee"))
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	///// Read

// 	err = db.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get([]byte("key"))
// 		if err != nil {
// 			return err
// 		}
// 		val, _ := item.ValueCopy(nil)
// 		fmt.Println("key =", string(val))
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
