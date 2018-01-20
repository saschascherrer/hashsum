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

type filemap map[string][]byte

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
		writeHashes(hashes, output)
	}

}

func CalculateHashes(files []string, algorithm string) filemap {
	hashes := make(filemap)
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

func readHashfile(path string) filemap {
	fmap := make(filemap)
	return fmap
}

func printHashes(hashes filemap) {
	for name, hash := range hashes {
		fmt.Printf("%x  %s\n", hash, name)
	}
}

func writeHashes(hashes filemap, filepath string) error {
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

func isDir(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

func createFilemap(files []string, hash hash.Hash) filemap {
	fmap := make(filemap)
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
