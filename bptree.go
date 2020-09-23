package bfile

import (
	"fmt"
)

func init() {
	_ = fmt.Print
}

//BpTree interface is unnecessary
//type BpTree interface {
//	findLeaf(key []byte) (*leafNode, bool)
//	getData(bid BlockID) ([]byte, bool)
//	//...
//}

//node is the tree structural interface for branch {interior,leaf} nodes.
//
type node interface {
	bID() BlockID
	equals(node) bool
	findLeftMostKey() []byte
	split() (node, []byte, node) //leftNode, key, rightNode
	isLeaf() bool
	isToBig() bool
	//halfFullSize() int
	String() string
}

const (
	dataNodeIdx  uint8 = iota               // iota=0
	intrNodeIdx  uint8 = 1 << iota          // iota=1
	leafNodeIdx  uint8 = 1 << iota          // iota=2
	dataNodeMask uint8 = ^(128 + dataNodeIdx) //0xfe 0b00000001
	intrNodeMask uint8 = ^(128 + intrNodeIdx) //0xfd 0b00000010
	leafNodeMask uint8 = ^(128 + leafNodeIdx) //0x
	spanMask     uint8 = ^dataNodeMask
	spanShift    uint8 = 1
)

type bpTree struct {
	root    node
	nents   int
	depth   int
	order int
}

func createBpTree() *bpTree {
	return &bpTree{
		root:    newLeafNode(incBID()),
		nents:   -1,
		depth:   -1,
		order:   Order,
	}
}

func (t *bpTree) findLeaf(key []byte) *leafNode {
	return nil
}

func (t *bpTree) Get(key []byte) ([]byte, bool) {
	return nil, false
}

func (t *bpTree) Put(key []byte) (bool, error) {
	return false, nil
}