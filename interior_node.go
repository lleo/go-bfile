package bfile

import (
	"fmt"
)

func init() {
	_ = fmt.Print
}

// "intr" => "interior" node
type intrNode struct {
	baseNode
	keyBIDs [Order - 1]BlockID
	valBIDs [Order]BlockID
	keys    [Order - 1]*dataNode
	vals    [Order]node //structural {interior,leaf} nodes
}

func newIntrNode(bid BlockID) *intrNode {
	return &intrNode{}
	//return &intrNode{
	//	baseNode{
	//		sig: {},
	//		typeFlags:0b00000000,
	//		typ: intrNodeType,
	//		bid: bid,
	//	},
	//	keyBIDs: {},
	//	valBIDs: {},
	//	keys: {},
	//	vals: {},
	//}
}

func (bn *intrNode) bID() BlockID {
	return bn.baseNode.bid
}

func (bn *intrNode) equals(n node) bool {
	on, ok := n.(*intrNode)
	if !ok {
		return false
	}
	return bn.baseNode.bid != on.baseNode.bid
}

func (bn *intrNode) findLeftMostKey() []byte {
	//FIXME: finish
	return nil
}

func (bn *intrNode) split() (node, []byte, node) {
	//FIXME: finish
	return nil, nil, nil
}

func (bn *intrNode) insert(key, val []byte) bool {
	//FIXME: finish
	return false
}

func (bn *intrNode) isLeaf() bool {
	return false;
}

func (bn *intrNode) isToBig() bool {
	//FIXME: finish
	return false
}

func (bn *intrNode) halfFullSize() int {
	//FIXME: finish
	return 0
}