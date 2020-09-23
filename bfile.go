package bfile

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func init() {
	_ = fmt.Print
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

//HeaderBlockSize is the number of bytes the two header blocks are.
const HeaderBlockSize = 2048

//BlockSize is the number of bytes blocks are a multiple of.
const BlockSize = 256

//Order is the base number of keys and vals in branch {interior,leaf} nodes;
//calculated from the number of 7byte values that fit in BlockSize Block.
//Order = Floor(((BlockSize - SizeOf(baseNode)) / ((SizeOf(key) + SizeOf(val)))
//Order = Floor((256 - (4 + 1 + 1 + 8)) / (SizeOf(BlockID) + SizeOf(BlockID)))
//Order = Floor(255 / (2 * SizeOf(BlockID)))
//Order = Floor(255 / (2 * 7))
//Order = Floor(18.2142857..)
const Order = 18

//BlockID FIXME:
// BlockID = 0 //unassigned (never used) BlockID
// BlockID = 1 //first valid BlockID
type BlockID uint64 //only the first 56bits used

var nextBID BlockID = 1

func incBID(int span) BlockID {
	bid := nextBID
	nextBID += span
	return bid
}

func offset(bid BlockID) uint64 {
	return (2 * HeaderBlockSize) + ((uint64(bid) - 1) * BlockSize)
}

//Header is the parsed contents of the header block.
type Header struct {
	root         BlockID
	creator      string
	createdOn    time.Time
	lastModified time.Time
	comment      string
}

func createHeader(rootBID BlockID, timeOfCreation time.Time, creator, comment string) *Header {
	return &Header{
		creator:      creator,
		createdOn:    timeOfCreation,
		lastModified: timeOfCreation,
		comment:      comment,
	}
}

//Marshal header data
func (h *Header) Marshal(rootBid BlockID) []byte {
	var b strings.Builder
	fmt.Fprintf(&b, "RootBID: %d\n", rootBid)
	fmt.Fprintf(&b, "Creator: %s\n", h.creator)
	createdOnStr, err := h.createdOn.MarshalText()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(&b, "CreatedOn: %s\n", createdOnStr)
	lastModifiedStr, err := h.lastModified.MarshalText()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(&b, "LastModified: %s\n", lastModifiedStr)
	fmt.Fprintf(&b, "Comment: %s\n", h.comment)

	s := b.String()
	bs := []byte(s)
	if len(bs) > HeaderBlockSize {
		log.Fatalf("len(bs),%d > HeaderBlockSize,%d", len(bs), HeaderBlockSize)
	}

	return bs
}

//BFile struct contains all the essential elements of a transactional file-backed
// B+Tree data structure.
type BFile struct {
	hdr   *Header
	fname string
	file  *os.File
	bpt   *bpTree
	txn   *Txn //current open transaction
}

//Create a b+tree backed file.
func Create(name string) (*BFile, error) {
	f, err := os.Create(name)

	if err != nil {
		log.Fatal(errors.Wrapf(err, "Failed to Create(%q): ", name))
		//TERM: log.Fatal -> os.Exit(1)
	}
	bid := incBID()
	bf := &BFile{
		hdr: createHeader(bid, time.Now(), "Sean M. Egan", "First Version.\n"),
		fname: name,
		file: f,
		bpt: &bpTree{
			root:  newLeafNode(bid),
			order: Order,
			nents: 0,
			depth: 0,
		},
		txn: nil,
	}
	//create *File
	bf.file, err = os.Create(name)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed trying to create file named %s", name)
	}

	return bf, nil
}

//Open a b+tree backed file.
func Open(name string) (*BFile, error) {
	//FIXME: finish
	return nil, nil
}

//Txn will create a  transaction object for BFile.
//
func (bf *BFile) Txn() (*Txn, error) {
	//FIXME: not implemented
	return nil, nil
}

//StartTxn will create a transaction and immediately start it.
//
//func (bf *BFile) StartTxn() (*Txn, error) {
//	txn, err := bf.Txn()
//	if 
//}
func (bf *BFile) StartTxn() (*Txn, error) {
	txn, err := bf.Txn()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Start()
	if err != nil {
		log.Fatal(err)
	}
	
}


func (bf *BFile) writeHeaders() error {
	err := bf.writeHeader(0)
	if err != nil {
		return errors.Wrapf(err, "failed bf.writeHeader(%d)", 0)
	}
	err = bf.writeHeader(HeaderBlockSize)
	if err != nil {
		return errors.Wrapf(err, "failed bf.writeHeader(%d)", HeaderBlockSize)
	}
	return nil
}

func (bf *BFile) writeHeader(off int64) error {
	bs := bf.hdr.Marshal(bf.bpt.root.bID())
	nw, err := bf.file.WriteAt(bs, off)
	if err != nil {
		log.Println(err)
		return errors.Wrapf(err, "failed bf.file.WriteAt(bs, %d", off)
	}
	if nw != len(bs) {
		return fmt.Errorf("nw,%d != len(bs),%d", nw, len(bs))
	}
	return nil
}

//Dump outputs a string representing the contents of BFile.
func (bf *BFile) Dump() string {
	var b strings.Builder
	b.WriteString("Building...\n")
	b.WriteString("done.\n")
	return b.String()
}
