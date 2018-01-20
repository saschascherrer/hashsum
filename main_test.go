package main

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const utdir = "test"

var helloMd5 = []byte{0x92, 0x67, 0x51, 0x1a, 0x22, 0xee, 0x72, 0xad, 0x81, 0xec, 0x3e, 0x0c, 0x67, 0x4b, 0x11, 0x5e}
var linuxMd5 = []byte{0xbe, 0x52, 0x0d, 0x1b, 0xd3, 0xf6, 0xf4, 0xdb, 0xd0, 0xeb, 0xad, 0xcd, 0x26, 0xb8, 0x31, 0xa1}
var helloSha1 = []byte{0x7d, 0x95, 0xa7, 0xab, 0x47, 0xed, 0x59, 0x49, 0x69, 0xfd, 0xac, 0xc1, 0x14, 0x79, 0x4e, 0x81, 0x33, 0xd2, 0xc7, 0x86}
var linuxSha1 = []byte{0x56, 0xef, 0x5d, 0x35, 0x83, 0x43, 0x83, 0x88, 0x06, 0x2f, 0x10, 0xf8, 0x68, 0x05, 0xe7, 0xc1, 0x00, 0xd1, 0x7d, 0x59}
var helloSha256 = []byte{0x81, 0x9f, 0xe4, 0xc7, 0x39, 0x55, 0xd9, 0xd0, 0x95, 0x57, 0xa3, 0x1f, 0x9e, 0x92,
	0x55, 0x89, 0x5d, 0x26, 0x4f, 0x94, 0xd3, 0x8b, 0x40, 0xe8, 0x73, 0x35, 0x5c, 0x52, 0x95, 0x6c, 0xb6, 0x47}
var linuxSha256 = []byte{0x7c, 0x29, 0xbc, 0x3f, 0xb9, 0x6f, 0x1e, 0x23, 0xd9, 0x8f, 0x66, 0x4e, 0x78, 0x6d,
	0xdd, 0xd5, 0x3a, 0x15, 0x99, 0xf5, 0x64, 0x31, 0xb9, 0xb7, 0xfd, 0xfb, 0xa4, 0x02, 0xa4, 0xb3, 0x70, 0x5c}
var helloSha512 = []byte{0x30, 0x63, 0x8e, 0x54, 0x4a, 0xcd, 0xf2, 0x66, 0xf0, 0x59, 0xfd, 0x0b, 0x23, 0xfa,
	0xb3, 0x67, 0xa6, 0x17, 0xa4, 0x4b, 0xbd, 0xd9, 0x03, 0xd9, 0x69, 0x7d, 0xdf, 0xf7, 0x38, 0xba, 0x91, 0x69,
	0x0e, 0x6a, 0x72, 0xbf, 0x57, 0x12, 0xc1, 0x1b, 0x1b, 0x22, 0xf8, 0x42, 0x43, 0xfa, 0xd3, 0x10, 0x9b, 0x8c,
	0x07, 0x76, 0xe7, 0x1d, 0x4b, 0x60, 0x91, 0x31, 0x22, 0x42, 0xba, 0x06, 0xf1, 0x75}
var linuxSha512 = []byte{0x85, 0x5b, 0xf8, 0x2a, 0xc9, 0xb1, 0x19, 0x98, 0xab, 0xb5, 0xec, 0xb7, 0x5b, 0x5d,
	0x71, 0x08, 0x81, 0xe8, 0x11, 0x4a, 0x5b, 0x86, 0x34, 0x0d, 0x33, 0x34, 0x90, 0x2a, 0x33, 0xda, 0xe2, 0x57,
	0x3a, 0xe8, 0xdf, 0x56, 0xb6, 0xc6, 0xba, 0xc7, 0x89, 0x9f, 0xbf, 0xc6, 0x93, 0xa5, 0x10, 0x06, 0x61, 0x65,
	0x78, 0x22, 0x23, 0x56, 0x37, 0x2d, 0xbe, 0x0d, 0x9c, 0xfc, 0xe1, 0xb6, 0x3b, 0x2d}

func TestIsDir(t *testing.T) {
	if !isDir(utdir) {
		t.Error("test is a directory")
	}

	if isDir(path.Join(utdir, "hello.txt")) {
		t.Error("hello.txt is a file")
	}
}

func TestCalculateHashesHello(t *testing.T) {
	files := []string{
		path.Join(utdir, "hello.txt"),
	}

	var got filemap
	var should filemap

	got = CalculateHashes(files, "md5")
	should = filemap{
		path.Join(utdir, "hello.txt"): helloMd5,
	}
	assert.Equal(t, should, got)

	got = CalculateHashes(files, "sha1")
	should = filemap{
		path.Join(utdir, "hello.txt"): helloSha1,
	}
	assert.Equal(t, should, got)

	got = CalculateHashes(files, "sha256")
	should = filemap{
		path.Join(utdir, "hello.txt"): helloSha256,
	}
	assert.Equal(t, should, got)

	got = CalculateHashes(files, "sha512")
	should = filemap{
		path.Join(utdir, "hello.txt"): helloSha512,
	}
	assert.Equal(t, should, got)
}

func TestCalculateHashesLinux(t *testing.T) {
	files := []string{
		path.Join(utdir, "linux-4.9.77.tar.xz"),
	}

	var got filemap
	var should filemap

	got = CalculateHashes(files, "md5")
	should = filemap{
		path.Join(utdir, "linux-4.9.77.tar.xz"): linuxMd5,
	}
	assert.Equal(t, should, got)

	got = CalculateHashes(files, "sha1")
	should = filemap{
		path.Join(utdir, "linux-4.9.77.tar.xz"): linuxSha1,
	}
	assert.Equal(t, should, got)

	got = CalculateHashes(files, "sha256")
	should = filemap{
		path.Join(utdir, "linux-4.9.77.tar.xz"): linuxSha256,
	}
	assert.Equal(t, should, got)

	got = CalculateHashes(files, "sha512")
	should = filemap{
		path.Join(utdir, "linux-4.9.77.tar.xz"): linuxSha512,
	}
	assert.Equal(t, should, got)
}
