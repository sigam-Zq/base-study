package badgertest

import (
	"fmt"
	"log"
	"testing"

	"github.com/dgraph-io/badger"
)

var badgerDb *badger.DB

func TestMain(m *testing.M) {
	if badgerDb == nil {
		// Open the Badger database located in the /tmp/badger directory.
		// It will be created if it doesn't exist.
		var err error
		badgerDb, err = badger.Open(badger.DefaultOptions("./tmp"))
		if err != nil {
			log.Println("init DB Fail")
			log.Fatal(err)
			return
		}
		if badgerDb == nil {
			log.Println("badgerDb is nil  , nit DB Fail")
			return
		}
		// Your code here…
		log.Println("init DB Succeed")
		m.Run()

		defer func() {
			badgerDb.Close()
			log.Println("DB close Succeed")
		}()
	}
}

func TestSet(t *testing.T) {
	A1 := "a1"
	A2 := "a2"
	fmt.Println("0---------SET------------------------0")
	badgerDb.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(A1), []byte("bbbbb1"))
		if err != nil {
			return err
		}

		err = txn.Set([]byte(A2), []byte("bbbbb2"))
		if err != nil {
			return err
		}
		// err = txn.Commit()
		// if err != nil {
		// 	return err
		// }
		return nil
	})
}

func TestGet(t *testing.T) {

	badgerDb.View(func(txn *badger.Txn) error {
		item1, err := txn.Get([]byte("a"))
		if err != nil {
			return err
		}
		log.Printf("item1 %v\n", item1)
		valStr := ""
		err = item1.Value(func(val []byte) error {
			valStr = string(val)
			return nil
		})
		if err != nil {
			return err
		}
		log.Printf("item1 %s \n", valStr)

		return nil
	})
}

func TestIterating(t *testing.T) {
	err := badgerDb.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	t.Logf("---err %v \n", err)
}
