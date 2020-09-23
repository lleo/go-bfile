//build-1.go example program.
//
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"path"
	
	bfile "github.com/lleo/go-bfile"
)

func init() {
	_ = fmt.Print
	_ = bfile.Create
}

func whereAmI() (string, string, int) {
	pc, funcName, lineNum, ok := runtime.Caller(1)
	//fmt.Printf("pc=%+v; funcName=%q; lineNum=%d; ok=%v;\n", pc, funcName, lineNum, ok)


	fnc := runtime.FuncForPC(pc)
	ent := fnc.Entry()
	funcName = fnc.Name()
	fmt.Printf("fnc=%+v; funcName=%s, ent=0x%x pc?=0x%x;\n", fnc, path.Base(funcName), ent, pc)


	fnFileName, fnLineNum := fnc.FileLine(pc)
	fmt.Printf("fnFileName=%q; lineNum=%d; fnLineNum=%d; ok=%v;\n", fnFileName, lineNum, fnLineNum, ok)

	return funcName, fnFileName, fnLineNum
}

func main() {
	xit := 0

	fmt.Println("hello world")

	fn := "test-0.bfile"
	var bf *bfile.BFile
	var err error
	if _, err = os.Stat(fn); os.IsNotExist(err) {
		bf, err = bfile.Create(fn)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		bf, err = bfile.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
	}
	txn, err := bf.Txn()
	if err != nil {
		log.Fatal(err)
	}

	//k := []byte{'k','e','y'}
	//v := []byte{'v','a','l','u','e'}
	k := []byte("key")
	v := []byte("value")

	prev, err := txn.Put(k, v)
	if err != nil {
		log.Fatal(err)
	}
	if prev != nil {
		log.Println("key/value was inserted")
	} else {
		log.Println("key/value was replaced")
	}

	log.Println(bf.Dump())
	log.Println("Success!")

	funcName, fileName, lineNum := whereAmI()

	log.Printf("%s[%d] inside> %s\n\n", path.Base(fileName), lineNum, funcName)

	os.Exit(xit)
}
