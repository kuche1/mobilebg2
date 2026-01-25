// TODO: make this into it's own package
// TODO: add the ability to delete (???)

package persistentstorage

import (
	"encoding/binary"
	"log"
	"math"
	"time"

	"github.com/dgraph-io/badger/v4"
)

const ExpirationDateNeverUnixSec = math.MaxInt64

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
	data, expired, found := self.ReadExpired(key)
	if !found {
		return nil, false
	}
	if expired {
		return nil, false
	}
	return data, true
}

func (self *PersistentStorage) ReadExpired(key []byte) (_data []byte, _expired bool, _found bool) {
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

	if !found {
		return nil, true, false
	}

	if len(value) < 8 {
		log.Fatal("Database is corrupted")
	}

	expirationDateUnixSec := bytesToInt64(value)
	value = value[8:]
	if expirationDateUnixSec == ExpirationDateNeverUnixSec {
		return value, false, true
	}

	nowUnixSec := time.Now().Unix()
	if nowUnixSec > expirationDateUnixSec {
		return value, true, true
	}

	return value, false, true
}

func (self *PersistentStorage) Write(key []byte, value []byte) {
	self.WriteWithExpirationAt(key, value, ExpirationDateNeverUnixSec)
}

// The value will be considered expired after that much time has passed relative to the moment of writing
func (self *PersistentStorage) WriteWithExpirationAfter(key []byte, value []byte, validitySec int64) {
	expirationDate := time.Now().Unix() + validitySec
	self.WriteWithExpirationAt(key, value, expirationDate)
}

// The value will be considered expired at the given date
func (self *PersistentStorage) WriteWithExpirationAt(key []byte, value []byte, expirationDateUnixSec int64) {
	expirationAsBytes := int64ToBytes(expirationDateUnixSec)

	expirationAndValue := append(expirationAsBytes, value...)

	err := self.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, expirationAndValue)
	})

	if err != nil {
		log.Fatal(err)
	}
}

func int64ToBytes(number int64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(number))
	return bytes
}

func bytesToInt64(bytes []byte) int64 {
	return int64(binary.BigEndian.Uint64(bytes))
}
