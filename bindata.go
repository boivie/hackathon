// Code generated by go-bindata.
// sources:
// answers.hashed
// DO NOT EDIT!

package main

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _answersHashed = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x14\x92\x4d\xea\x65\xb7\x0e\xc4\xff\x97\xf7\x20\x10\x7a\x9e\x2d\x78\x28\x59\xb6\x2c\x2d\x22\x8b\xd0\x67\xc8\x24\xf4\x16\x92\x85\xa7\xa3\x86\x7b\xe1\x70\x74\xac\x72\xfd\xaa\x7e\xfb\xf1\xf7\xff\x3f\x9f\x5f\x7e\xb7\xef\xdf\xff\xfc\xeb\x8f\xcf\x8f\x7f\xbe\x3e\xdf\x3e\xdf\xbe\xbe\xfe\xfd\x9f\xcc\xf3\xaf\xeb\x88\x34\x10\xde\x3e\x2c\x21\x74\x99\xc5\x83\x4e\x61\x68\x66\x88\xea\x46\x31\xf6\xb3\x2a\x77\x2a\x77\xcc\xff\x31\x64\xaa\xb3\x40\x50\xde\x0e\x78\x6f\x37\x6c\xb5\xcb\x50\x0b\xfc\xa1\x26\x90\x75\x5c\x41\x48\x98\x09\x95\xc0\x66\x40\xad\xc6\x12\xf7\x06\xb3\x85\x77\xb3\x9c\xd7\x00\x89\x7c\x98\xe0\x7a\x5f\x7d\x3d\xf2\xa4\x8f\x19\x95\x55\x8b\x7a\x25\xa6\x04\x6f\x3b\x73\xde\x0e\x51\xc1\x2d\x55\x74\x71\xac\xe7\x25\xc7\xf1\x94\xd4\x59\xe7\xe8\x86\x1d\xfd\x4c\x62\x34\x00\x84\x7b\x3b\xee\xad\xd8\xd5\x63\x35\x20\xfc\x16\x2d\x68\xa5\xf9\xa8\xa0\xee\xde\x37\xdf\x23\x7f\x19\xd1\x04\xed\x20\x1d\x51\x2c\xdb\xc7\xb5\x6d\x6e\x84\x52\x01\xc4\xc3\x0a\x01\x67\xbf\x3c\x86\xf8\xf0\x98\x31\x28\x14\x22\xdd\x95\xcd\x5d\xb0\x67\x3c\x10\x5e\xa2\x78\xf5\xab\x8e\xec\x4d\x52\x75\x33\x3c\x05\x8b\x57\x3f\xe0\xf9\xf1\xa5\xed\xa9\xf7\x6a\xc5\x11\x4c\x7e\xdb\x2b\x0f\xe0\xb3\x56\x9d\xab\x2e\x35\x10\x23\x25\x21\x0a\xf7\x7b\xd1\x28\x5f\x42\x3d\x26\x3e\xe2\xe3\x17\x90\x9c\x78\x0d\x15\xba\x32\xe4\xee\x3b\xde\x13\x24\xb1\xdb\x09\x99\xf8\x00\x74\x74\x78\xcb\x20\xac\x25\xfb\xd2\x18\x9b\xf5\x1d\xe7\xe7\xfd\xc2\x11\x67\x04\x2c\xe6\x59\x74\xc4\xcc\xe3\xee\x55\xd0\x16\xb9\xe5\x0d\xf8\xbd\x2b\x7c\xc7\xed\x1e\xd9\x28\x99\x88\x59\xa6\x2b\xa6\xa5\xab\xee\xf1\x3d\xdb\xd5\xfc\x60\xeb\xb0\x9c\xfc\x2e\xe3\x4e\xe9\x33\x6f\x9b\x39\x55\x6a\x2f\x0d\x3e\x15\x88\x81\x10\xba\xaf\xbc\xd1\x1a\x77\x70\x07\xb8\xf4\xf5\xdb\x13\x9b\xa8\xad\xe0\xe7\x2d\xa7\xd2\xf6\x05\x39\x3f\x17\x22\x95\x51\xd8\xe0\x70\x00\xbb\x06\x96\xc9\x8b\xb0\x11\x0d\x7a\xac\xcc\x11\x9c\x2a\x55\x84\xc5\x00\xe9\x11\x9e\x2e\x4f\x23\x06\xff\x5d\xc5\x94\x09\x97\x4f\xe6\xd0\x89\xd6\x36\x07\xdc\x42\xf6\xfc\x4c\x8f\x2c\xed\x02\x29\x2e\x1e\xbb\x36\xd5\xe2\x6c\x3c\x4d\x5a\x77\xda\x7f\x89\xb7\x8a\x76\x85\x21\x31\x79\xed\xfd\x5f\x00\x00\x00\xff\xff\x2e\x52\x3a\xa5\x53\x03\x00\x00")

func answersHashedBytes() ([]byte, error) {
	return bindataRead(
		_answersHashed,
		"answers.hashed",
	)
}

func answersHashed() (*asset, error) {
	bytes, err := answersHashedBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "answers.hashed", size: 851, mode: os.FileMode(420), modTime: time.Unix(1463056774, 0)}
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
	"answers.hashed": answersHashed,
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
	"answers.hashed": &bintree{answersHashed, map[string]*bintree{}},
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

