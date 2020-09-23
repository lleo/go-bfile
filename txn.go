package bfile

import (
	"sync"
)

//Txn is a struct defining all transactional operations to bfile.
type Txn struct {
	rwLock        sync.RWMutex
	bpt           *bpTree
	writeOpsOrder []*WriteOp
	writeOpsByBid map[BlockID]*WriteOp //spammed hash of BlockIDs
	freeOpsOrder  []*FreeOp
	freeOpsByBid  map[BlockID]*FreeOp //spammed hash of BlockIDs
}

//newTxn properly allocates a *Txn
//It also is a place holder for potential slab allocator of Txn structs
func newTxn() *Txn {
	return &Txn{
		rwLock:        sync.RWMutex{},
		bpt:           bpTree{},
		writeOpsOrder: make([]*WriteOp, 8),
		writeOpsByBid: make(map[BlockID]*WriteOp),
		freeOpsOrder:  make([]*FreeOp, 8),
		freeOpsByBid:  make(map[BlockID]*FreeOp),
	}
}

func (t *Txn) Start() error {
	t.rwLock.
}

//freeTxn is a place holder for potential slab allocator of Txn structs
func freeTxn(txn *Txn) {
	//noop
	return
}

//Get ...
func (t *Txn) Get(key []byte) (val []byte, err error) {
	return nil, nil
}

//Put ...
func (t *Txn) Put(key, val []byte) ([]byte, error) {
	return nil, nil
}

//Del ...
func (t *Txn) Del(key []byte) ([]byte, error) {
	return nil, nil
}

//ReadBlockBytes ...
func (t *Txn) ReadBlockBytes(bid BlockID, span int) ([]byte, error) {
	return nil, nil
}

//WriteBlockBytes ...
func (t *Txn) WriteBlockBytes(bid BlockID, span int, blk []byte) {
	//FIXME:
}

//FreeBlock ...
func (t *Txn) FreeBlock(bid BlockID, span int) {
	//FIXME:
}
