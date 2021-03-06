// Package bindata Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// templates/kvschema.gotmpl
package bindata

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

var _templatesKvschemaGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6d\x6f\xdb\xba\x15\xfe\x6c\xfe\x8a\x33\x7f\x18\xe4\x56\x95\xd3\x7e\xea\x5a\x64\x80\x9b\x64\x9d\xb1\xde\xe4\x22\x76\x57\x5c\x04\xc1\x40\x4b\xc7\x36\x61\x99\x54\x49\x4a\x8e\x27\xe8\xbf\x0f\x87\xa2\x5e\x6c\xe7\xe5\xf6\x02\x17\x18\xb6\xfb\x29\xb1\xc8\xf3\xfa\x3c\x7c\x74\xa8\xb2\x1c\xbf\x02\x76\xa1\xb2\xbd\x16\xab\xb5\x85\x77\x67\x6f\xff\x02\x9f\x95\x5a\xa5\x08\x5f\xbe\x5c\x30\xf6\x45\xc4\x28\x0d\x26\x90\xcb\x04\x35\xd8\x35\xc2\x24\xe3\xf1\x1a\xc1\xaf\x84\xf0\x4f\xd4\x46\x28\x09\xef\xa2\x33\x08\x68\xc3\xd0\x2f\x0d\x47\x1f\xd9\x5e\xe5\xb0\xe5\x7b\x90\xca\x42\x6e\x10\xec\x5a\x18\x58\x8a\x14\x01\x1f\x62\xcc\x2c\x08\x09\xb1\xda\x66\xa9\xe0\x32\x46\xd8\x09\xbb\x76\x41\xbc\x8b\x88\xfd\xe2\x1d\xa8\x85\xe5\x42\x02\x87\x58\x65\x7b\x50\xcb\xfe\x2e\xe0\x96\x31\x00\x80\xb5\xb5\x99\xf9\x30\x1e\xef\x76\xbb\x88\xbb\x34\x23\xa5\x57\xe3\xb4\xde\x66\xc6\x5f\xa6\x17\x57\xd7\xb3\xab\x37\xef\xa2\x33\xc6\xbe\xca\x14\x8d\x01\x8d\xdf\x73\xa1\x31\x81\xc5\x1e\x78\x96\xa5\x22\xe6\x8b\x14\x21\xe5\x3b\x50\x1a\xf8\x4a\x23\x26\x60\x15\x25\xba\xd3\xc2\x0a\xb9\x0a\xc1\xa8\xa5\xdd\x71\x8d\x2c\x11\xc6\x6a\xb1\xc8\xed\x41\x87\x9a\xb4\x84\x81\xfe\x06\x25\x81\x4b\x18\x4e\x66\x30\x9d\x0d\xe1\xd3\x64\x36\x9d\x85\xec\xdb\x74\xfe\xf7\x9b\xaf\x73\xf8\x36\xb9\xbd\x9d\x5c\xcf\xa7\x57\x33\xb8\xb9\x85\x8b\x9b\xeb\xcb\xe9\x7c\x7a\x73\x3d\x83\x9b\xbf\xc1\xe4\xfa\x17\xf8\xc7\xf4\xfa\x32\x04\x14\x76\x8d\x1a\xf0\x21\xd3\x94\xbb\xd2\x20\xa8\x77\x98\x44\x6c\x86\x78\x10\x7c\xa9\xea\x64\x4c\x86\xb1\x58\x8a\x18\x52\x2e\x57\x39\x5f\x21\xac\x54\x81\x5a\x0a\xb9\x82\x0c\xf5\x56\x18\x42\xcf\x00\x97\x09\x4b\xc5\x56\x58\x6e\xdd\xef\x93\x72\x22\xf6\x6a\x5c\x55\x8c\x95\x65\x82\x4b\x21\x11\x86\x9b\xc2\xc4\x6b\xdc\xf2\x68\xa5\x86\x55\x35\x1e\xc3\x85\x4a\x10\x56\x28\x51\x73\x5b\x77\xb4\xdd\x33\xfc\x08\x97\x37\x70\x7d\x33\x87\xab\xcb\xe9\x3c\x62\x2c\xe3\xf1\x86\xb2\x29\xcb\xe8\xe7\xfa\xdf\xe8\x9a\x6f\x91\x22\x88\x6d\xa6\xb4\x85\x80\x0d\x86\x2b\x61\xd7\xf9\x22\x8a\xd5\x76\xbc\x72\xb4\x1c\x4b\x65\xf1\xcd\x96\x67\x66\xbc\x29\x86\x6c\xc4\xd8\x78\x0c\xf3\x07\x09\x99\x56\x85\x48\xd0\x00\x4a\x2b\xac\x40\x13\x3a\x62\x29\x89\xd2\x9a\x90\xca\x03\x21\x13\x7c\x40\x03\x0b\x1e\x6f\x3c\xe0\xb0\xc1\xfd\x9b\x82\xa7\x39\x82\xb1\x4a\x63\xc4\xec\x3e\x43\xe7\xd0\x58\x9d\xc7\xb6\x84\x4d\x11\xfd\xcc\x35\xf9\x54\x12\x13\xa8\x18\x5b\xe6\x32\x86\x6b\xdc\x05\x96\x16\xe7\x0f\x72\xe4\x0c\x4a\xd0\x68\x73\x2d\xe9\x47\x79\x68\x55\xda\x10\xce\xaa\x0a\x2a\x56\x96\x9a\xcb\x15\x42\x74\xd1\x24\x37\xdf\x67\x68\xaa\xaa\x2c\x2d\x6e\xb3\x94\x5b\x84\x61\x9b\xf8\x10\x22\x5a\x41\x99\xb4\x7f\xfa\x00\x74\xfb\xaa\x8a\xfa\x30\x43\x5b\x96\xbe\x8d\x60\xd0\x1a\x87\x5f\xf7\x88\x1b\xa3\x62\xe1\xb0\x71\x27\x0d\x89\xd8\x45\xc4\xc6\x63\xe6\xd0\xd3\x1a\x4d\xa6\x64\x42\xdc\x68\x9a\xc5\x35\x42\x9e\x25\x64\x14\xd5\x95\x07\x06\x5c\xcd\xfd\x68\x01\x52\x2b\xae\xa8\xf5\xfb\x10\x0a\x28\x4b\xb1\x84\xe8\x52\x68\x8c\xed\x95\x8c\x55\x82\xda\x55\x90\x1a\xac\xaa\x57\x6d\x45\xde\x7a\x04\xa8\xb5\xd2\x50\xb2\xc1\x06\xf7\xf0\xe1\x1c\xb6\x7c\x83\x01\xf5\x50\xe3\x52\x3c\x84\xf0\xfe\xf5\xbb\xd7\xef\x47\x6c\x60\xba\xae\x46\xb5\xdf\x89\x0d\x36\xb8\x1f\xb1\x01\x11\xc9\xed\xae\x7d\x1e\x2c\xdf\xbd\xff\x70\x3f\x62\x03\x3c\x7c\xf8\xf6\xcc\x3d\x2d\x4b\xa0\x64\xa7\xbe\xe0\xaa\x2a\xb8\x06\x95\x26\x5d\xe3\xd8\x40\x2c\x29\x45\xca\xcc\x44\x9f\xd1\x99\x87\xb4\x27\xba\x44\x72\x38\xfa\xe8\x96\xff\x74\x0e\x52\xa4\x54\xc6\xc0\x53\x01\xb5\x66\x83\x23\xfb\x59\x63\x5f\xf8\x74\x82\xd1\x8b\xf6\xe3\x31\x4c\x20\x73\xe5\xc1\x22\x5f\x2e\x51\xbb\x03\xce\xd3\xb4\x06\x8a\x78\x6c\x22\x36\xf0\x5b\x3e\x9c\x13\x1c\x17\x4a\xc6\xdc\x7e\xda\x5b\x9c\x91\x04\x9a\x3a\xaa\x5b\xf0\xbc\x09\xce\x46\x5d\x0e\x6c\xd0\x42\xd8\x3d\x9f\xd8\xa0\xf6\xd9\x74\x8b\x9a\x83\xa6\x43\xdb\xb9\x2e\x4b\xf0\xb4\xee\xba\xc8\x5c\xd6\x5f\x1d\x75\x7a\x1c\x74\xe9\x3e\x83\x56\x13\xad\x46\x8c\x8a\xfc\x57\x08\xa2\xa0\x92\xea\x10\xd4\xf5\xb2\x8c\x7e\x42\xbb\x56\x89\x67\xdf\xc8\xf5\x6c\xf3\x54\xdd\x99\x67\x91\x28\xfa\xd5\x9e\x82\x1a\x02\x9a\x27\x11\x3d\x80\x84\x30\x71\xf6\x26\xba\xc5\xad\x2a\x30\xc0\x3a\x87\x53\xa4\x9d\xd3\xa7\x81\x3e\xf2\xeb\x1c\x57\x0e\xf3\x47\x6a\x2f\xfe\xab\x2a\x9f\x4a\x83\xda\xfe\x0e\x95\xfb\xe7\x52\xa4\x65\x09\x28\x13\x20\xad\x00\x12\x0f\xa8\x2a\xbf\xf8\xf8\x39\x6a\xf7\xb3\xca\xbd\x19\x2e\x31\x45\x8b\x1d\xfb\x12\xf7\xfb\x45\x5d\xfc\xad\x92\x78\x14\xae\xaf\x8a\xff\x5f\x1a\x57\x37\xc2\x65\xfd\x92\xd9\x1f\x92\xf5\xbf\x23\x59\xbf\xee\xe0\xf6\xc8\x71\x7c\x5e\x3f\xf7\x27\x98\xda\xe0\x57\x1f\xd6\xe9\x12\xa4\xea\x6d\x5c\x73\x03\x0b\x44\x49\xe3\x72\x2a\x62\x61\xd3\x3d\x0d\x45\xee\xc5\x89\xf5\x44\x78\x10\x6e\x27\xd2\xd4\xc7\x24\x77\x14\x55\xa3\xc9\x53\x4b\xd7\x8d\x84\x5a\x4c\x22\xc0\x7b\x11\x96\x5a\x6d\x69\xa6\xc7\x6d\x66\xf7\x60\x08\x39\xda\xbb\xd8\x5b\x34\x47\xca\xf0\xf9\x89\x61\x69\x04\x41\xfb\x3c\xac\x25\xc2\xa1\x42\x9c\x2d\xfa\xa7\xb3\x30\xe1\x01\xf4\xed\x92\x23\x4c\x70\x77\xdf\xba\x2c\xb1\x1a\xb9\xd3\x98\xa2\x0c\x0a\x33\x82\xbf\x9e\xc3\x5b\x07\x5c\x01\xe7\x50\x98\xbb\xb3\xfb\x3e\x58\x85\xf3\xfb\x48\xff\x9d\xe3\x16\x84\x83\xba\xa9\x83\x3c\x5e\xd7\xb3\xf6\x9e\xee\x46\x54\xf0\x8f\xc3\x40\xbd\xf3\x33\x23\xc1\xd1\x6b\x39\x81\x41\xde\x16\x78\x10\xd9\xae\xb9\xed\x3c\x3a\x50\x30\xf9\xad\x38\xd4\x9d\x43\x03\xbd\xe6\x8d\x20\xb8\xbb\x7f\x14\x11\x9f\x58\x23\xdc\x07\xbb\xa8\xd3\x68\x9c\x1e\xfd\x9e\xda\x4e\x2d\x13\x21\x60\xa7\x2c\x68\x1c\xb0\x8f\x8b\xfe\xe0\x54\xce\x83\x3f\xd7\x65\xdc\x89\xfb\x51\x23\x1b\x9d\xac\x3c\x22\x1d\x52\xa4\x61\xa7\x1f\x1d\x6b\x6a\x37\x21\xad\x7b\xea\x4c\xd2\xb4\xed\xc8\x95\xbf\x83\x1d\x1c\xe1\xa5\xd0\xc6\x82\x6c\x2f\x68\x0d\x98\xc5\x01\xc4\x21\x2c\x70\x25\x24\xdd\x4f\xc9\x6b\xfb\x45\xa0\xb6\xf6\x84\x5b\x69\xe4\xd6\x5d\x4f\xb9\xa4\x5b\x30\x7e\xcf\x79\x4a\x97\x99\x57\xc6\x72\x6d\x1b\x2a\x4e\x5c\x39\xee\x11\xd4\x97\x3c\x77\xc6\x17\x08\x42\x5a\xd4\x99\x46\x52\x11\x4e\xe4\xce\x94\x7b\x44\x3e\xfe\x8d\x5a\x75\x1e\x6a\x3b\xb5\x04\x09\xee\x7b\xc1\x49\x48\xda\x7e\xea\xd7\x3b\xa6\xcc\x53\xae\x57\x68\x2c\xb9\xcb\x94\x31\x62\x91\x62\xed\xf5\x88\x9a\x8f\x35\x30\xa8\x93\x7f\xd5\xbb\x5c\x49\x0a\x32\x82\x23\xde\xd6\xe2\xd0\x67\xab\x17\xdb\x49\x9a\xb6\xef\xce\xd6\xeb\x11\xd7\xc2\xba\x47\x21\xc8\x11\xab\xda\x1b\xaa\x7f\x2f\x56\x35\xbc\x8d\xed\x4f\xdc\xc6\x6b\x21\x57\x65\xd9\xbd\x93\x6b\x2f\xa7\xc2\xdd\x22\xed\x50\x3c\xb5\xa8\xdb\xe0\x89\xe0\x33\xe6\xb0\xf5\x11\xc8\x80\x2e\xc7\x57\x0f\x99\x6e\xc4\xd6\xae\x51\x68\x38\x7a\x97\xc2\xd6\xfd\x68\x30\x9b\xaf\x1b\xd5\xc2\x04\x7a\x6f\x7c\x10\x06\x78\xaa\x91\x27\x7b\x30\x4a\x9f\x4e\x6f\x3f\x50\x62\x50\x1c\x66\x37\x82\xe0\x70\xbc\xe8\x2b\xc7\x73\x9a\xf0\xfa\xdd\x8b\xaa\xd0\xe6\xf0\x92\x3c\x3c\x3e\x0b\x1d\x4c\x81\xcf\x48\x8c\x1f\x5f\x28\xd9\x73\xe0\x59\x86\x32\x39\x1e\xaf\xa3\x28\x7a\x6a\x9c\x6a\x19\x87\x26\xec\x0b\x4e\x37\x9d\xb0\x43\x1e\x7d\xda\xff\x30\x83\xc8\xfc\x69\x12\x29\x9d\xa0\xff\x8e\xe7\x07\x86\x1e\x79\xfc\x9e\x8e\x43\xde\xd7\x33\x34\xba\x45\xee\x5e\xf8\x4e\x8e\x0c\x70\x0b\x71\xae\x8d\xd2\xf5\xab\x0a\x65\x62\x60\xb7\x46\x59\x1f\x71\x94\x2b\xbb\x6e\x3e\x4b\x1e\x91\x8f\x9c\x99\x86\x80\x9d\x86\xc8\x08\xbe\x91\xbd\xf6\x71\x84\x71\x5f\x49\xdd\xd7\x50\xb4\x18\xfa\x70\xf4\xdc\x5f\x36\xc0\xe4\xf1\xba\x1e\x4d\xb8\x85\xdc\x38\x2b\xf7\x09\x95\x83\xc9\x17\xf8\x3d\x47\x69\x21\xe6\xa9\xd3\x25\xd7\xe0\x66\xb4\x51\x79\x9a\x34\x27\x4c\xe2\x83\x05\x37\xe2\x34\xdd\x7d\xe2\x1c\x3c\x0b\x51\xe0\xd3\x23\x69\x72\x4a\x71\xe1\xbb\xf3\x83\xfa\xd4\x05\x6b\x23\x39\x77\xc1\x53\xcc\x0f\xe1\x44\xbd\x1a\x60\x6a\xf9\xaa\x3f\x93\x35\x7f\xff\x13\x00\x00\xff\xff\x38\xc7\x5a\x59\xe3\x16\x00\x00")

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

	info := bindataFileInfo{name: "templates/kvschema.gotmpl", size: 5859, mode: os.FileMode(420), modTime: time.Unix(1566671246, 0)}
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
