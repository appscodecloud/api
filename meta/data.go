// Code generated by go-bindata.
// sources:
// meta/config.yaml
// DO NOT EDIT!

package meta

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

var _metaConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x94\xc1\x8e\xda\x30\x10\x86\xef\x79\x8a\x11\x9c\x9b\x88\x1e\xb9\x51\xba\x6a\x91\xb6\xd2\xaa\x6c\x7b\x9f\xd8\xb3\x64\x84\x33\xb6\x3c\xe3\x54\xbc\x7d\x15\x10\x88\xd5\x02\x82\x3d\x45\x89\xbf\x6f\x66\xec\x3f\xc9\x14\x96\x51\xde\x78\x53\x32\x1a\x47\x51\x78\x8b\x19\x16\x2f\xab\x27\xf1\x29\xb2\x98\xd6\xe3\xdd\x9a\xf2\x40\x19\xfe\x71\x08\xd5\x14\xd0\x19\x44\x01\x16\xcf\x03\xfb\x82\x01\x7a\xb2\x2e\x7a\x85\x92\xa2\x80\x75\x04\xee\x5d\xd1\x6a\x0a\x29\xc7\x81\x3d\xf9\xba\xaa\x6c\x97\x68\x0e\x93\x95\x18\x65\x47\xc9\x62\x3e\x8c\x30\xa9\x06\x0c\xec\xf7\xca\xd9\xe2\xbc\x02\x88\xd9\x53\x9e\xc3\xac\x02\x90\xf8\xf7\x44\x8d\x4b\x00\x5f\x60\xd2\x60\x4a\xea\xa2\xa7\xba\x23\x0c\xd6\xd5\x3f\xf7\x97\x66\x6d\x68\x45\x27\x15\x16\xeb\x48\x8c\xdd\x8d\xe2\x5f\x2b\x80\x22\x67\x24\xf9\xe3\x21\xdc\xdd\xe7\x03\x36\x96\xab\x87\x59\x4b\x86\xb3\x7a\xf1\x6e\x8a\xe6\x39\x6e\x58\x1e\x77\x62\xb1\x07\xa5\xd7\xb8\xa5\x4b\x8d\xb6\xa5\xa5\x2c\x64\xa4\x27\x73\x19\x8a\x1a\x65\x1d\x77\x94\xad\xa4\x63\x32\x0f\xb9\x2b\x51\x43\x71\xf4\x6d\xb7\x7a\xb9\xa0\xf6\xc8\x81\x65\x13\x58\xed\xe4\xfe\x3a\x3c\x7b\x66\xb5\x66\x4d\xe2\x9f\x46\xe8\x33\x6e\x69\xd5\x65\x6e\xe9\x13\xee\x1f\xd1\x1b\xb6\x60\x4f\x9a\xd0\xd1\xc9\x7d\x25\xec\xb5\x59\x66\x42\x7b\x44\xf8\x41\x97\xf2\xbb\x46\xaf\x74\x31\x20\x07\x6c\xc3\xbe\x87\xee\xd4\xa8\x5f\xdc\xf9\x92\x3a\x3e\xdb\xe8\x98\xcd\xf5\x69\x3f\xa2\xdf\x29\xd0\x45\xb4\x45\xb7\x2d\xe9\x2c\x76\x26\x31\x6d\x7e\xd3\xf1\xa3\xbf\x43\x3a\xfc\x50\x6e\x9c\xde\x15\xfe\x38\xd4\xff\x00\x00\x00\xff\xff\x3d\x0d\xe9\x37\xb7\x04\x00\x00")

func metaConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_metaConfigYaml,
		"meta/config.yaml",
	)
}

func metaConfigYaml() (*asset, error) {
	bytes, err := metaConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta/config.yaml", size: 1207, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"meta/config.yaml": metaConfigYaml,
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
	"meta": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{metaConfigYaml, map[string]*bintree{}},
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

