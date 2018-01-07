package rocks

import (
	"log"
	"os"

	"github.com/tecbot/gorocksdb"
	"sync/atomic"
	"time"
	"errors"
)

type Rocks struct {
	path             string
	db               *gorocksdb.DB
	dbOptions        *gorocksdb.Options
	wo               *gorocksdb.WriteOptions
	ro               *gorocksdb.ReadOptions
	compactionFilter *shardingCompactionFilter

	// used for locking
	clientCounter int32
}

var (
	ERR_SHUTDOWN = errors.New("shutdown in progress")
)

func New(path string) *Rocks {
	r := &Rocks{
		compactionFilter: &shardingCompactionFilter{},
	}
	r.setup(path)
	return r
}

func (d *Rocks) setup(path string) {
	d.path = path
	d.dbOptions = gorocksdb.NewDefaultOptions()
	d.dbOptions.SetCreateIfMissing(true)
	// TODO Required but not avaiable for now
	d.dbOptions.SetAllowIngestBehind(true)
	d.dbOptions.SetCompactionFilter(d.compactionFilter)

	var err error
	d.db, err = gorocksdb.OpenDb(d.dbOptions, d.path)
	if err != nil {
		log.Fatal(err)
	}

	d.wo = gorocksdb.NewDefaultWriteOptions()
	//d.wo.DisableWAL(true)
	d.ro = gorocksdb.NewDefaultReadOptions()
}

func (d *Rocks) Put(key []byte, msg []byte) (err error) {
	// println("put", string(key), "value", string(msg))
	if newClientCounter := atomic.AddInt32(&d.clientCounter, 1); newClientCounter > 0 {
		err = d.db.Put(d.wo, key, msg)
	} else {
		err = ERR_SHUTDOWN
	}
	atomic.AddInt32(&d.clientCounter, -1)
	return
}

func (d *Rocks) Get(key []byte) (data []byte, err error) {
	if newClientCounter := atomic.AddInt32(&d.clientCounter, 1); newClientCounter > 0 {
		data, err = d.db.GetBytes(d.ro, key)
	} else {
		err = ERR_SHUTDOWN
	}
	atomic.AddInt32(&d.clientCounter, -1)
	return
}

func (d *Rocks) Delete(k []byte) (err error) {
	// println("del", string(k))
	if newClientCounter := atomic.AddInt32(&d.clientCounter, 1); newClientCounter > 0 {
		err = d.db.Delete(d.wo, k)
	} else {
		err = ERR_SHUTDOWN
	}
	atomic.AddInt32(&d.clientCounter, -1)
	return
}

func (d *Rocks) Destroy() {
	os.RemoveAll(d.path)
}

func (d *Rocks) Close() {
	for {
		swapped := atomic.CompareAndSwapInt32(&d.clientCounter, 0, -100)
		if swapped {
			break
		}
		log.Printf("waiting to close db %s ...", d.path)
		time.Sleep(300 * time.Millisecond)
	}
	d.wo.Destroy()
	d.ro.Destroy()
	d.dbOptions.Destroy()
	d.db.Close()
	log.Printf("closed db %s", d.path)
}

func (d *Rocks) Size() (sum uint64) {
	for _, lifeFileMetadata := range d.db.GetLiveFilesMetaData() {
		sum += uint64(lifeFileMetadata.Size)
	}
	return sum
}

func (d *Rocks) GetLiveFilesMetaData() ([]gorocksdb.LiveFileMetadata) {
	return d.db.GetLiveFilesMetaData()
}
