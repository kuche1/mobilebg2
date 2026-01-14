// TODO: make this into it's own package AND then add expiration date for each value

package persistentstorage

import (
	"log"

	"github.com/dgraph-io/badger/v4"
)

type PersistentStorage struct {
	db *badger.DB
}

func NewPersistentStorage(dbFolder string) *PersistentStorage {
	opts := badger.DefaultOptions(dbFolder)

	// disable prints
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
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return value, found
}

func (self *PersistentStorage) Write(key []byte, value []byte) {
	err := self.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})

	if err != nil {
		log.Fatal(err)
	}
}
