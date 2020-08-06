package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var input string
var output string
var offset int64
var limit int64

func init() {
	flag.StringVar(&input, "from", "", "source file")
	flag.StringVar(&output, "to", "", "destination")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
	flag.Int64Var(&limit, "limit", 0, "limit to copy")
}

func checkFlags() {
	if input == "" || output == "" {
		flag.Usage()
	}
}

func printAndExitIfError(err error, preAction func()) {
	if err != nil {
		preAction()
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func openFiles(inFilePath string, outFilePath string) (*os.File, *os.File) {
	inFile, err := os.Open(inFilePath)
	printAndExitIfError(err, func(){})
	outFile, err := os.Create(outFilePath)
	printAndExitIfError(err, func(){ _ = inFile.Close()})
	return inFile, outFile
}

func closeFiles(files... *os.File) {
	for _, file := range files {
		_ = file.Close()
	}
}

func processLimitFlag(inFile *os.File, closeFunc func()) {
	statFile, err := inFile.Stat()
	if err == nil {
		if limit == 0 || limit + offset > statFile.Size() {
			limit = statFile.Size() - offset
		}
	} else {
		if limit == 0 {
			printAndExitIfError(errors.New("limit should be greater then zero"), closeFunc)
		}
	}
}

func processOffsetFlag(inFile *os.File, closeFunc func()) {
	_, err := inFile.Seek(offset, io.SeekStart)
	printAndExitIfError(err, closeFunc)
}

func copyStream(in io.Reader, out io.Writer, limit int64) error {
	bufIn := bufio.NewReader(in)
	bufOut := bufio.NewWriter(out)
	var readLen int64
	var tempSymbol byte
	var err error
	for readLen < limit {
		tempSymbol, err = bufIn.ReadByte()
		readLen++
		if err != nil {
			return err
		}
		err = bufOut.WriteByte(tempSymbol)
		if err != nil {
			return err
		}
		fmt.Printf("Progress: %.2f%%\r", (float32(readLen) / float32(limit)) * 100)
	}
	return bufOut.Flush()
}

func main() {
	flag.Parse()
	checkFlags()

	inFile, outFile := openFiles(input, output)
	closeFunc := func() {closeFiles(inFile, outFile)}

	processLimitFlag(inFile, closeFunc)
	processOffsetFlag(inFile, closeFunc)

	err := copyStream(inFile, outFile, limit)
	closeFunc()
	printAndExitIfError(err, func(){})
}