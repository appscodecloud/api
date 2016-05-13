// Code generated by go-bindata.
// sources:
// data/files/ci_products.json
// data/files/cloud_providers.json
// data/files/db_products.json
// data/files/kubernetes_products.json
// data/files/subscriptions.json
// DO NOT EDIT!

package files

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _dataFilesCi_productsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x94\xb1\x6a\xc3\x30\x10\x86\x77\x3f\x85\xd0\x9c\x94\x3b\xc7\x36\x26\x6b\xa7\xae\xa5\x53\x4b\x09\xaa\x75\x04\x51\x57\x0e\xb2\x5c\x0a\x21\xef\xde\x53\xa8\xeb\x14\x67\x30\x41\x14\x81\xd0\xf0\xe9\xa4\x5f\x7c\x1c\x77\xcc\x84\x90\x6f\x83\x69\xf5\x4e\xed\xc9\xfa\x5e\x6e\xc5\x0b\x33\x21\x8e\xe7\x9d\x4f\xfb\xf7\x81\xa1\xbc\x7f\x58\x3f\x96\x98\xcb\xd5\xc8\x0f\xae\xfb\x34\x9a\x5c\x38\xd4\x66\x6f\xbc\x6a\xbb\x86\x94\x9d\x2a\xba\xf0\x9a\x6c\x8d\x1d\xbe\x2e\xaf\x99\x86\x98\x6f\xca\x5f\xa4\xc9\x2b\xd3\x86\xe2\x31\x94\x61\x73\x08\xb1\xb8\x9a\x88\x53\x1f\x4c\xf8\x0f\x3f\xe8\x34\x3d\xa0\x3c\xed\x7a\xaf\x9c\x27\x1d\x22\x73\xc0\x6a\x0d\xc8\xeb\x09\x60\x7b\x5e\x77\x00\xf0\x2c\xff\xde\x20\xab\xc7\xfa\x0d\x5c\xad\xcf\x2e\x72\xae\x09\xc1\x68\x3a\xca\x1b\x75\x60\x3a\x32\xe2\xf5\x06\x02\x2c\xb1\x91\xcf\x6c\x24\xd4\x1a\x45\x3c\x1b\xd5\x8d\x36\x8a\x74\x6c\xd4\xd1\x6c\xe4\xe5\x22\x1b\xc5\xcc\x46\x9d\x8e\x0d\xac\xe2\x8d\xd1\x65\xcd\x51\xcf\x07\x47\xf5\xdf\x3e\x78\x7f\xcd\x4e\xd9\x77\x00\x00\x00\xff\xff\x47\x15\xcd\xef\x73\x06\x00\x00")

func dataFilesCi_productsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesCi_productsJson,
		"data/files/ci_products.json",
	)
}

func dataFilesCi_productsJson() (*asset, error) {
	bytes, err := dataFilesCi_productsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/ci_products.json", size: 1651, mode: os.FileMode(436), modTime: time.Unix(1463142321, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataFilesCloud_providersJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x99\xcf\x6e\xe3\xb6\x13\xc7\xef\x79\x0a\x42\x67\x5b\xb0\x28\xad\x63\xec\xed\x17\x63\x91\x5f\x80\x16\x1b\xc0\x9b\xf6\x50\x2c\x16\x5c\x89\x91\x89\x48\xa4\x40\x49\x4d\xbd\x8b\x3c\x4d\x1f\xa5\x2f\x56\xfa\x1f\x63\x59\xa4\xcd\x11\x4c\xd4\x37\x99\x43\xd1\x9f\xd1\x7c\x67\x24\x72\x7e\xde\x20\x14\xa4\x85\x68\xb3\x6f\x95\x14\x7f\xb2\x8c\xca\x3a\xf8\xb8\x1e\x55\xe3\xe4\x55\x5f\xab\x5f\x9c\x94\x34\xf8\xb8\x19\x1c\xed\xc7\x24\xcd\x99\xe0\x6a\xd6\x1f\xbb\x11\x84\x7e\xea\x2b\x65\x2f\x44\x4a\x1a\x35\x43\xdd\xf7\x1b\x93\x39\xe3\x8c\xe8\x9b\x0f\x16\x50\xe6\xb6\x1e\x53\x52\x37\xe3\xa8\x6b\xff\x21\x38\xed\x2c\xbf\x19\xd5\x93\xbb\xab\x1d\x5a\xbe\x5b\x2d\xa9\xd5\x92\x59\x2d\x34\x38\x30\x7c\xd5\xd7\x6f\xa3\xb3\x6e\xcf\x49\xc1\x9e\x85\x3c\xe9\xf8\x2b\x05\x38\xbe\x9d\x6c\x72\x7c\x6b\x31\x39\xbe\xb5\xa4\x83\x9d\xf8\xac\x68\xd5\xc5\x69\x07\x70\xdf\x01\x2b\x3f\xb6\xf2\x63\x2b\x3f\x1e\xce\xff\x20\x69\x41\x78\x66\x71\x80\xb6\x80\x08\xe8\xc9\x3d\x0f\xb4\xa5\xe7\x81\xb6\x0c\xf7\x60\xc1\x78\x4e\x2a\x21\xa9\xc5\x07\x52\x8d\x6b\xd1\x36\x4b\xf7\x1c\xea\xde\xd1\xf3\xa6\x6b\xfe\x3e\x18\xfc\x8b\x78\x59\x09\x3b\x34\x17\x12\x08\x7d\x70\x87\x09\xfa\xc0\xdc\x8b\x43\xd7\x0c\x0c\x86\xf9\x51\x1f\x89\xfe\x30\x62\xab\x8c\xd3\x15\x38\x12\xfd\xcc\xe8\x9a\x81\x91\xd0\xd4\x35\x31\xd6\xd7\x43\xe0\x7f\xfe\x16\xe8\x91\xb4\x85\x70\x81\xd6\xeb\xf5\x78\xb5\xc5\x86\xba\xbb\xfa\xaa\xdf\x22\x8c\xd7\x0d\xe1\x29\xfd\xd6\xac\x2a\x6a\x7f\x99\x28\x52\x55\x86\xe4\x4a\xb1\xde\x53\x4e\x25\x29\xd0\x63\x2b\x2b\x51\x1f\x25\x45\xfd\xd2\xaa\x29\x0d\x0e\x4b\x96\xca\x23\x67\xd2\x4a\xd9\xa2\xce\x50\x49\xcb\xcd\xa2\xd1\xd9\x87\x09\x06\xa8\x4b\x52\x14\xce\x00\xf8\xf2\x00\x25\xcd\x58\x5b\x1a\x08\xb0\x91\x20\xb9\x3c\x41\x41\x64\x4e\x9d\x01\x66\x17\x06\x28\x93\xff\x1e\xe0\x2f\x1b\x41\x62\x96\xe1\xf4\xf2\x08\xd8\xca\x30\x33\x32\xc4\x97\x96\xa2\x62\x48\xac\x0c\xd1\xd4\x08\x31\xbd\xb4\x1a\x15\x44\x34\xb1\x47\x63\x62\x09\xc7\xe4\xd2\x18\xb1\x3d\x2d\xcd\x85\x21\x0e\x6f\x3f\x5c\x1e\x02\x96\x18\xb7\xa1\x07\x04\x68\x6a\x78\x40\x00\xa7\x06\x48\x0e\x73\x51\x56\x6d\x43\xd1\xe7\xaa\x61\x25\xfb\x41\x33\x13\x45\x0a\x2d\x52\x50\x3d\xb8\x52\x00\xc3\x01\x94\x84\x2b\x05\x34\x22\x30\x55\xb8\x52\x80\xcb\x95\x1f\x61\xcc\xac\x18\xb1\xa5\x6a\x7a\xc0\x80\xd6\x0a\x3f\xfa\x04\x97\x0b\x2f\xfa\x84\x57\x0c\x1f\xfa\x8c\xaf\x43\x9f\xf1\x29\x7d\x9a\xb5\x01\x7c\x9d\x3e\x3e\xa1\x87\xdd\xee\xa0\x36\x21\xe4\xd8\x73\x3c\x5c\x08\x3c\x3f\x84\x5f\x37\x77\x21\x71\x2a\x14\x12\x9a\xa3\xb0\xa7\xe0\x88\x00\x4c\xd0\x78\x02\xcb\x50\x47\x0a\xa8\x20\xa6\xa0\x9d\x9f\x23\x04\x38\x3f\x23\x0c\xfa\xe8\x76\xc4\x00\x4b\x13\x27\xa0\xcf\xee\x45\x23\x24\xc9\xcf\xd4\x09\x86\x7d\x2b\xc3\x15\xc3\xaf\x34\x5c\x29\x7c\x6b\xc3\x95\xe3\x2a\xc4\x91\x5d\x87\x38\xb2\xab\x10\x47\x76\x25\xe2\xc8\x4e\x8a\xc3\xcc\xd1\x11\xc7\xee\x6a\x7b\xea\xb7\xa3\x0a\xf2\x94\xf6\xfb\x49\xeb\xc1\xfd\x7a\x67\xfb\x49\xfa\x38\xf3\x13\xa9\x1b\x2a\x39\x7a\x5a\x74\x01\x7f\x79\x3f\xcf\xbc\xa3\xf2\x85\x16\x74\x85\xe6\xa2\xe5\xcd\x6a\x84\x16\xeb\x53\x54\x34\x27\x52\x14\x8c\x13\x40\x9f\x29\x1a\xdb\xba\x49\xd1\xd8\xd6\x4d\x8a\xc6\xd9\xc0\x93\xda\x39\xe5\xcd\x7a\xdf\x7a\xc2\xb5\xb5\x47\x29\x2b\xd0\x5d\xd1\x3e\x3f\xd7\x23\xf4\x20\x5e\x5d\xfd\x49\xb7\xab\x47\x63\x53\x07\x46\x1b\x4d\xfe\x6a\xa3\xc9\x65\x6d\x7c\x1e\xe8\xf5\xef\x74\x1b\xd0\x4f\xad\x14\x15\xb5\x7a\xbe\x68\x42\x74\xbf\x64\x75\x41\x18\x1f\xa1\x3b\x5a\xe4\xbd\x83\x14\x5b\xdf\x66\xbd\xee\xa6\x0f\x63\x70\xaf\x63\xed\xf9\xd7\xb1\x0e\x0d\xeb\x5a\xb1\xe8\x7f\xf5\x71\x1f\xf0\x30\xaa\x4b\xc2\xf3\x65\x4b\xb4\x60\xbf\x10\xf6\x4a\xb8\x53\x0b\x41\xad\xbb\x93\x5d\xbf\x7f\xf0\x6e\xeb\x77\x44\xde\x6d\xb6\x76\xc8\x3e\x93\x75\x8e\xc2\x4f\xeb\x17\xfd\x63\xf0\xdd\xb7\x73\x34\x86\x1d\x91\x07\x51\x78\x3b\xb9\xbf\x0b\x60\x05\x8f\xf0\x8c\x48\x63\x9d\xe3\x0a\x60\x67\x3e\xee\x8c\x9c\xa2\x58\x6f\x70\x3d\x51\x60\xe7\x2f\xf8\x40\xed\x6f\x3d\x41\x24\xce\xef\xe2\x20\xf2\xc5\x30\x73\x7e\x0f\x07\xb1\x37\x49\x4c\xdd\x5f\xc3\xc1\xd4\x17\x45\x8c\x4d\x18\x96\x4f\xb4\x20\xc2\x50\x8e\xff\xb3\x7c\x89\xb6\xfb\x09\x0b\xca\x52\xcd\x50\x7f\x00\x11\x67\x14\x7b\xa3\x00\xa8\x13\x4f\xbd\x51\x00\xf4\xf9\x01\x7b\xa3\x00\x29\x34\x9a\x24\xde\x40\x60\x22\xc5\x93\xd9\x10\x92\xf9\xe3\xd3\x09\x0c\xf5\x8f\x20\x85\x86\xb3\x41\x99\x72\x1e\x02\x20\xd0\x38\x04\x97\x0d\x47\x08\x80\x3e\x6f\xc3\x61\x35\xe3\x3c\x04\x4c\x9e\x49\x98\x78\xe2\x00\xaa\x73\x76\x2c\x8d\xfd\x47\xd0\xcd\xfe\xd7\xdb\xcd\xdb\xcd\xbf\x01\x00\x00\xff\xff\x26\x57\x9f\x74\x3c\x27\x00\x00")

func dataFilesCloud_providersJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesCloud_providersJson,
		"data/files/cloud_providers.json",
	)
}

func dataFilesCloud_providersJson() (*asset, error) {
	bytes, err := dataFilesCloud_providersJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/cloud_providers.json", size: 10044, mode: os.FileMode(436), modTime: time.Unix(1463034467, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataFilesDb_productsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x93\x5f\x4f\x83\x30\x14\xc5\xdf\xf9\x14\xa4\xcf\xa3\xb9\x40\xa6\x91\x37\xfe\x4c\x5d\x42\xd4\x38\x49\x8c\xc6\x2c\x1d\x54\x25\x22\x90\xb6\x18\xcd\xb2\xef\x6e\x65\x5d\x64\xac\xf3\x61\x0f\x2a\x49\x13\xe0\xf4\xde\xf2\xbb\x3d\x39\x4b\xc3\x34\x51\x5d\x71\xf1\xc4\x28\x47\xde\xd7\xa7\x14\x4a\xf2\x4a\x91\xf7\xad\x8f\xd6\xf2\x1b\x65\x3c\xaf\x4a\x59\x77\xdf\x0a\x52\x3a\xc1\x63\xd4\xbe\x3f\xa8\x9a\x6c\x31\x17\x1f\x35\xed\xd4\x2c\xd5\x53\x6e\xf2\x97\x46\x1e\x1b\x05\xd6\xd5\x99\x15\xf8\xb3\x69\xa8\x4e\x6e\x37\x6b\x96\xa7\xf2\xaf\xd0\x91\x14\x47\x40\x78\x9e\x76\x4b\x33\x22\xe8\x9c\x0b\xc2\x04\xcd\x90\x67\x22\x07\xec\x23\x0b\x6c\xb9\x6e\x00\xbc\x76\x61\x00\xb8\xdb\xe9\xa1\x65\xb6\xe9\x70\x41\xdb\xa1\x1a\x56\xa3\x9f\xe9\xa3\xe4\xda\x0f\xe2\x89\x86\x7f\xbc\xcb\x1f\x35\x8c\x2c\x0a\xfa\xbf\x26\x38\xf7\x35\xf0\x8e\xe6\xf6\xb7\x0b\xff\x1e\x3c\x8c\x2f\x93\x48\xc7\xae\xb9\xf9\xb0\xa8\x9a\xec\x97\xf1\xd7\x71\x30\xd4\x20\x88\x16\x84\x8b\x3c\xe5\x94\xb0\xf4\xb9\x9f\xb0\xed\xcd\xfd\x31\xb3\xf1\xb1\x0c\xda\x86\x09\x39\xd8\xc5\xf6\x21\xb9\x9b\xcc\x86\x9c\x3b\x49\x3f\x80\xdc\xf5\xfc\xcf\xcb\xc7\xa2\x79\xef\x1b\xaf\xd4\xfd\x8e\x03\xb6\x01\xbb\x87\x78\x3c\xbd\x38\x8d\x93\xdb\x21\xfb\xac\x26\x18\x94\xd7\xc6\xca\xf8\x0c\x00\x00\xff\xff\xe9\x6e\xb0\x0c\x4c\x07\x00\x00")

func dataFilesDb_productsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesDb_productsJson,
		"data/files/db_products.json",
	)
}

func dataFilesDb_productsJson() (*asset, error) {
	bytes, err := dataFilesDb_productsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/db_products.json", size: 1868, mode: os.FileMode(436), modTime: time.Unix(1463143133, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataFilesKubernetes_productsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x50\x50\xca\x2e\x4d\x4a\x8d\x4f\x4c\x4f\xcd\x2b\x29\x56\xb2\x02\x0a\x28\x28\x54\x83\x49\xa0\x54\x71\x76\xa9\x92\x95\x92\x9f\xbf\x8b\xab\xae\xb3\x92\x0e\x4c\xb4\xa0\x28\x33\x39\x55\xc9\xca\x5c\xcf\x14\x2e\x94\x5c\x00\x54\x68\x08\xe7\xa6\x24\x96\xa4\xc6\x17\x97\x24\x16\x95\xa4\xa6\x28\x59\x29\x28\x19\x19\x18\x9a\xe9\x1a\x18\x02\x51\x88\x81\x81\x15\x18\xe9\x19\x18\x18\x44\x29\xa1\xea\x48\xcd\x4b\x81\xa9\x37\x36\xc0\xaa\x1e\xac\xbc\x96\xab\x96\x0b\x10\x00\x00\xff\xff\xb5\x23\xe4\xd5\xba\x00\x00\x00")

func dataFilesKubernetes_productsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesKubernetes_productsJson,
		"data/files/kubernetes_products.json",
	)
}

func dataFilesKubernetes_productsJson() (*asset, error) {
	bytes, err := dataFilesKubernetes_productsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/kubernetes_products.json", size: 186, mode: os.FileMode(436), modTime: time.Unix(1463045973, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataFilesSubscriptionsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\x51\x6b\xab\x30\x14\xc7\xdf\xfd\x14\x92\xe7\x7a\x49\x52\xec\x05\xdf\xfa\x72\xb9\x7b\x28\x14\xd6\xa7\x8d\x51\x52\x3d\x6d\x43\x9d\x4a\x12\x61\x5b\xe9\x77\x5f\xb4\xb4\xb5\xd6\x38\x7d\x6b\x67\x40\x84\x1d\xcf\x3f\x47\xdc\x8f\xf3\x83\xee\x1d\xd7\x45\x99\xe0\xef\x4c\x7c\xa2\xe0\x55\xff\xe5\xba\xfb\xf2\xae\xeb\x72\x97\xa3\x00\xfd\xf7\x66\x74\xe6\xfd\x13\x00\x68\x74\x7a\xa2\x13\x21\xa0\x00\x9f\x0b\x11\x28\xc6\x63\x89\x82\x53\x58\xd7\x72\x09\x02\x05\xe3\xd1\xa5\x92\x6d\xd9\x4a\x27\x99\x4a\x45\xb5\xb3\xc8\x73\xb9\x5b\x4a\xfe\xa5\x0f\xa5\xe7\xfa\xa1\x12\x65\x42\xf1\x35\x0b\x95\x39\x47\x70\x63\x70\x97\xaf\xa0\x16\x0a\xe3\x5c\xaa\xe2\xdd\xc8\xa8\x5a\x4e\xd2\x08\xa4\x69\x7e\xc8\x6b\x87\xb0\x0d\x24\xfa\x6d\x9a\x87\x46\xab\x5a\x77\x96\x4a\xb5\x11\xc5\xf1\xf8\x6a\x26\xc4\x4c\x2a\x1e\x4a\x60\x22\xdc\xd6\x1f\xf2\x64\x1d\xe7\x1f\xc5\x59\x95\x29\x4e\x6d\x1a\x8a\x98\x82\xa5\x54\xfa\x03\x41\x84\x02\x17\x51\x4c\x26\x1e\x26\xfa\x5a\x60\x1c\x94\xd7\x1f\x8c\xf1\x0b\xba\x4e\x40\x12\x9d\xfa\xc7\xb8\xb1\xdf\xa9\xcc\x69\x64\xe2\x29\x89\xf8\x2d\x14\x84\x76\xc1\x82\xe0\xde\x5c\x34\x7f\xea\x2e\x60\xf4\x27\xc3\x6b\x44\xc3\x23\xbd\xd8\x30\xb4\xb7\xc0\xe1\xb7\xc1\xe1\x1b\xe0\xf0\xef\x12\x8e\xe7\x62\x68\x9e\xdd\xe2\xf1\xb7\x0b\x1e\xd4\xef\x8b\x87\xe9\xbf\xfc\x33\x1f\xfe\xe3\xf0\x41\x5a\xb7\x07\x31\xad\x0f\x72\x9f\xfb\x63\x2e\xd2\x35\x48\xc9\xd3\x84\xc5\x37\x9c\xd0\x8e\x6b\xa4\xff\x1e\xf1\x07\x00\x0a\x6d\x05\x85\x9a\x40\xa1\xf7\x07\xca\x74\x6e\x14\x0d\xb1\xa2\x19\xb8\x68\x0a\x38\x8c\xa2\x99\x58\xd1\x58\xd1\x94\x88\xb4\x8b\xa6\xe3\x1a\xb1\xa2\xb9\xb4\xff\x4e\xd1\x4c\x4d\xa2\xc1\x56\x34\x83\x17\xcd\xd4\x2c\x1a\xdf\x8a\xc6\x8a\xa6\x44\xa4\x5d\x34\x1d\xd7\x88\x15\xcd\xa5\xfd\xa1\x45\xa3\xef\x6f\x45\x12\x29\xc1\x35\x0f\xcd\x3f\xaf\x2e\x84\x06\xc1\x20\x1e\xab\x9d\x41\x6b\xe7\x88\x86\x49\x3b\x56\x3a\x83\x97\xce\x11\x90\x56\xe9\x58\xe5\x0c\x4e\x39\xce\xc1\xf9\x0e\x00\x00\xff\xff\xde\xb2\x6f\x62\xd7\x1b\x00\x00")

func dataFilesSubscriptionsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesSubscriptionsJson,
		"data/files/subscriptions.json",
	)
}

func dataFilesSubscriptionsJson() (*asset, error) {
	bytes, err := dataFilesSubscriptionsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/subscriptions.json", size: 7127, mode: os.FileMode(436), modTime: time.Unix(1463047210, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/files/ci_products.json": dataFilesCi_productsJson,
	"data/files/cloud_providers.json": dataFilesCloud_providersJson,
	"data/files/db_products.json": dataFilesDb_productsJson,
	"data/files/kubernetes_products.json": dataFilesKubernetes_productsJson,
	"data/files/subscriptions.json": dataFilesSubscriptionsJson,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"files": &bintree{nil, map[string]*bintree{
			"ci_products.json": &bintree{dataFilesCi_productsJson, map[string]*bintree{
			}},
			"cloud_providers.json": &bintree{dataFilesCloud_providersJson, map[string]*bintree{
			}},
			"db_products.json": &bintree{dataFilesDb_productsJson, map[string]*bintree{
			}},
			"kubernetes_products.json": &bintree{dataFilesKubernetes_productsJson, map[string]*bintree{
			}},
			"subscriptions.json": &bintree{dataFilesSubscriptionsJson, map[string]*bintree{
			}},
		}},
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

