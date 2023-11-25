// dhash recursively hashes all of the files in a directory and produces
// a cumulative hash string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

const (
	appName = "dhash"

	usage = appName + `

SYNOPSIS
  ` + appName + ` [options] <directory-path>

DESCRIPTION
  ` + appName + ` recursively hashes all of the files in a directory and produces
  a cumulative hash string.

  For more information regarding the go.mod hash type, please refer to:
  https://pkg.go.dev/golang.org/x/mod/sumdb/dirhash

OPTIONS
`

	helpArg     = "h"
	hashTypeArg = "t"
	eachFileArg = "e"
	outputArg   = "o"

	goModHashType  = "go.mod"
	sha256HashType = "sha256"
	sha512HashType = "sha512"
)

func main() {
	log.SetFlags(0)

	err := mainWithError()
	if err != nil {
		log.Fatalln("fatal:", err)
	}
}

func mainWithError() error {
	help := flag.Bool(
		helpArg,
		false,
		"Display this information")

	hashType := flag.String(
		hashTypeArg,
		sha256HashType,
		fmt.Sprintf("The hash to use. Supported types are: %q, %q, %q\n",
			goModHashType, sha256HashType, sha512HashType))

	eachFile := flag.Bool(
		eachFileArg,
		false,
		"Write each individual file's hash to output")

	outputPath := flag.String(
		outputArg,
		"-",
		"The output file path (specify '-' for stdout)")

	flag.Parse()

	if *help {
		_, _ = os.Stderr.WriteString(usage)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if flag.NArg() > 1 {
		return errors.New("only one directory path should be supplied")
	}

	dirPath := flag.Arg(0)
	if dirPath == "" {
		return errors.New("please specify a directory path to hash as the last argument")
	}

	var optNewHashFn func() hash.Hash

	switch *hashType {
	case goModHashType:
		// OK.
	case sha256HashType:
		optNewHashFn = sha256.New
	case sha512HashType:
		optNewHashFn = sha512.New
	default:
		return fmt.Errorf("unsupported hash type: %q", *hashType)
	}

	output := os.Stdout
	if *outputPath != "-" {
		var err error

		output, err = os.OpenFile(*outputPath, os.O_CREATE|os.O_WRONLY, 0o600)
		if err != nil {
			return fmt.Errorf("failed to open output file - %w", err)
		}
		defer output.Close()
	}

	filePaths, err := dirhash.DirFiles(dirPath, "")
	if err != nil {
		return fmt.Errorf("failed to find files in directory %q", dirPath)
	}

	var optEachFileOutputFn func(filePath string, h hash.Hash) error
	if *eachFile {
		if *hashType == goModHashType {
			optNewHashFn = sha256.New
		}

		optEachFileOutputFn = func(filePath string, h hash.Hash) error {
			_, err := output.WriteString(filePath + " - " + hashToString(h) + "\n")
			return err
		}
	}

	opener := newFileOpener(optNewHashFn, optEachFileOutputFn)

	goModHash, err := dirhash.DefaultHash(filePaths, opener.Open)
	if err != nil {
		return fmt.Errorf("failed to dir hash %q - %w", dirPath, err)
	}

	if *hashType == goModHashType {
		_, err := output.WriteString(goModHash + "\n")
		return err
	}

	_, err = output.WriteString(hashToString(opener.optCumulativeHash) + "\n")
	return err
}

func newFileOpener(optNewHashFn func() hash.Hash, optEachFileOutputFn func(filePath string, h hash.Hash) error) *fileOpener {
	var optCumulativeHash hash.Hash
	if optNewHashFn != nil {
		optCumulativeHash = optNewHashFn()
	}

	return &fileOpener{
		optNewHashFn:       optNewHashFn,
		optCumulativeHash:  optCumulativeHash,
		optOutputHashStrFn: optEachFileOutputFn,
	}
}

type fileOpener struct {
	optNewHashFn       func() hash.Hash
	optCumulativeHash  hash.Hash
	optOutputHashStrFn func(filePath string, h hash.Hash) error
}

func (o *fileOpener) Open(filePath string) (io.ReadCloser, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var optFileHash hash.Hash
	if o.optOutputHashStrFn != nil {
		optFileHash = o.optNewHashFn()
	}

	return &fileHasher{
		file:   f,
		hash:   optFileHash,
		parent: o,
	}, nil
}

type fileHasher struct {
	file   *os.File
	hash   hash.Hash
	parent *fileOpener
}

func (o *fileHasher) Read(b []byte) (int, error) {
	i, err := o.file.Read(b)
	if err != nil {
		return i, err
	}

	if o.hash != nil {
		_, err := o.hash.Write(b[0:i])
		if err != nil {
			return 0, fmt.Errorf("failed to write bytes to file hash - %w", err)
		}
	}

	if o.parent.optCumulativeHash != nil {
		_, err := o.parent.optCumulativeHash.Write(b[0:i])
		if err != nil {
			return 0, fmt.Errorf("failed to write bytes to cumulative hash - %w", err)
		}
	}

	return i, nil
}

func (o *fileHasher) Close() error {
	err := o.file.Close()

	if o.hash != nil {
		if o.parent.optOutputHashStrFn == nil {
			panic("file hash object is non-nil, but output function is nil")
		}

		err := o.parent.optOutputHashStrFn(o.file.Name(), o.hash)
		if err != nil {
			return err
		}
	}

	return err
}

func hashToString(h hash.Hash) string {
	return hex.EncodeToString(h.Sum(nil))
}
