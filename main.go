/*
package main is the hashsum tool for calculating hash sums of arbitrary files.
This is a command line tool and is tested under windows and linux x86_64 systems.
It tries to mimick the md5sum, sha1sum, sha256sum and sha512 from GNU coreutils
using only the Go Standard Library (and github/stretchr/testify/assert for tests)

Usage:
hashsum [-a md5|sha1|sha256|sha512] [-o <outfile>] <file> [<file> [...]]
	calculate hashsums over the specified files, optionally writing it to <outfile>
hashsum [-a md5|sha1|sha256|sha512] -r <hashstring> <file>
	calculate hashsum of specified file and compare it to the provided hashsum

NOT IMPLEMENTED YET:
hashsum [-a md5|sha1|sha256|sha512] -c <sumfile> [<file>, [<file>[, ...]]]
  compares priviously generated hashsum files with present files. Usually
	checks every entry of the sumfile, but can be limited to files of interest
	by specifying them as arguments
*/
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

// Filemap is a map of filenames to their hashsum values
type Filemap map[string][]byte

func main() {
	var algorithm string
	flag.StringVar(&algorithm, "a", "sha256", "md5, sha1, sha256, sha512")
	// var check string
	// flag.StringVar(&check, "c", "", "file with computed hashes")
	var reference string
	flag.StringVar(&reference, "r", "", "hash to compare file to (single file only)")
	var output string
	flag.StringVar(&output, "o", "", "output file for computed hashes")
	flag.Parse()

	files := flag.Args()

	// No target provided
	if len(files) < 1 {
		fmt.Printf("No file(s) specified. Use %s -h to get usage information.\n", os.Args[0])
		os.Exit(0)
	}

	// Provided r Argument as hash to compare the file to
	if len(reference) > 1 {
		if len(files) > 1 {
			fmt.Println("Multiple files provided. Selecting first file for hash comparison.")
		}
		files = files[0:1]
		hashes := CalculateHashes(files, algorithm)
		file := files[0]
		hash := fmt.Sprintf("%x", hashes[file])
		if hash == reference {
			fmt.Printf("%s  ok\n", file)
			os.Exit(0)
		} else {
			fmt.Printf("%s  mismatching hashes\n%s  (expected)\n", file, reference)
		}
	}

	hashes := CalculateHashes(files, algorithm)

	printHashes(hashes)

	if len(output) > 1 {
		postfix := ""
		if len(hashes) != 1 {
			postfix = "es"
		}
		fmt.Printf("Write %d filehash%s to '%s' \n", len(hashes), postfix, output)
		WriteHashes(hashes, output)
	}

}

// func readHashfile(path string) Filemap {
// 	fmap := make(Filemap)
// 	return fmap
// }

// printHashes outputs a list of '<hash>  <filename>' rows to the command line
func printHashes(hashes Filemap) {
	for name, hash := range hashes {
		fmt.Printf("%x  %s\n", hash, name)
	}
}

// WriteHashes saves the calculated hashmap as a list of '<hash>  <filename>'
// rows to the file specified by the filepath, ignoring probably existing files.
func WriteHashes(hashes Filemap, filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	for name, hash := range hashes {
		fmt.Fprintf(f, "%x  %s\n", hash, name)
	}

	return nil
}

// Check, if a path is a directory. Return true if it is and false otherwise
func isDir(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

// CalculateHashes for the provided list of files using the algorithm
// defined by algorithm. Returns filenames mapped to their hashsum.
// This is a wrapper selecting the hash algorithm. Processing takes place
// in createFilemap(files, hash) function.
func CalculateHashes(files []string, algorithm string) Filemap {
	hashes := make(Filemap)
	switch algorithm {
	case "sha512":
		hashes = createFilemap(files, sha512.New())
	case "sha256":
		hashes = createFilemap(files, sha256.New())
	case "sha1":
		hashes = createFilemap(files, sha1.New())
	case "md5":
		hashes = createFilemap(files, md5.New())
	default:
		fmt.Printf("Unknown hash algorithm. Supports only md5, sha1, sha256 and sha512.\n")
	}
	return hashes
}

// createFilemap goes through the provided lists of files and calculates
// a hashsum on them. The algorithm is injected as hash.Hash instance.
// Returns filenames mapped to their hashsum.
// If a provided filepath is a directory it will log this to command line but
// will skip the entry to the file->hashsum map. This is also the case, if a
// file cannot be read.
// As of now, there is no parallelism used when processing the files.
func createFilemap(files []string, hash hash.Hash) Filemap {
	fmap := make(Filemap)
	for _, file := range files {

		if isDir(file) {
			fmt.Printf("'%s' is a directory (not hashable).\n", file)
			continue
		}

		f, err := os.Open(file)
		if err != nil {
			log.Fatal("Error: ", err)
			continue
		}
		defer f.Close()

		if _, err := io.Copy(hash, f); err != nil {
			log.Fatal("Error: ", err)
			continue
		}

		fmap[file] = hash.Sum(nil)
	}
	return fmap
}
