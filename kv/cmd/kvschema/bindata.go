// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// templates/kvschema.gotmpl
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

var _templatesKvschemaGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5f\x6f\xdb\xc8\x11\x7f\xd6\x7e\x8a\x39\x3d\x1c\x48\x47\xa6\x1c\xa3\x40\x53\xb9\x2e\xe0\xb3\xdd\x54\x48\xce\x0e\x2c\x27\xc1\xc1\x30\x8a\x15\x39\x92\x16\xa2\x76\x99\xdd\xa5\x64\x82\xe0\x77\x2f\x66\xb9\x24\x25\x59\xb1\x1b\x14\x01\x7a\x2f\x89\x49\xce\xce\xfc\x66\xe6\x37\x7f\x56\x65\x39\x3c\x02\x76\xa9\xb2\x42\x8b\xf9\xc2\xc2\xe9\xc9\xdb\xbf\xc1\x7b\xa5\xe6\x29\xc2\xc7\x8f\x97\x8c\x7d\x14\x31\x4a\x83\x09\xe4\x32\x41\x0d\x76\x81\x70\x91\xf1\x78\x81\xe0\xbf\x0c\xe0\x0b\x6a\x23\x94\x84\xd3\xe8\x04\x02\x12\xe8\xfb\x4f\xfd\xf0\x8c\x15\x2a\x87\x15\x2f\x40\x2a\x0b\xb9\x41\xb0\x0b\x61\x60\x26\x52\x04\x7c\x8a\x31\xb3\x20\x24\xc4\x6a\x95\xa5\x82\xcb\x18\x61\x23\xec\xc2\x19\xf1\x2a\x22\xf6\x87\x57\xa0\xa6\x96\x0b\x09\x1c\x62\x95\x15\xa0\x66\xdb\x52\xc0\x2d\x63\x00\x00\x0b\x6b\x33\x33\x1a\x0e\x37\x9b\x4d\xc4\x1d\xcc\x48\xe9\xf9\x30\xad\xc5\xcc\xf0\xe3\xf8\xf2\xfa\x66\x72\x7d\x7c\x1a\x9d\x30\xf6\x59\xa6\x68\x0c\x68\xfc\x96\x0b\x8d\x09\x4c\x0b\xe0\x59\x96\x8a\x98\x4f\x53\x84\x94\x6f\x40\x69\xe0\x73\x8d\x98\x80\x55\x04\x74\xa3\x85\x15\x72\x3e\x00\xa3\x66\x76\xc3\x35\xb2\x44\x18\xab\xc5\x34\xb7\x3b\x11\x6a\x60\x09\x03\xdb\x02\x4a\x02\x97\xd0\xbf\x98\xc0\x78\xd2\x87\xdf\x2e\x26\xe3\xc9\x80\x7d\x1d\xdf\xff\xeb\xf6\xf3\x3d\x7c\xbd\xb8\xbb\xbb\xb8\xb9\x1f\x5f\x4f\xe0\xf6\x0e\x2e\x6f\x6f\xae\xc6\xf7\xe3\xdb\x9b\x09\xdc\xfe\x13\x2e\x6e\xfe\x80\x0f\xe3\x9b\xab\x01\xa0\xb0\x0b\xd4\x80\x4f\x99\x26\xec\x4a\x83\xa0\xd8\x61\x12\xb1\x09\xe2\x8e\xf1\x99\xaa\xc1\x98\x0c\x63\x31\x13\x31\xa4\x5c\xce\x73\x3e\x47\x98\xab\x35\x6a\x29\xe4\x1c\x32\xd4\x2b\x61\x28\x7b\x06\xb8\x4c\x58\x2a\x56\xc2\x72\xeb\x9e\x9f\xb9\x13\xb1\xa3\x61\x55\x31\x56\x96\x09\xce\x84\x44\xe8\x2f\xd7\x26\x5e\xe0\x8a\x47\x73\xd5\xaf\xaa\xe1\x10\x2e\x55\x82\x30\x47\x89\x9a\xdb\x3a\xa2\xad\x4c\xff\x0c\xae\x6e\xe1\xe6\xf6\x1e\xae\xaf\xc6\xf7\x11\x63\x19\x8f\x97\x84\xa6\x2c\xa3\x4f\xf5\x9f\xd1\x0d\x5f\x21\x59\x10\xab\x4c\x69\x0b\x01\xeb\xf5\xe7\xc2\x2e\xf2\x69\x14\xab\xd5\x70\xee\x68\x39\x94\xca\xe2\xf1\x8a\x67\x66\xb8\x5c\xf7\x59\xc8\xd8\x70\x08\x13\xab\x34\x42\xa6\xd5\x5a\x24\x68\x00\xa5\x15\x56\xa0\x19\x38\x6a\x29\x89\xd2\x9a\x01\x39\x08\x42\x26\xf8\x84\x06\xa6\x3c\x5e\xfa\x94\xc3\x12\x8b\xe3\x35\x4f\x73\x24\x4d\x86\x34\x45\x6c\x38\xa4\x87\xcf\x86\xcf\x71\xc4\x86\xc3\xb2\x74\xc4\x74\xa7\x21\xba\x6c\x94\xde\x17\x19\x1a\x38\xa9\x2a\x12\x06\xf2\x64\xf2\x85\xeb\xaa\x1a\x00\x6a\x0d\xa3\xf3\x1a\x57\xe9\xfe\x1d\xd5\xaa\xab\xa8\x2c\xbd\x9f\xad\x9a\xe0\x24\x8c\x26\x31\x97\xc1\xc3\xe3\x72\x1d\x5d\x13\xf8\xa2\xfc\xeb\x00\xfe\x72\x5a\x85\xce\x38\xca\xa4\xaa\x98\x2d\x32\xf4\x9e\x1a\xab\xf3\xd8\x42\xc9\x7a\xcb\x75\xe4\x5e\xb1\x5e\xc6\x35\x79\xad\x24\xb4\x4a\x58\xe5\xa2\xf3\xa9\xfd\xa2\xd1\xe6\x5a\x1a\x97\xd5\x5a\x04\xec\x82\x5b\x22\x6a\x4e\x45\xce\x0d\x70\xe8\x14\x11\x85\x78\x9a\x82\xca\x28\xa1\xc4\x8a\x88\xcd\x72\x19\x43\x60\x6a\x20\x61\xa7\x3b\x08\x3b\xbb\x04\xac\x36\x05\x26\x6a\xd5\x79\x34\x5f\x85\x5d\x3c\x47\xc4\x41\xe2\xc6\x7b\xe7\x62\x8d\x04\x86\x70\xfe\x10\x9c\x1d\xe5\x01\x76\x90\x42\x38\xaa\x95\x97\xac\xb7\x85\x09\xce\x01\x5b\xac\xbf\x1a\x56\xb1\xb2\xd4\x5c\xce\x71\x3f\xcb\x55\x55\x96\x16\x57\x59\xca\x2d\x42\xbf\xa5\x55\x1f\x22\xfa\xe2\x12\xd4\xe4\x69\xab\x40\x3a\xb9\x9a\x23\x13\xb4\x6d\xfa\xc1\xa0\xad\x3d\xec\x5e\x71\x63\x54\x2c\x5c\xed\xf8\x20\x58\x05\xeb\x86\x8e\x97\x4a\x6b\x34\x99\x92\x09\xd5\x6e\x43\x65\xae\x11\xf2\x2c\xa1\x43\x5d\x38\x8e\x7c\x3c\xb6\x0d\x6e\x87\x63\x00\x6b\x28\x4b\x31\x83\xe8\x4a\x68\x8c\xed\xb5\x8c\x55\x82\xda\x39\x91\x1a\xac\xaa\xa3\xd6\x29\x7f\x3a\x24\x4e\x2b\xed\x38\x87\x05\x91\x7b\xc5\x97\x18\x2c\xd7\xd1\x27\x8d\x33\xf1\x34\x80\x77\x6f\x4e\xdf\xbc\x0b\x77\xc2\x1b\xd5\x7a\x2f\x6c\xb0\xc4\x22\x64\x3d\xaa\x75\x27\x5d\xeb\xdc\xf9\xfc\xf0\x6e\xf4\x18\xb2\x1e\xee\xbe\x7c\x7b\xe2\xde\x96\x25\x10\xd8\xb1\xf7\xb9\xaa\xd6\x5c\x83\x4a\x93\x2e\x76\xac\x27\x66\x4d\xd9\x99\xe8\x3d\xba\xe3\x03\x92\x89\xae\x90\x14\x86\x67\xee\xf3\x2f\xe7\x20\x45\x4a\x6e\x34\x69\x47\xad\x59\x6f\xef\xfc\xa4\x39\xbf\xf6\x70\x82\xf0\xd5\xf3\x29\x2e\xe9\x70\x8a\xd2\x7b\xdb\x46\x9b\x0a\xfc\x90\x57\x14\xc8\x73\x1a\x3a\x28\x93\xda\xdc\x72\x1d\xed\xb4\x85\xc6\x78\x14\x45\x21\xeb\x91\xd3\x01\xeb\xf5\x52\xb1\x84\x6d\x43\x3d\x34\xd0\xe5\x76\x42\xe3\x8e\xf5\xc2\xb2\x04\x4f\xe5\x2e\x6c\x8c\xf5\xa8\xaf\x39\xba\x6c\xf1\xce\x71\xa9\xc1\x43\x00\x47\x29\x2e\x1f\xa3\x0b\x87\xac\x03\xb4\x97\xbe\x90\xf5\xa8\x24\xff\x3d\x00\xb1\x26\xcf\x6b\x6b\x14\xf1\xb2\x8c\x7e\x47\xbb\x50\x89\x67\x5e\xe8\xe2\xb5\xef\xee\xc3\x28\x15\xcb\x47\x3a\xbd\xe7\xe7\xe1\x54\xa2\xf9\x6e\x26\x77\x52\x41\xb9\x70\x1a\x4c\x74\x87\x2b\xb5\xc6\x00\x6b\xfb\x87\x33\x8c\xe6\x85\x14\xef\x69\x76\xaa\x2b\x97\xed\x03\x9e\xaf\xff\x6f\xfc\x1e\x4b\x83\xda\xfe\x14\xbf\xfd\x7b\x29\xd2\xb2\x04\x94\x09\x50\x8f\x00\x6a\x1a\x50\x55\x6d\xd3\x3f\x54\x3f\xad\xbc\x1f\x04\xef\xb7\x7a\x93\xe3\xec\xd6\x30\xe8\xa8\x49\x71\x46\x1e\x2f\xea\xa1\x5e\xd0\x1a\x86\xa6\xe9\x88\xe3\x7a\x0d\x74\xab\x4a\x5a\x50\x57\xa4\x21\x4b\x5b\x45\xab\x4a\x02\xae\x32\x5b\x80\x9b\xf0\x75\x57\x95\xca\xf7\x32\xd2\xbd\xc4\xc2\xb8\x21\x48\xfa\x12\xe5\x16\x55\x7c\x12\xc6\xd6\x0b\x43\x0b\xc4\xe7\x20\x08\x21\xe6\x12\x12\xf7\xd0\x69\x9f\x16\x16\xc1\x90\x13\x03\x42\x24\x81\x93\xba\x6f\x39\xea\xa2\xf6\xc0\x6f\x24\xf5\xb8\xa5\x96\x4d\x76\xf6\xdb\xfd\xb6\xdf\x66\xa1\xf2\x34\x81\x26\xdc\x8a\xf4\x39\xd4\xe6\x79\x97\x7f\x16\xc9\x00\x0d\x6c\x2d\x13\x21\x04\x0f\x8f\xad\xc4\xa0\xd6\x13\xd6\x53\xda\xe4\xa9\x6d\xbb\xf9\x8e\x14\xf5\x17\x34\x61\xf8\x93\x1b\x3e\x85\x47\x0c\x00\xbb\x3a\x42\xe3\x38\x78\x78\x12\xf4\x9e\x17\x48\xf0\x6b\xed\xc6\x83\x78\x0c\x9b\x52\xe9\x8a\xe9\x40\xb9\x48\x91\x0e\xba\x9a\xe9\x48\x5d\xab\x19\xd0\x77\x56\xb5\x8b\x80\xef\x9e\x55\xcd\xda\x6b\x9f\xc9\xdf\xb9\x8d\x17\x42\xce\xcb\xb2\xeb\xd8\xb5\x93\x5d\x0e\x1b\x12\xb6\xd9\x77\x59\x7e\x7e\xa2\x26\xa7\x27\x87\xc7\xc2\x61\xe5\x2d\xd0\x01\xda\x41\xae\x9f\x32\x4d\x05\xa1\xd5\x8a\x38\x26\x34\xec\xb5\x1b\x58\xb9\x87\xa6\x38\xee\x17\x4d\x19\x60\x02\x5b\xa3\x81\xb6\x3d\x9e\x6a\xe4\x49\x01\x46\xe9\x83\x7b\xc3\x0f\x78\x19\xac\x77\x01\x86\x10\xec\x8e\xa2\x6d\xbe\xbd\xc4\xa4\x37\xa7\xaf\x72\xa9\xc5\xf0\x1a\xa9\x5e\x1d\xbd\x2f\x12\xf3\xed\xbb\xef\x8c\xe7\xfd\x86\x4d\x03\x19\xcd\xb3\xd1\xdb\xf4\x4e\x33\x38\xdc\xc7\xd9\x2e\x95\x7e\x2b\x7e\x98\x44\x74\xfc\xfb\x3c\x52\x3a\x41\x7f\xa3\xf5\xfb\xe5\x16\x7f\xbc\x4c\x47\x23\xaf\xeb\x05\x26\xdd\x21\x77\x2b\xe7\x14\xe7\x82\x7a\xaa\x85\x38\xd7\x46\xe9\xba\x45\xa2\x4c\x0c\x6c\xa8\xe9\x91\xb1\x14\xe5\xdc\x2e\x9a\x0b\xfa\x1e\xff\xdc\x0d\xab\xe1\xa0\xbb\x7d\xdb\x05\x97\x20\x23\xf8\x4a\xe7\xb5\xb7\x23\x8c\x6b\x8f\xee\x77\x01\xb4\x38\xf0\xe6\xdc\x35\xa5\xde\x74\xc1\xe4\xb1\x0b\x82\x2b\x99\xdc\xb8\x53\xee\xc7\x04\x0e\x26\x9f\xe2\xb7\x1c\xa5\x85\x98\x6e\x0b\x56\x81\x0b\xb0\xf7\x6c\xb3\xd3\x56\xf1\xc9\x82\x74\xad\xd5\x47\xf7\xfb\xa5\xf0\x62\x96\x02\x8f\xf0\x68\xb9\xae\xfb\xc5\xa5\x0f\x90\x04\x21\x6d\x08\x7b\x0d\xb9\xbe\x20\xfe\x49\xcb\x42\x58\xf4\x1d\xb8\x16\x1d\x5b\xba\x8e\x29\xed\xf1\x24\x38\x43\x0d\x24\x14\x5d\x09\x13\x73\x9d\x04\xfe\x50\x34\x41\x5c\xfa\x48\x45\x1f\x9c\xb0\x98\xc1\x2f\xee\xd3\x17\x9e\x8a\xc4\xef\x4b\x75\x72\x5c\x63\xa6\x02\x9b\xe6\xb3\x67\x15\xe6\x7b\xfb\x39\x34\x87\x73\x0c\xa6\xf9\xec\x95\x55\xbf\x59\xf3\x3d\x84\xdb\xd9\xcc\xa0\x85\xbf\xbb\x49\x37\xcd\x67\xb5\x75\x34\x5d\xd1\x53\x05\x4f\xf3\xd9\xc3\xce\x81\xd1\x63\xb7\xad\xf9\x19\x09\xff\x38\x07\x59\x8f\x98\x5d\xdd\x6f\xce\x41\xfa\x0d\xac\x15\x6d\x24\x6b\x4b\x68\x1e\x46\xf2\xb1\xd9\xb1\x5a\xa0\x5b\x5b\xa6\x73\xf1\x06\x9f\x6c\x10\x9e\xc1\x76\xb0\xce\xb6\x3f\x39\x9d\x3f\x1a\x96\x1d\x73\x87\x3c\xff\x2f\x3c\xfd\xb0\xdd\x25\xbb\x77\x0f\x27\xa3\x13\xda\x71\x09\xc8\x07\x2c\xda\x15\x77\x2f\x40\xe7\x5d\xf0\x8f\x21\x68\x8c\x1c\x83\x0c\xff\x97\xb0\xf9\xc7\xf6\x4a\xde\xfc\xff\x9f\x00\x00\x00\xff\xff\x97\x20\x9b\x77\xef\x14\x00\x00")

func templatesKvschemaGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesKvschemaGotmpl,
		"templates/kvschema.gotmpl",
	)
}

func templatesKvschemaGotmpl() (*asset, error) {
	bytes, err := templatesKvschemaGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/kvschema.gotmpl", size: 5359, mode: os.FileMode(420), modTime: time.Unix(1564182148, 0)}
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
	"templates/kvschema.gotmpl": templatesKvschemaGotmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"kvschema.gotmpl": &bintree{templatesKvschemaGotmpl, map[string]*bintree{}},
	}},
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
