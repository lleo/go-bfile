package bfile

import (
	"fmt"
)

func init() {
	_ = fmt.Print
}

type WriteOp struct {
	bid  BlockID
	span int
	blk  []byte
}

type FreeOp struct {
	bid  BlockID
	span int
}

//newWriteOp is place holder for potential slab allocator for WriteOps
func newWriteOp(bid BlockID, span int, blk []byte) *WriteOp {
	//FIXME: make sure it doesn't conflict with existing BlockID/Spans
	//FIXME: make sure it doesn't conflict with other allocations
	return &WriteOp{
		bid:  bid,
		span: span,
		blk:  blk,
	}
}

//freeWriteOp is place holder for potential slab allocator for WriteOps
func freeWriteOp(op *WriteOp) {
	//noop
	return
}

//newFreeOp is place holder for potential slab allocator for FreeOps
func newFreeOp(bid BlockID, span int) *FreeOp {
	//FIXME: same stuff as newWriteOP?
	return &FreeOp{
		bid:  bid,
		span: span,
	}
}

//freeFreeOp is place holder for potential slab allocator for FreeOps
func freeFreeOp(op *FreeOp) {
	//noop
	return
}
