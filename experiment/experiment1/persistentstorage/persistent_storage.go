package persistentstorage

// import (
// 	"log"
// 	"time"

// 	"go.etcd.io/bbolt"
// )

// type PersistentStorage struct {
// 	db         *bbolt.DB
// 	bucketName []byte
// }

// func NewPersistentStorage(dbPath string) *PersistentStorage {
// 	db, err := bbolt.Open(
// 		dbPath,
// 		0600,
// 		&bbolt.Options{Timeout: 5 * time.Second}, // Opening an already open Bolt database will cause it to hang until the other process closes it. To prevent an indefinite wait you can pass a timeout option to the Open() function
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &PersistentStorage{
// 		db:         db,
// 		bucketName: []byte("A"), // "SomeBucket"
// 	}
// }

// func (self *PersistentStorage) Close() {
// 	self.db.Close()
// }

// func (self *PersistentStorage) Read(key []byte) (_data []byte, _found bool) {
// 	var value []byte
// 	var found bool

// 	err := self.db.View(func(tx *bbolt.Tx) error {
// 		b := tx.Bucket(self.bucketName)
// 		val := b.Get(key)

// 		if val == nil {
// 			found = false
// 			return nil
// 		}

// 		// `val` needs to be copied
// 		// otherwise it the value in `value` "dissapears" after the function returns

// 		// value = make([]byte, len(val))
// 		// copy(value, val)

// 		value = append([]byte(nil), val...)
// 		found = true

// 		return nil
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return value, found
// }

// func (self *PersistentStorage) Write(key []byte, value []byte) {
// 	err := self.db.Update(func(tx *bbolt.Tx) error {
// 		bucket, err := tx.CreateBucketIfNotExists(self.bucketName)
// 		if err != nil {
// 			return err
// 		}
// 		return bucket.Put(key, value)
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
