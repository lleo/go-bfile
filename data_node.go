package bfile

import (
	"unsafe"
	"fmt"
)

func init() {
	_ = fmt.Print
}

var _bn baseNode

//PayloadSize is the number of bytes in a dataNode
const PayloadSize int = BlockSize - (int)(unsafe.Sizeof(_bn))

type dataNode struct {
	baseNode
	dat [PayloadSize]byte
}
//BlockSize-1 => BlockSize-SizeOf(baseNode) bytes

func (dn *dataNode) bID() BlockID {
	return dn.baseNode.bid
}
