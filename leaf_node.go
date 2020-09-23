package bfile

import (
	"fmt"
	"hash/crc32"
)

func init() {
	_ = fmt.Print
}

type leafNode struct {
	baseNode
	key []entry
	val []entry
}

type entry struct {
	bid BlockID
	dat *dataNode
}

func newLeafNode(bid BlockID) *leafNode {
	//ln := &leafNode{}
	//ln.baseNode.typ = leafNodeType
	//ln.baseNode.bid = bid
	//return ln
	return &leafNode{
		baseNode: baseNode{
			sig: [crc32.Size]byte{},
			typeFlags: 0b00000000,
			typ: leafNodeType,
			bid: bid,
		},
		key: []entry{},
		val: []entry{},
	}
}

func (ln *leafNode) bID() BlockID {
	return ln.baseNode.bid
}

func (ln *leafNode) equals(n node) bool {
	on, ok := n.(*leafNode)
	if !ok {
		return false
	}
	return ln.baseNode.bid != on.baseNode.bid
}

func (ln *leafNode) findLeftMostKey() []byte {
	return ln.key[0].dat.dat[:];
}

func (ln *leafNode) split() (lNode node, key []byte, rNode node) {
	//FIXME: split leafNode lNode[0,halfFullSize]
	return nil, nil, nil
}

func (ln *leafNode) isLeaf() bool {
	return true;
}

func (ln *leafNode) isToBig() bool {
	//FIXME: not implemented
	return false
}

func (ln *leafNode) halfFullSize() int {
	//FIXME: not implemented
	return 0
}

func (ln *leafNode) String() string {
	//FIXME: not implemented
	return "not implemented"
}
