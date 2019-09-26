// Package gencosmosdb Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// cosmosdb/collection.go
// cosmosdb/cosmosdb.go
// cosmosdb/database.go
// cosmosdb/document.go
// cosmosdb/template.go
package gencosmosdb

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _collectionGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xdd\x6e\xdb\xb8\x12\xbe\x96\x9e\x82\xd5\x45\x61\x27\xaa\x7c\x1f\xc0\x17\x45\xd2\xe6\x04\x3d\xe9\x09\xda\x9c\xc5\x02\x41\xd0\x32\xd4\xd8\xe6\x46\x26\xb5\x24\xd5\x26\x08\xf2\xee\x0b\x92\x92\xc5\x3f\xc5\x4e\x77\x37\x17\x85\xc5\xf9\xe1\xcc\x37\x33\x1f\xc9\xb6\x98\xdc\xe3\x35\x20\xc2\xe5\x96\xcb\xfa\x2e\xcf\xe9\xb6\xe5\x42\xa1\x59\x9e\x15\x0c\xd4\x62\xa3\x54\x5b\xe4\xf3\x3c\x5f\x2c\xd0\x29\x6f\x1a\x20\x8a\x72\x86\x04\xb4\x02\x24\x30\x25\x11\x46\x64\xb7\x9e\xab\xc7\x16\x5c\x3d\xa9\x44\x47\x14\x7a\xca\xb3\x8b\x33\x94\xfe\x93\x4a\x50\xb6\x4e\x49\xbe\xff\x21\x39\x3b\x29\x68\x5d\xf2\x2d\x55\xb0\x6d\xd5\x63\xf1\x3d\xcf\xbe\x80\xe4\x9d\x20\x10\x79\xdc\xeb\xe9\x9b\x08\x7d\x5d\xd3\x2d\x48\x85\xb7\x6d\x68\x41\x99\x9a\x88\x77\xf0\xa5\xa4\xef\xea\x2b\x34\xab\x5f\x4b\xf0\x9b\x84\x66\xe5\x3b\xfb\x70\x8d\x93\x16\x07\x38\x03\x85\xd7\xbe\xb3\x33\x4e\xba\xad\xa9\xd5\xab\x9d\xd5\x9c\x84\x69\x2a\x2e\xa0\xbe\x12\x9c\x40\xdd\x09\x90\xaf\x49\xb3\x15\x91\xbb\x6b\x41\xd7\x6b\x10\x51\x68\x07\xb8\x53\xbd\xa9\xef\xf0\xff\x12\xc4\x19\xac\x28\x83\xfa\x63\xc7\x4c\x17\xca\x03\x1d\x76\xf5\x2a\x70\x76\xca\xd9\xaa\xa1\xe4\x57\x90\x23\x83\xa9\xef\xf1\x82\xd5\xf0\x40\xd9\xfa\x8a\x37\x94\x3c\x3a\x76\x47\x93\x92\x61\x0c\x3c\xb9\xef\xf5\x0a\x0b\x45\x75\xaa\x9f\xe0\xd1\x8b\xe6\x68\x52\xd2\x7b\x6d\x1d\x79\x3a\x77\x3d\x6e\x4d\xa7\x55\xfa\xc8\x8e\x26\x25\xbd\x4f\x32\x21\xf7\xfd\x9f\x03\x97\x2d\x56\x14\x37\xda\x1f\x1d\xc1\x3c\x9a\x94\xf4\xfe\xd7\x81\xdc\xf3\xfb\x6c\xa8\x2a\x00\xd3\xa5\x2b\x86\x06\x24\x51\x6b\xa4\x96\xb4\x02\x8b\x91\xb8\xde\x77\x8a\x6f\xb1\xa2\xc4\x84\x70\xc7\x79\x13\xa3\x88\x07\x9d\x74\xb1\x2f\x79\x0d\x51\xcf\x04\x65\xd5\x3a\xa1\x35\x69\xba\x1a\xea\x2b\xac\x36\x12\xdd\xdc\xba\xdf\xa3\xb5\xa3\x13\x50\xc8\xc3\x7e\x73\x78\x98\x32\x1f\x60\x74\x6c\x42\x10\xad\x08\xb5\x58\x6d\x06\x08\x1d\xed\x11\x40\xf3\xe9\xcc\xcc\xae\xf3\xd4\x26\x01\x17\xd8\x58\x6b\x78\xf0\x10\x82\x74\x78\x5a\x2d\x51\x5c\xa7\xa4\x4e\x20\x67\x58\xe1\x6b\x2d\x18\x42\xe9\x37\xa8\xfb\x75\x3f\x9a\x4f\x94\xd5\xde\xb0\xf7\xda\xf7\x94\x05\x87\xc8\x95\x00\x42\xa5\x3e\xee\xfa\x53\x63\xc8\x70\x58\x4f\x84\xee\x56\x27\xc8\x60\x28\x8a\x83\xac\xa7\x1d\x20\xeb\x47\x17\xa1\x6a\xb7\xf3\x68\xc0\x3b\xbc\x77\x04\x80\xee\xa1\x1f\x05\x4f\xd9\xdf\x4d\x13\xe1\xcd\x6d\xbc\xa5\x4c\x63\xb7\x6b\xf9\x49\xec\x7e\x03\xe1\x22\xb7\xd3\xfc\x61\xd7\x13\xa9\x4c\xb2\x4f\x70\x27\xb1\x5a\x48\xec\xd4\xbc\x71\x9f\xf4\x32\xe6\x6b\xa7\x36\xfd\xe7\x23\xb0\x8d\x66\x37\xe1\xbe\x9f\x82\xd0\x36\x41\x96\xd1\x60\x24\xbc\x0d\xe7\xef\x7e\x6f\x83\x66\x02\xca\x88\x68\x3d\x08\x47\x9a\x35\x68\xd2\xb5\x45\x2e\x32\x1a\x11\x33\xc3\xe5\x07\xa4\xc2\xb9\x7a\x0e\xae\x92\xd2\xdd\x74\xbc\x49\xca\xf0\x2a\x29\x9d\x7d\x4e\x79\xb7\xbb\x9e\xf9\x17\xb5\xf1\xf4\xed\x98\x9a\xbe\x36\xfa\x54\x3c\x79\x3b\x74\x77\xbf\xb9\x3d\x72\xee\xb5\xbd\xc9\x70\xb3\x72\x14\xf7\x0c\xdf\x17\xcc\xd6\xe0\xe5\xec\x0d\x20\x12\x46\x1e\xcf\x61\x6f\x97\xc6\x20\x09\xc5\xab\x00\xf1\xdb\x3a\xe9\x21\x42\x27\x11\xde\xcd\x6d\xb4\x38\xd8\xc7\xda\x87\x20\xf5\x02\x57\x59\xa8\x26\x90\x7a\xf9\xe1\x11\xd2\xd2\x01\x2f\x8c\xd0\xe4\xa0\xa7\x44\x48\x6a\x87\xbc\x19\xa2\x8d\x0e\x7a\x1c\x44\x56\xf1\x2b\xe0\x12\x3f\x98\x63\x44\xd2\x1f\x30\x61\xb5\x75\x54\x02\x63\xca\xcc\xe9\xfe\x92\xb1\xa3\x32\x05\xe7\x95\x80\x15\x7d\x48\x20\x23\x68\x6d\x65\x01\xa6\x1b\xc1\xbb\xf5\xa6\xed\xd4\x47\x81\xed\xe4\x85\x96\x2a\x52\x09\xdf\x2b\x58\x75\xf2\x45\xb0\xa4\x51\x89\x9a\x3b\x7c\x31\xc5\xe7\x9e\x51\x09\xdb\xd8\x34\xe4\x48\x64\xa7\x0d\x05\xa6\x9c\x7e\x3c\xd2\xd7\x8d\x3b\x2c\xc1\x4a\xf2\xac\x1d\x4f\xf1\x88\x1d\x7b\x6b\xea\x3f\xb3\x11\xb1\xa6\x01\x45\x0e\xca\x4c\x81\x58\x61\x02\x86\x26\x04\x60\x05\x33\x87\xbb\xe6\xc8\xfd\x2a\x11\x08\xc1\xc5\x3c\xcf\xfe\x4b\xa5\x9a\xcd\x1d\x77\x17\x0a\x04\x56\x5c\x58\xd1\xfb\xa6\x99\xf9\xa6\x72\xb4\x3d\x07\x35\xb3\x19\x4c\x79\x3f\x83\x06\xc2\x38\x8c\x4c\xb7\x47\xdb\x60\x72\x50\x8c\x31\x89\x38\xdb\x26\x18\x66\x30\x8c\xcb\xa2\x53\x1a\x12\x74\x8b\x13\x16\x2e\xcf\x08\x67\x8a\xb2\x0e\x0f\xff\x9f\xa1\xcb\x94\xd5\x9c\x8d\x77\x03\xfd\x26\x88\x2a\xb7\x73\x1e\xd6\x8e\x0e\xb0\x06\xd5\x1b\x0d\xdc\xfa\x7d\x86\x07\x35\x89\xbb\xdd\xf3\x33\xfc\x8c\x7a\x40\x80\xea\x04\xd3\x3b\x33\xf8\x99\xe8\x9c\x55\xc7\x48\xca\x70\x46\xd0\x99\xd7\x9e\x25\xaa\xef\x68\x8d\x06\x94\xa3\x8d\x9e\xf2\xcc\xee\x85\xde\x86\xd0\x3d\xe5\x59\xe6\xf7\xfa\x09\x22\xd5\x2c\xe8\xff\x79\x99\x67\x66\x04\x4e\x9c\x51\x2b\xea\x3b\xb9\x28\xd0\xb1\xd9\xbc\xcc\xb3\x67\x9d\xaa\x89\x79\x46\x50\x54\xa3\x39\xc2\x4d\x33\xa3\x09\x2c\x27\x80\xd3\x51\xe3\xa6\xd1\x7e\x24\x3a\x59\xa2\xb7\x8e\xce\xd3\x73\x9e\x67\x2b\x2e\xb4\x4e\x66\x34\x8c\x95\x56\xa3\x95\xad\x46\x9e\x65\x74\x65\x16\xdf\x2c\x11\xa3\x8d\x51\x1d\x60\x60\xb4\x31\x06\x79\x96\x3d\x5b\x45\xbb\xcd\xd2\x51\xbd\x13\x80\xef\x8d\x42\x9e\xed\x02\xa9\xec\x89\x7e\xbc\x44\xce\xa7\x2b\x77\x4e\xa5\x41\x67\x5c\xf2\x1d\x8d\xd7\x96\x25\xc2\x6d\x0b\xac\x9e\xa5\xa4\x25\x8a\x96\xaa\xaa\x9a\xe7\x26\xb0\x3e\x9f\xc1\xae\xd4\xe1\xef\x29\x43\x4f\x36\x0c\x7e\x6a\x19\xf2\x07\x3a\x5c\xb2\xb0\x8e\x05\xd1\x5f\x4b\x44\xaa\x9a\xcf\x36\x4a\xb5\xd5\x25\xa8\x0d\xaf\xaf\xb8\x54\x25\x22\x95\xee\x90\xe3\x62\x61\x42\x29\x4a\x54\x0c\x3f\xac\xa4\x44\xc6\xc4\x72\xbd\x8d\xa2\x2e\xd1\xdb\x3e\x90\xd2\xf6\xa6\xc9\x60\x3e\x24\xb6\x27\x95\x49\x36\x4c\x77\xbc\x4b\x27\x4f\xa1\xb7\x13\x44\xf6\x35\xf0\x1e\x86\x75\x36\x25\x95\x6e\x76\x52\xd9\x00\xe7\x7b\xfc\x6a\x62\xd6\xab\xce\x08\xff\x62\x21\xce\x21\xaa\xc3\xa2\x38\xb6\xbe\xa3\x7a\x24\x34\x9c\x02\xfd\xef\x53\x69\xc7\xe4\xf5\x65\xe9\x8f\x91\xb8\xbd\x4c\xf8\x3a\xfa\x7e\xe0\x2a\x73\x49\x5a\x2e\x51\x51\x98\x99\xeb\xd1\xfb\x20\x84\x16\x7c\x81\x3f\x3b\x2a\xa0\xd6\x9d\x9e\x6d\x00\xd7\x20\x0c\x11\x98\x20\xff\x63\xbe\x9f\x46\x49\xf5\x15\xd4\xac\xb8\x58\xbd\xbb\xc4\x8a\x6c\x8a\x72\xdc\x60\xee\x94\xc5\x87\xcb\x06\x3a\x81\x47\x75\x71\xb6\x0f\x32\xa3\xe2\x60\xf6\x99\x9f\x72\xa6\x0c\x21\x1b\xe8\xcc\x3f\x7d\x7c\xfb\x9a\x60\x38\x60\xff\xed\xb9\x5c\x14\xc7\xfd\x16\x7b\x12\x74\xb5\xfe\xe9\xc1\x4d\x5c\x11\xc2\x01\x68\xef\x85\x44\x53\xd7\x85\xbf\x35\x08\xc7\xc5\xa2\xbd\xb7\x4f\x38\x4d\x52\xce\xef\xd7\xcd\x85\x8e\x70\x2a\x6b\xea\x66\xed\xf2\xce\x1c\x0d\x97\x05\x7b\xe2\x44\x44\xe2\xa4\x46\x57\x88\x56\xe6\x12\x33\x0e\x87\xa5\xfd\xe9\x69\x30\x36\xde\x6d\xe8\xcd\x6e\xbc\xbc\x49\xf9\xfd\xdd\xa5\x7c\x77\xea\x28\x16\x65\x60\xd9\x1f\x31\x16\x60\x9a\x02\x98\x4e\x32\x3e\x8d\x19\xdf\x27\x14\xe9\x0c\x46\xe2\xa0\x76\x93\x0d\xf2\x59\x0e\x86\xd5\x79\x3a\x0f\xed\xd0\xc2\xb6\x0c\xb1\x30\x54\x93\x3b\xe5\xfa\x2b\x00\x00\xff\xff\x54\x34\x4f\x68\x5b\x1b\x00\x00")

func collectionGoBytes() ([]byte, error) {
	return bindataRead(
		_collectionGo,
		"collection.go",
	)
}

func collectionGo() (*asset, error) {
	bytes, err := collectionGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "collection.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cosmosdbGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xdf\x4f\xdc\xba\x12\x7e\x4e\xfe\x8a\x69\x24\x50\x02\xc1\xbb\x70\x4b\x75\xb5\xbd\x79\x68\x61\xe9\x4f\x68\x6f\xd9\x3e\xb5\x47\xaa\x49\x26\x1b\x97\xc4\x4e\x6d\x07\xba\x45\xfc\xef\x47\x63\x7b\xd9\x85\xe5\x54\x95\x0e\x12\xda\x64\x6c\xcf\x7c\xf3\x7d\x33\xe3\xf4\xbc\xbc\xe4\x73\x84\x52\x99\x4e\x99\xea\x22\x8e\x45\xd7\x2b\x6d\x21\x8d\xa3\xe4\x62\x61\xd1\x24\x71\x94\x94\x7a\xd1\x5b\x35\x6a\x3a\x5e\xae\xbd\x9a\x86\x1f\x1c\x3e\x23\x03\xca\x52\x55\x42\xce\x47\x17\xdc\xe0\xb3\xa7\x64\xaa\x3b\x4b\x3f\x42\x8d\x84\x1a\xac\x68\xe9\x45\xa2\x1d\x35\xd6\xf6\xcb\x67\x8b\x3f\x6d\xaf\x95\x55\x4b\xc3\xa0\xdd\x3e\x63\xb5\x90\x73\x17\xd9\x8a\x0e\x93\x38\x8e\x92\xb9\xb0\xcd\x70\xc1\x4a\xd5\x8d\x86\xb9\xd2\xdf\xc5\x68\xae\x46\xa5\xaa\xb0\x4c\xe2\x2c\x8e\x47\x23\x98\x6a\xad\x34\x68\xec\x35\x1a\x94\xd6\x00\x97\x80\x64\x8b\xed\xa2\xc7\xb0\x6c\xac\x1e\x4a\x0b\x37\x71\x74\x6e\xb9\x1d\xcc\x91\xaa\x10\x84\xb4\x71\xe4\x9e\xfc\x9f\x8f\x0f\xdf\xbe\x1b\x25\x27\x09\x05\x49\xbe\xc5\xd1\x29\x1a\x43\x54\x6d\x6c\xe8\xfc\x42\xf2\x2d\xbe\x8d\xe3\x7a\x90\x25\xa4\x21\x5c\xe6\x7f\xd2\x6c\x79\xe2\x26\x8e\x34\xda\x41\x4b\xa8\x3b\xcb\xce\x7b\x2d\xa4\xad\xd3\x64\xab\x82\x2d\x33\x81\x2d\x93\xe4\x80\x6c\x85\x8c\xde\x96\xbf\x21\x7c\x46\x41\x46\x23\x78\x63\x9c\xeb\xb5\x2c\xbc\x63\x03\x56\x0f\x08\xa2\xa6\xdc\x41\x18\x50\x35\xac\xe5\xcf\x65\x05\xc2\x1a\x58\x9d\x23\x67\x1d\xb7\x65\x83\x06\xcc\xca\xea\xf2\xd8\x08\x92\x92\x53\x47\x6a\xbe\xb6\x99\x08\xcc\xe0\x42\xa9\x96\x12\xf4\xa1\x73\x50\x97\x30\x29\xe8\x91\xa5\x9e\x8c\xe7\x64\xba\x89\xa3\x25\x05\xb4\xb4\x86\xbf\x28\xd6\xe3\x47\xb7\x2b\xaa\x78\x6b\x30\xa4\x3d\xd5\x7a\x3a\xe3\xf3\x4f\xf8\x63\x10\x1a\x2b\x4a\xd0\x36\xe8\x21\x05\x06\xc8\x5a\x3b\x2b\xed\x84\x5a\x60\xeb\xf6\x49\x65\xa1\x57\xfd\xd0\x72\x8b\xa0\x24\x70\xf2\xf7\xf1\xf3\x0c\x94\x86\xe3\xe9\xfb\xe9\x6c\x0a\xaa\x47\xcd\xad\x50\x32\xbe\xe2\x7a\x23\x56\xe1\x44\x73\xb9\xd4\x69\xe2\x9c\x0b\x03\x3a\x2c\x27\xbe\x0a\x3f\xa1\xd5\x8b\x0f\xf2\xa3\xc6\x52\xc9\x4a\x90\xb3\x13\x2e\x5a\xac\x08\x9d\x16\x68\x80\x03\x71\x4b\x0b\x84\x53\x58\xa8\xb9\x68\x0d\x54\x03\x82\x55\x0e\xd3\xc6\x59\xaf\xc6\x3f\xba\x4e\x6b\xe7\x32\xcd\x3c\x0f\x19\xac\x64\xca\x88\xf0\x5a\x69\x10\x24\xc6\xf8\x39\x08\xf8\x1f\x1c\x3e\x07\xb1\xbb\xeb\xa4\xa0\x8d\x05\xd4\x69\x16\x47\x24\xdc\x93\x47\x15\xcf\x81\xba\x36\x68\xb5\x19\xde\xc5\x58\xaa\x1a\x47\xa4\x5c\x44\x7d\xcb\xce\x5b\xc4\x3e\x75\x8f\xc7\x83\x27\x36\xdd\x1f\x8f\x77\x44\x06\x3b\xe0\xcc\xa7\xa2\x6d\x85\x71\x0e\xb3\x35\xc9\x49\x6c\x52\xa0\x81\x02\xb6\x5d\x93\xb3\xb7\x46\xc9\xd7\x5c\x56\x2d\xde\xc4\xd1\x4b\x6e\x44\xe9\xdf\x26\xe0\xd7\xd7\x4c\x84\xe6\x18\xc9\xfc\xa1\xa7\x98\x66\xb9\xe7\x9e\xd1\x61\x76\xc9\xbe\xa9\xcf\xd4\x09\x55\xc9\xc4\xf5\x4e\x4e\x29\xe4\x31\xfd\xdf\xf5\x73\x09\x3b\x15\xb7\x9c\xc6\xdb\x51\x2b\x90\xca\x9d\x0f\xb6\x51\x5a\xfc\x42\xaa\x0f\x34\x36\xd5\xf8\x03\x76\x1c\x51\xc1\x92\x83\x46\xa3\x06\x5d\xe2\x6c\xd1\xe3\xea\xed\xbd\x90\x97\x61\x20\x38\xea\x2a\x2a\xc8\x49\xe1\x19\x39\x53\xd7\x69\xc6\x3e\xcf\x8e\xd2\x8c\x9d\x28\xdd\x71\x9b\x26\xa7\x4a\xe6\x30\x3e\x80\xb7\x5c\xc2\xc1\x78\xfc\x0c\xf6\x0f\x27\xe3\xa7\x93\xf1\x21\xbc\x3a\x9d\x51\xdd\x45\x0d\x9d\xa7\xf1\xcc\xce\xf0\x3a\xf5\x83\x99\x1e\x73\x28\x59\xc7\x8d\x45\xfd\x0e\x17\x59\x1c\x51\x01\x9f\x84\xa9\xd3\xe4\x90\x6c\x99\xaf\x72\xfd\xff\xab\x4c\xf2\x00\xcd\xb0\x99\x7a\xaf\xae\x51\x53\x62\xec\x14\x6d\xa3\xaa\xec\x77\x29\x6d\x1e\xa4\xc4\x32\x82\x47\x1e\x5e\x23\xaf\x50\xb3\x73\xb4\x69\xf2\x22\x70\xe7\x4a\x22\xc9\x61\xd0\x2d\xfb\xff\x80\x7a\x31\x35\x25\xef\x31\xbd\x37\x1c\x69\x7a\x15\x3e\x89\xed\x2b\xd4\xc5\x3e\x1b\x6f\x1b\x31\x2f\xdc\xb4\xf4\x17\x0e\x3b\xb7\xd5\x34\xdc\x41\xcc\x3d\xe0\x4c\x9d\x3b\x34\x69\xc3\xce\x87\x2e\x95\xa2\xcd\xe8\x6f\x13\xcb\xcf\xbd\xce\xec\x11\xd2\x24\x07\x07\xf8\xb7\xaa\x57\x2a\xed\x1c\x15\x39\xf4\xdc\x36\x7f\x20\x71\x0e\xf8\xb3\xc7\xd2\x62\x75\xff\xc6\xc9\x41\xc8\x1c\xd4\x60\xe9\x05\x75\xcd\x4b\xbc\xb9\xcd\xa1\x71\xd0\x8c\xef\x38\x8f\x33\xb4\xb5\xbf\x3c\x7e\xe4\x6e\xb0\x93\xdc\xb4\xe3\x0c\xaf\x97\xf5\xb7\xc4\x95\xd0\x82\x99\x8c\x46\xc9\x6e\xc9\x96\xf8\x5f\x94\xa5\x1a\xa4\xdd\x4d\x58\xa5\xca\xa1\xa3\xfb\x91\xf1\x5f\x83\x46\x77\xa3\x26\xbb\x3e\x19\x62\x69\x39\xc0\xe1\x49\x41\xef\x0f\x06\x36\x35\xa9\xdb\x21\xe4\xfa\x86\x8b\xa1\x26\x48\xdb\xee\x8b\x81\xbd\x1c\xea\x1a\xf5\xcd\x6d\x98\x2f\x93\x22\xf4\xdf\x19\x5e\x7b\x71\x74\x7a\x31\xd4\x39\x34\x59\x10\x2b\x15\x32\x0c\xa0\x07\x81\xef\x45\x76\x83\x85\xe4\x7b\xa9\xaa\x05\x14\xe0\x3f\x2c\xd8\x99\xea\x8f\x5a\x65\xbc\xd3\x2c\x6c\x59\x57\xf8\x48\x49\x8b\xd2\xee\x91\x44\x49\x0e\x09\xef\xfb\x56\x94\xae\xf6\x46\x74\x83\x27\x99\xcf\x8a\xc6\xe4\x65\x0e\x57\x04\x58\x73\x39\xc7\x3b\x31\x6e\xee\x79\xfd\x72\xf7\xdd\xc2\x8e\xb8\x54\x52\x94\xbc\x3d\x7d\x73\x3a\xf5\xab\xef\x70\x91\x5e\x66\x7f\x41\x01\x57\xde\xed\xa3\x05\x77\x85\xda\xf8\xda\x4f\x0e\xc6\xfb\xff\xdd\xdb\x3f\xd8\xfb\xcf\xbe\x6b\xe5\x92\x3d\x36\x59\x7e\x57\x68\xbe\xc3\x4c\x7f\x57\x19\x25\x6b\x4a\x76\xac\xe8\xdc\x1f\xc8\x19\x55\x58\x23\x5d\x9e\xa6\x77\xcc\x32\xc7\x66\x9a\x79\x9d\x97\x14\xac\x1d\x77\x3c\x3d\x4e\x52\x54\x61\x8b\x16\xd3\x60\xcb\xe1\x32\x0b\xba\x6d\x92\xeb\x02\x7a\x62\xfc\xd9\x70\xe8\xcb\x65\x20\x8f\xce\x11\x81\xd5\xbd\x0a\xf2\x43\x9c\xa6\x52\xc0\x4b\x75\xe4\xb1\x3a\xd3\x5a\x97\x3d\x29\x1e\xeb\x3d\x8a\x45\x97\x0b\x91\xe2\xe6\xbf\xaf\xbc\x35\x38\xec\xd5\x46\xd9\x64\xf4\x79\xb2\x59\x39\x21\xe9\x70\xb3\xa4\xdb\xa8\xf5\x32\xe1\x87\x5f\x37\x0f\xd1\x3d\xde\x56\x34\x0f\x02\xd5\xdb\xdb\xff\x02\x53\x70\xbd\x42\xa6\x06\x9b\x2d\xeb\xd1\x2d\x49\xd1\xc6\xb7\xf1\xdf\x01\x00\x00\xff\xff\x42\xf7\xa3\x62\x00\x0c\x00\x00")

func cosmosdbGoBytes() ([]byte, error) {
	return bindataRead(
		_cosmosdbGo,
		"cosmosdb.go",
	)
}

func cosmosdbGo() (*asset, error) {
	bytes, err := cosmosdbGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cosmosdb.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _databaseGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x96\xdf\x6f\xdb\x36\x10\xc7\x9f\xc5\xbf\xe2\xaa\x87\x40\x4a\x15\xf9\x65\xd8\x83\x01\x3d\x14\x76\x90\x05\x5d\xb2\xa1\xee\x80\x01\x45\xd1\x52\xd4\xd9\xe6\x2a\x93\x1e\x49\x2f\x0b\x8c\xfc\xef\x05\x49\xfd\x20\x65\xbb\x45\xfc\xe0\x44\xba\xe3\xdd\x97\x1f\x7e\x75\xf2\x9e\xb2\x6f\x74\x83\xc0\xa4\xde\x49\xdd\xd4\x84\xf0\xdd\x5e\x2a\x03\x19\x49\x52\x14\x4c\x36\x5c\x6c\x66\x35\xd5\xf8\xeb\x2f\x29\x49\x52\x81\x66\xb6\x35\x66\x9f\x92\x9c\x90\xd9\x0c\x96\xd4\x50\x1b\x05\x85\x7b\x85\x1a\x85\xd1\x40\xa1\xe9\xee\x12\xf3\xbc\xc7\x31\x47\x1b\x75\x60\x06\x8e\x24\xb9\x5f\xc2\xf0\xd1\x46\x71\xb1\x81\xaf\xff\x68\x29\xe6\x29\x6f\x0a\xb9\xe3\x06\x77\x7b\xf3\x9c\x7e\x25\xc9\x07\xd4\xf2\xa0\x18\xda\x15\x71\xe6\x17\x35\xcd\xfd\xc8\x77\xa8\x0d\xdd\xed\x01\x80\x0b\x63\x8b\xf7\xb9\x46\xc7\xa9\x2b\x6c\xd7\xe7\x05\x7c\xd1\xd8\xae\xe3\xe4\xdb\x8f\x74\x73\x21\x19\x0d\xdd\xc4\xc9\x0b\xd9\xb6\xc8\x0c\x97\x42\x4f\x93\x99\x6c\xdb\x89\x8e\xbf\x34\x2a\x7d\xbe\xf4\xc1\x86\xa2\xec\x97\x08\xb9\x0e\x99\xf7\xc4\x75\x8c\x5c\x07\xcc\x17\xf2\xe0\x99\x0c\x70\x20\x24\xc4\x6c\xf8\x22\xfb\x4e\x5b\xb4\xe2\x84\xff\xd8\x15\x3e\x7d\xbe\x1e\x8e\xbd\xcb\x1f\xa2\xd3\x3d\x39\xc5\xfd\x06\x16\x2d\x47\x61\x02\xd9\x5b\x06\xd1\xe7\xda\xda\xaf\xf4\x69\x24\xe9\x97\xbd\x63\x4e\x7f\xa7\x93\x24\x3b\xaa\x0d\xaa\xf7\xf8\xdc\xad\xfa\xf4\xb9\x7e\x36\x38\x21\xd8\xf5\xe2\xa1\x65\x81\xf9\xca\x11\xc6\x3e\x51\x18\x54\x6b\xca\xd0\xe1\x54\x48\x0d\x66\xc3\x36\x73\x18\xff\x2f\x00\x95\x92\x2a\x27\xc9\xef\x5c\x9b\x2c\x1f\x0a\xdd\x1b\x54\xd4\x48\xe5\x03\xef\xda\x36\x0b\x97\xe9\x71\xdd\x1d\x9a\xcc\xef\xe5\x7c\xdd\x25\xb6\x18\x77\x77\x91\x13\x9c\xb6\x4d\xdf\x34\x80\x7a\x1d\xe3\x26\x09\x93\xc2\x70\x71\xa0\xd6\xb7\x03\xc4\x46\x0a\x1c\xb8\xd7\x52\xb6\x13\x7e\x43\xe1\x98\x20\xef\x37\x19\x31\x1c\x93\x43\x8a\x8f\xf8\xbf\xb9\xc0\xc0\xf7\x7a\xc4\xa7\xc9\x29\x28\x34\x07\x25\x6c\x47\x81\x4f\x27\xe7\xb6\x3e\x08\x76\xba\x28\xdb\xb2\xc8\x38\x05\x4c\x8c\x53\xc0\xe8\x98\x81\x7b\x5c\xa3\x17\x66\x65\xff\x47\x95\xbd\xea\xa0\x93\x84\xc1\xbc\x82\xab\x18\xea\x91\x24\xc9\x96\xcd\x63\xf7\x6e\x59\x41\x92\xa9\x6b\xe7\x27\x6a\x48\xf2\x62\xab\x96\x83\x28\xd7\x1c\x2a\xf0\x93\xb8\x5c\x99\xe6\xb6\x1b\xce\xe5\x12\x99\x6c\x70\xe5\x44\x67\xc3\x82\x9c\x24\x7c\xed\x16\xbd\xa9\x40\xf0\xd6\xaa\x4e\x3c\x3a\x7b\xe9\xea\xf9\x2e\xdd\x4d\x56\xd8\xfb\x96\xba\x43\x98\x31\x98\x98\x24\x07\xda\xb6\x19\x3f\x39\xd0\xb3\xa7\x67\xdb\xd1\xb6\x6d\x6a\xed\xc8\x0c\xf1\xa3\xed\xb8\x96\xca\xc9\x69\x6a\x9f\x6f\x53\x78\xe9\xad\x40\x92\x33\xba\x4f\x85\x27\x2f\x3e\xd1\x36\xa8\x82\xc4\x5a\x21\xfd\xe6\xc2\x24\xe9\x04\x94\x7e\xf0\xbd\xad\x60\xb8\x18\x63\xc1\x88\xf3\xf1\xf1\xc6\x98\x34\xce\xb5\x0a\xe8\x7e\x8f\xa2\xc9\xa6\x91\x02\xa2\xcb\xb2\x2c\xf3\x88\xae\xcf\xff\x39\xe2\x6e\xaa\x08\x7c\x6a\x6a\x08\x67\x4b\x78\x59\x8c\xde\x73\xa0\xbd\x35\x58\xd9\xc8\xcc\x79\xfc\x01\xcd\x56\x36\x7f\x4a\x6d\x0a\x48\x9b\x5a\xa7\xe3\x9f\xb4\x00\x97\xb2\x32\xd4\x1c\xb4\xef\xd6\x14\x70\xe5\x1a\x16\x70\x65\xbf\x04\x6f\xf3\x5e\xf8\x0f\xc5\x5e\x18\x6e\x56\x52\xb7\xed\xab\x73\x73\xe8\x18\xd7\x99\x03\x7b\xf9\x69\x9b\x8b\xa3\x32\x68\xc6\x4a\x6b\x50\x56\x7a\x59\xf9\x0f\x6b\xda\xf9\xda\xd4\xbc\x19\x1f\xf6\xd7\x03\xbe\xc3\x8e\xef\x2c\x7d\x6b\x6b\x45\xb0\x87\x7b\x01\xee\x3f\xde\x17\xde\xc2\xaf\xc3\xdc\xcd\xfa\xd8\x10\x4e\x9c\xd5\xe6\x1e\x82\xd2\xfd\x44\xa9\x2a\x48\xd3\xf0\x41\xbf\x55\xca\x06\x3e\xe0\xbf\x07\xae\xb0\xb1\x8e\x4c\xb6\x48\x1b\xfb\x9b\x63\x5e\x79\x69\xbf\xb9\xeb\xe3\x18\x29\x57\x68\xb2\xf4\x7e\x7d\xf3\x40\x0d\xdb\xa6\x45\x5f\x3e\x0f\x40\xc7\x20\xbc\xc0\x60\xdf\xe5\xfd\xf2\x14\x86\xbb\x19\xd0\x78\x94\x0b\x29\x8c\x1b\xac\x0e\x8a\xfb\xea\x34\x04\x87\xc7\x47\x20\xa1\x89\x72\xe8\x5f\x1d\x76\x04\x4c\x7c\x11\x1c\x1d\x5f\x03\x2f\xdd\x8b\x6c\xe4\xe2\x9f\xcc\xcb\x20\xdc\x9a\xe8\x8d\xf8\x66\x20\x1b\x41\xfa\xfb\xe6\x41\xdf\x2c\x82\xc4\xb4\x98\xac\xec\xa6\x80\x37\x10\xbf\x64\xa0\x8b\x0f\x68\xe8\x18\x1d\xd0\xb9\x38\xdb\x7d\xbb\x89\xfa\xaa\x5f\x58\xde\x9d\x57\x6d\x0b\x7a\x48\xd5\x74\xe7\xce\x53\x24\xf0\xe9\xf7\x00\x00\x00\xff\xff\x2c\x98\x5c\xc0\x32\x0c\x00\x00")

func databaseGoBytes() ([]byte, error) {
	return bindataRead(
		_databaseGo,
		"database.go",
	)
}

func databaseGo() (*asset, error) {
	bytes, err := databaseGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "database.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _documentGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xcd\xaa\xc2\x40\x0c\x85\xd7\x9d\xa7\x08\x5d\x5f\xe8\xfe\x3e\x84\xe8\xc6\x8d\x08\x8d\x35\x94\xaa\x99\x19\x93\x54\x28\xe2\xbb\x4b\x7f\x9c\x32\x9d\xd5\x24\xe7\x7c\xe4\x8b\xd8\xdc\xb1\x25\x68\x82\x72\xd0\xeb\xc5\xb9\xaa\x82\x43\x4f\x32\x80\x50\x14\x52\xf2\xa6\x80\xf0\x1c\x57\xce\x86\x48\x4b\xaa\x26\x7d\x63\xf0\x76\xc5\x3c\x4f\x4f\x4d\x3a\xdf\xce\xff\xfa\xa6\xc1\xff\x97\x13\xf8\x17\xb8\x33\xe2\x68\x43\x59\xbb\x62\x8f\x82\x4c\x46\xa2\x70\x3a\xa7\xe1\x07\xc4\x94\x66\xd4\x67\x32\x5b\xdb\x99\x5d\x62\x66\xc3\xb5\xb5\x5a\xee\x90\x29\x09\x2e\xa7\x3c\x32\xe5\x6a\x47\x7c\xf4\xb4\x69\xbd\xc6\xdd\xc6\xe5\x1b\x00\x00\xff\xff\xa1\x09\x70\x18\x37\x01\x00\x00")

func documentGoBytes() ([]byte, error) {
	return bindataRead(
		_documentGo,
		"document.go",
	)
}

func documentGo() (*asset, error) {
	bytes, err := documentGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "document.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x49\x6f\xe3\x36\x14\x3e\x8b\xbf\x82\xa3\x43\x20\xc7\xb2\x74\x37\xe0\x43\xe1\x04\x69\xd0\x26\x9d\x66\x72\x28\x50\x14\x08\x4d\x3d\xdb\x1c\xcb\xa2\x86\xa2\x30\x0d\x0c\xff\xf7\x82\xa4\x16\x52\x5b\x92\x22\x4d\x5b\x1f\xa2\x88\x7c\xfb\xf7\x36\xe5\x84\x1e\xc8\x0e\x30\xe5\xc5\x91\x17\xc9\x06\x21\x76\xcc\xb9\x90\x38\x40\x9e\x9f\x81\x8c\xf7\x52\xe6\x3e\x42\x5e\x7e\xd8\x61\x7f\xc7\xe4\xbe\xdc\x44\x94\x1f\xe3\xaf\xec\xb8\x38\xb2\x4c\x82\x88\x77\x7c\x51\xb3\xc7\xf9\x61\x17\xef\x20\x6b\xde\x9b\x7f\x92\xf2\x78\x7c\xf6\xd1\x0c\x21\xf9\x9c\x03\x96\x70\xcc\x53\x22\x61\x9d\x32\xc8\x24\x2e\xa4\x28\xa9\xc4\x27\xe4\x5d\x26\x44\x92\x0d\x29\xaa\x1b\xe4\xe5\x44\xee\xd5\x3d\xcb\x76\xe8\x8c\x50\x1c\xe3\x47\x97\x97\x15\x98\x34\xf2\x30\x35\x6c\x5a\x49\x97\x50\x59\xbb\x25\x14\x94\x9e\xb5\x00\x22\x21\x30\x82\x43\x7c\x99\x1f\x76\x51\x4d\x3f\xc3\x81\xf3\x1e\x62\x10\x82\x8b\x19\xf2\x7e\x66\x85\x0c\x66\x8d\xe0\x5b\x09\x82\x48\x2e\xcc\xc5\x0f\x69\x1a\x74\x59\x8b\x96\xf7\x06\x64\xa3\xce\x3c\xc7\xf5\x3c\x40\x9e\x12\xfa\x66\xf3\xae\x20\x85\x71\xa7\x34\x15\xf2\x7e\x2d\x41\x3c\xb7\x34\xfa\x75\xc8\x25\x7d\xa1\x7c\xea\x92\x8e\x79\x78\xee\x60\xab\x62\x52\x8b\xb3\x11\x76\xb1\x47\x1e\xe5\x99\x64\x59\x49\x24\xe3\x59\x8d\xb4\x97\xf0\x0c\x70\xfd\xdb\x70\x9e\xf6\xc4\x6b\x6b\x5e\x23\x3f\x27\x42\x32\x25\xfc\x00\xcf\x8d\xfc\x6f\x8a\xbb\x96\x6f\x3c\x7b\x8b\x25\x56\x1a\x36\x26\xb8\x89\xc8\xea\x40\x3a\xa9\xd8\x12\xdb\xc9\x78\x0f\x7f\xca\x89\xd4\x31\xfa\xee\xe1\x7b\x27\xa1\x05\xc8\x52\x64\x4a\x6b\x06\xdf\x7b\x25\xb0\x2d\x33\xda\x67\x0a\x28\x4f\x53\x8a\xd7\x3c\x4d\x81\x2a\x3f\xcd\x71\x88\xd5\x39\x4b\x9a\xcc\xec\xa8\x3a\x21\xcf\x68\xc3\x17\x6e\x78\x4f\xc8\xf3\xdc\x92\x5d\x6a\x51\x34\x0a\x2e\x69\x47\xc9\x2c\x72\x29\x43\xe4\xe9\xf2\x5e\xe2\xf6\x37\xce\xab\x1b\xc1\x1c\xfb\xb1\xba\x2a\x62\x1f\xcf\x2b\x9b\x43\xe4\x9d\x55\x8c\xb4\xc3\x01\xc5\x9d\x04\x98\x61\x92\xa6\x01\xeb\x41\x30\x1a\x6f\xe5\x2c\x49\xd3\x5a\x4a\x81\x97\x2b\x7c\xe1\x50\x9e\xce\x08\x79\x5b\x2e\x14\xa5\x27\x1d\x7e\x45\xcc\x22\x03\x28\xf2\x3c\xb6\xd5\x87\x9f\x56\x38\x63\xa9\x26\xaf\xe3\x98\xb1\x54\x33\x20\xcf\x3b\x1b\xc2\x56\xe1\xca\x22\xdf\x08\x20\x07\x4d\x84\x3c\xc7\xac\x68\xcd\xcb\x4c\xe2\xf9\x0a\x77\x8e\xba\x74\x0f\x50\xf0\x52\x50\xb8\xbd\xc2\x36\x6d\x7b\xdc\x65\x78\x6c\x0d\xc1\x24\xcf\x21\x4b\x82\xe1\xfb\x10\x0f\x1c\x46\x51\x34\x43\xda\xdc\xca\x53\x9b\x37\x54\x8e\x4d\x82\x55\x35\xe6\x81\xa2\x0d\x55\x9a\x37\x59\xde\x6d\x89\xc3\x17\x06\x93\x16\xd7\x3d\x90\x04\x84\x86\x54\x4d\xb6\xe8\x47\xfd\x7e\x3a\x23\x05\x80\xa3\xf3\xd3\x0a\xfb\xbe\x86\xa0\xe2\x89\xbe\x80\x0c\xfc\xdf\x16\x77\xc5\xe2\x8a\xd3\xf2\x08\x99\x4c\x36\x8b\xcf\x16\x8f\x1f\xe2\xa7\xdf\xfd\xa7\xb9\x2d\x67\xfe\xe4\xff\xf1\xa4\xe3\xe1\x29\x4b\x56\x98\x46\x09\x0f\xb4\xee\x3b\x90\x7b\x9e\x7c\xe6\x85\xaa\x3e\x9d\xde\x73\x3f\x4e\x38\x2d\xfc\x10\xfb\xd5\xd3\x9c\x87\xc6\xd8\x2f\x92\xc8\xb2\x30\x01\x4a\x42\x7c\x61\x85\x23\x6c\xeb\x32\xc4\x95\xc1\xb3\x1a\x81\xc9\x78\x8f\x8c\xb4\xa1\x8a\xb7\x1b\xfa\xc9\x95\xb3\xc4\x74\xba\x06\x5f\x1c\x90\x96\x42\x1a\xa9\x8a\xa5\x91\x31\x6d\x36\x29\x57\x4d\x55\x3b\xde\x6d\x4a\x5a\xcd\xec\x7f\x9b\x1c\x37\xd0\xcd\x8d\xd8\x9f\xb7\x0e\x76\xf3\x64\x98\xc6\xca\x9d\x5f\x7e\x0a\x4d\xe3\xf9\xbb\xd9\x52\x2f\x26\xff\x50\x79\xb2\xad\x2d\x24\xba\x7e\x24\x3b\xd5\x0c\xab\x58\xdb\x9d\xf3\x5a\x08\x75\xfb\x00\xdf\x4a\x26\x20\xd1\x31\x1c\xc7\xcf\x41\xe9\x76\xbb\xb8\x23\x92\xee\xfd\xb0\xa7\x6c\xf6\x2f\xf6\x81\x72\x00\x6a\xdb\xbe\xdb\xab\x71\xb8\xbb\x74\x1d\xc8\xdf\xa7\x53\x54\xdb\xe5\x20\xf4\x63\xb8\x6b\x68\x2b\x64\x5f\x80\xf5\x7d\x10\xfd\x78\x38\x9b\xa6\xe5\xe2\x69\xa2\x35\x5e\x99\x93\x78\x8e\x81\x79\xcf\xd7\x3c\x93\x7a\x5f\xd3\x55\xa0\xff\x34\x20\x4e\x81\x67\xb6\xfe\x41\xec\xcc\x2a\x3c\xf6\x21\x30\x34\x08\x9c\xd5\xbb\x3f\x09\x42\x27\xe4\x4b\xec\xf6\x67\xad\x6e\x69\x1e\xd3\x33\xa3\xf9\x02\x79\x85\xd9\xaf\x9f\x2a\xfd\x48\x54\xb2\xec\x49\xc3\x5a\x6b\xec\xa9\x37\xc3\xf5\xaa\xde\x2e\x6b\x03\x8a\xdd\x8e\xc6\x22\xfd\x09\xd1\xa6\xba\xd9\x8e\x26\xa7\x0d\x8b\x9c\x6f\x91\xa9\xac\x5d\x5b\x84\x7e\xd8\xe1\xac\x36\x31\xd3\x72\xd8\xd0\x74\x61\x23\x9b\x07\xeb\x6f\x1e\xfd\xe9\x51\xd8\x2d\xa4\xbf\xeb\xda\xee\x76\x3c\x5a\xd5\x8c\xd1\xcd\xb0\x27\x4a\xa0\x09\xdc\xaa\x1b\x0d\xdd\x3a\x50\xbf\x67\x59\xa0\x39\x19\xfa\x31\xa8\x4d\xf6\x93\xdb\x42\xa7\x98\x8a\xef\xa3\x28\x41\x79\xe7\xd0\x57\x65\xbd\x78\x7c\xce\x41\x11\x91\x3c\x4f\x19\xd5\xee\xc6\x9a\x73\xfe\xb5\xa8\x82\xb2\xd5\xd0\xbc\x4f\x4b\x73\x25\x35\x4d\x0d\x43\x5a\xc0\xcb\x12\x75\x90\x17\xd7\x19\xd9\xa4\x40\x05\x2f\x8a\x46\x98\xed\xe7\x07\x26\xb4\xd9\xa5\xdf\x98\xd1\x17\x2c\xd2\x11\xfe\xef\xa5\xf5\x5f\x01\x00\x00\xff\xff\x87\x26\x92\xd3\x9d\x13\x00\x00")

func templateGoBytes() ([]byte, error) {
	return bindataRead(
		_templateGo,
		"template.go",
	)
}

func templateGo() (*asset, error) {
	bytes, err := templateGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"collection.go": collectionGo,
	"cosmosdb.go":   cosmosdbGo,
	"database.go":   databaseGo,
	"document.go":   documentGo,
	"template.go":   templateGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"collection.go": &bintree{collectionGo, map[string]*bintree{}},
	"cosmosdb.go":   &bintree{cosmosdbGo, map[string]*bintree{}},
	"database.go":   &bintree{databaseGo, map[string]*bintree{}},
	"document.go":   &bintree{documentGo, map[string]*bintree{}},
	"template.go":   &bintree{templateGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
