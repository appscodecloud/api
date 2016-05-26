// Code generated by go-bindata.
// sources:
// data/files/ci.latest.json
// data/files/cloud_provider.json
// data/files/cluster.latest.json
// data/files/configs.yaml
// data/files/db.latest.json
// data/files/package.latest.json
// DO NOT EDIT!

package files

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

var _dataFilesCiLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x96\xdf\x6b\xdb\x30\x10\xc7\xdf\xf3\x57\x18\x3d\x2f\x45\x52\x1c\x63\xf2\x96\xb5\x63\x04\xd6\x75\x98\xbd\x74\x63\x08\xc5\x3a\x52\x51\x5b\x32\xfa\x31\x52\x4a\xff\xf7\x49\x5e\x9b\x7a\x55\x56\x0a\x85\x24\x60\x30\xc2\xdc\x7d\xef\x7b\x3a\xf1\x41\xe8\x7e\x92\x65\x68\xed\x65\x23\x18\xdf\x80\x72\x68\x91\xfd\x0c\xa1\x2c\xbb\xef\xd7\x90\xb4\xb7\x3e\x04\xd1\xf9\x6a\x7a\x4e\x2a\x32\xbd\xb8\x42\x1f\x9e\x52\x42\xda\xae\xe1\x77\x4c\xf1\x16\xa2\xe6\x63\xf4\xc9\x96\xbd\xcf\x4e\xd4\x19\x59\x4b\xb5\x61\xad\x16\xd0\xa0\x05\xfa\xb6\xbc\xfe\xfc\x9c\xb5\x7e\x6d\x6b\x23\x3b\x27\xb5\x62\xee\xae\x83\xa8\xa8\x56\x97\xcb\xea\x7a\x8f\x05\xb8\xf0\x17\x15\x9f\x2a\x76\xb9\xfa\xfa\xac\xf0\x4a\x3a\x16\x65\xc0\xbc\x15\x41\x81\xcf\x48\xba\xcf\x81\xe0\x69\xbc\xe8\x0e\x86\xdd\x68\x6f\x42\xd5\x6c\x57\xf3\x18\x6f\xb5\x72\x37\x21\x31\x9f\xa3\xc7\xc4\xc3\xce\x35\xec\x86\x0b\xee\x78\x98\x7c\xe8\x66\xf4\x6f\x29\xc0\xc4\xf3\xb8\x90\x1b\xe9\x78\x73\x55\x03\x57\x43\x67\x6d\x63\xf6\x8b\x54\x7e\x3b\x0c\xd7\x5d\x3c\x69\x32\x88\x18\xde\xbe\x88\x84\x51\x6e\x43\x68\x86\x87\x2a\xa8\x75\xdb\x82\x12\x20\x18\x6c\xa1\xf6\x4e\x1b\xfb\xa2\x0e\xb6\x0e\x8c\xe2\x0d\x93\x22\xf6\x26\x9b\x75\x3a\x50\x18\x06\x98\x75\xdc\x38\xe8\x45\x14\x93\x62\x8a\x49\xf8\xbe\x63\xbc\xe8\xbf\x33\x8c\xf1\x0f\xf4\x6f\x45\xdf\xf8\xaf\x7e\x86\xf7\xea\x27\x83\x3e\xfb\xb8\xa2\x15\x1d\x29\x57\x04\xe3\x83\x81\x45\x13\xb0\x68\x0a\x56\xfe\x16\xb0\xe8\xff\xc1\xa2\x27\x06\x56\x3e\x56\xb0\x8a\x63\x82\x95\xa7\x60\x15\xef\x04\x2b\x3f\x29\xb0\xf2\xaa\x1c\x29\x58\x74\x7e\x38\xb0\xf2\x04\xac\x32\x05\xab\x7c\x27\x58\xe5\x49\x81\x55\x56\xa1\xd1\x38\xc9\x9a\x1d\xf0\xca\x2a\xd3\x47\x56\x91\xa2\x45\xde\x74\x69\xe5\xaf\x3c\xb3\x8a\x23\xc0\x15\xd6\x5f\x93\x87\xc9\x9f\x00\x00\x00\xff\xff\xc9\xc8\x18\x23\xda\x0b\x00\x00")

func dataFilesCiLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesCiLatestJson,
		"data/files/ci.latest.json",
	)
}

func dataFilesCiLatestJson() (*asset, error) {
	bytes, err := dataFilesCiLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/ci.latest.json", size: 3034, mode: os.FileMode(420), modTime: time.Unix(1464258041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFilesCloud_providerJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x99\xcf\x6e\xe3\xb6\x13\xc7\xef\x79\x0a\x42\x67\x5b\xb0\x28\xad\x63\xec\xed\x17\x63\x91\x5f\x80\x16\x1b\xc0\x9b\xf6\x50\x2c\x16\x5c\x89\x91\x89\x48\xa4\x40\x49\x4d\xbd\x8b\x3c\x4d\x1f\xa5\x2f\x56\xfa\x1f\x63\x59\xa4\xcd\x11\x4c\xd4\x37\x99\x43\xd1\x9f\xd1\x7c\x67\x24\x72\x7e\xde\x20\x14\xa4\x85\x68\xb3\x6f\x95\x14\x7f\xb2\x8c\xca\xe0\xe3\x7a\x50\x0d\x93\xd7\x7a\x7f\xad\x7e\x71\x52\xd2\xe0\xe3\x66\x70\xb4\x1f\x93\x34\x67\x82\xab\x59\x7f\xec\x46\x10\xfa\xa9\xaf\x94\xbd\x10\x29\x69\xd4\x0c\x75\xdf\x6f\x4c\xe6\x8c\x33\xa2\x6f\x3e\x58\x40\x99\xdb\x7a\x4c\x49\xdd\x8c\xa3\xae\xfd\x87\xe0\xb4\xb3\xfc\x66\x54\x4f\xee\xae\x76\x68\xf9\x6e\xb5\xa4\x56\x4b\x66\xb5\xd0\xe0\xc0\xf0\x55\x5f\xbf\x8d\xce\xba\x3d\x27\x05\x7b\x16\xf2\xa4\xe3\xaf\x14\xe0\xf8\x76\xb2\xc9\xf1\xad\xc5\xe4\xf8\xd6\x92\x0e\x76\xe2\xb3\xa2\x55\x17\xa7\x1d\xc0\x7d\x07\xac\xfc\xd8\xca\x8f\xad\xfc\x78\x38\xff\x83\xa4\x05\xe1\x99\xc5\x01\xda\x02\x22\xa0\x27\xf7\x3c\xd0\x96\x9e\x07\xda\x32\xdc\x83\x05\xe3\x39\xa9\x84\xa4\x16\x1f\x48\x35\xae\x45\xdb\x2c\xdd\x73\xa8\x7b\x47\xcf\x9b\xae\xf9\xfb\x60\xf0\x2f\xe2\x65\x25\xec\xd0\x5c\x48\x20\xf4\xc1\x1d\x26\xe8\x03\x73\x2f\x0e\x5d\x33\x30\x18\xe6\x47\x7d\x24\xfa\xc3\x88\xad\x32\x4e\x57\xe0\x48\xf4\x33\xa3\x6b\x06\x46\x42\x53\xd7\xc4\x58\x5f\x0f\x81\xff\xf9\x5b\xa0\x47\xd2\x16\xc2\x05\x5a\xaf\xd7\xe3\xd5\x16\x1b\xea\xee\xea\xab\x7e\x8b\x30\x5e\x37\x84\xa7\xf4\x5b\xb3\xaa\xa8\xfd\x65\xa2\x48\x55\x19\x92\x2b\xc5\x7a\x4f\x39\x95\xa4\x40\x8f\xad\xac\x44\x7d\x94\x14\xf5\x4b\xab\xa6\x34\x38\x2c\x59\x2a\x8f\x9c\x49\x2b\x65\x8b\x3a\x43\x25\x2d\x37\x8b\x46\x67\x1f\x26\x18\xa0\x2e\x49\x51\x38\x03\xe0\xcb\x03\x94\x34\x63\x6d\x69\x20\xc0\x46\x82\xe4\xf2\x04\x05\x91\x39\x75\x06\x98\x5d\x18\xa0\x4c\xfe\x7b\x80\xbf\x6c\x04\x89\x59\x86\xd3\xcb\x23\x60\x2b\xc3\xcc\xc8\x10\x5f\x5a\x8a\x8a\x21\xb1\x32\x44\x53\x23\xc4\xf4\xd2\x6a\x54\x10\xd1\xc4\x1e\x8d\x89\x25\x1c\x93\x4b\x63\xc4\xf6\xb4\x34\x17\x86\x38\xbc\xfd\x70\x79\x08\x58\x62\xdc\x86\x1e\x10\xa0\xa9\xe1\x01\x01\x9c\x1a\x20\x39\xcc\x45\x59\xb5\x0d\x45\x9f\xab\x86\x95\xec\x07\xcd\x4c\x14\x29\xb4\x48\x41\xf5\xe0\x4a\x01\x0c\x07\x50\x12\xae\x14\xd0\x88\xc0\x54\xe1\x4a\x01\x2e\x57\x7e\x84\x31\xb3\x62\xc4\x96\xaa\xe9\x01\x03\x5a\x2b\xfc\xe8\x13\x5c\x2e\xbc\xe8\x13\x5e\x31\x7c\xe8\x33\xbe\x0e\x7d\xc6\xa7\xf4\x69\xd6\x06\xf0\x75\xfa\xf8\x84\x1e\x76\xbb\x83\xda\x84\x90\x63\xcf\xf1\x70\x21\xf0\xfc\x10\x7e\xdd\xdc\x85\xc4\xa9\x50\x48\x68\x8e\xc2\x9e\x82\x23\x02\x30\x41\xe3\x09\x2c\x43\x1d\x29\xa0\x82\x98\x82\x76\x7e\x8e\x10\xe0\xfc\x8c\x30\xe8\xa3\xdb\x11\x03\x2c\x4d\x9c\x80\x3e\xbb\x17\x8d\x90\x24\x3f\x53\x27\x18\xf6\xad\x0c\x57\x0c\xbf\xd2\x70\xa5\xf0\xad\x0d\x57\x8e\xab\x10\x47\x76\x1d\xe2\xc8\xae\x42\x1c\xd9\x95\x88\x23\x3b\x29\x0e\x33\x47\x47\x1c\xbb\xab\xed\xa9\xdf\x8e\x2a\xc8\x53\xda\xef\x27\xad\x07\xf7\xeb\x9d\xed\x27\xe9\xe3\xcc\x4f\xa4\x6e\xa8\xe4\xe8\x69\xd1\x05\xfc\xe5\xfd\x3c\xf3\x8e\xca\x17\x5a\xd0\x15\x9a\x8b\x96\x37\xab\x11\x5a\xac\x4f\x51\xd1\x9c\x48\x51\x30\x4e\x00\x7d\xa6\x68\x6c\xeb\x26\x45\x63\x5b\x37\x29\x1a\x67\x03\x4f\x6a\xe7\x94\x37\xeb\x7d\xeb\x09\xd7\xd6\x1e\xa5\xac\x40\x77\x45\xfb\xfc\x5c\x8f\xd0\x83\x78\x75\xf5\x27\xdd\xae\x1e\x8d\x4d\x1d\x18\x6d\x34\xf9\xab\x8d\x26\x97\xb5\xf1\x79\xa0\xd7\xbf\xd3\x6d\x40\x3f\xb5\x52\x54\xd4\xea\xf9\xa2\x09\xd1\xfd\x92\xd5\x05\x61\x7c\x84\xee\x68\x91\xf7\x0e\x52\x6c\x7d\x9b\xf5\xba\x9b\x3e\x8c\xc1\xbd\x8e\xb5\xe7\x5f\xc7\x3a\x34\xac\x6b\xc5\xa2\xff\xd5\xc7\x7d\xc0\xc3\xa8\x2e\x09\xcf\x97\x2d\xd1\x82\xfd\x42\xd8\x2b\xe1\x4e\x2d\x04\xb5\xee\x4e\x76\xfd\xfe\xc1\xbb\xad\xdf\x11\x79\xb7\xd9\xda\x21\xfb\x4c\xd6\x39\x0a\x3f\xad\x5f\xf4\x8f\xc1\x77\xdf\xce\xd1\x18\x76\x44\x1e\x44\xe1\xed\xe4\xfe\x2e\x80\x15\x3c\xc2\x33\x22\x8d\x75\x8e\x2b\x80\x9d\xf9\xb8\x33\x72\x8a\x62\xbd\xc1\xf5\x44\x81\x9d\xbf\xe0\x03\xb5\xbf\xf5\x04\x91\x38\xbf\x8b\x83\xc8\x17\xc3\xcc\xf9\x3d\x1c\xc4\xde\x24\x31\x75\x7f\x0d\x07\x53\x5f\x14\x31\x36\x61\x58\x3e\xd1\x82\x08\x43\x39\xfe\xcf\xf2\x25\xda\xee\x27\x2c\x28\x4b\x35\x43\xfd\x01\x44\x9c\x51\xec\x8d\x02\xa0\x4e\x3c\xf5\x46\x01\xd0\xe7\x07\xec\x8d\x02\xa4\xd0\x68\x92\x78\x03\x81\x89\x14\x4f\x66\x43\x48\xe6\x8f\x4f\x27\x30\xd4\x3f\x82\x14\x1a\xce\x06\x65\xca\x79\x08\x80\x40\xe3\x10\x5c\x36\x1c\x21\x00\xfa\xbc\x0d\x87\xd5\x8c\xf3\x10\x30\x79\x26\x61\xe2\x89\x03\xa8\xce\xd9\xb1\x34\xf6\x1f\x41\x37\xfb\x5f\x6f\x37\x6f\x37\xff\x06\x00\x00\xff\xff\xa3\x88\xf0\x08\x3b\x27\x00\x00")

func dataFilesCloud_providerJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesCloud_providerJson,
		"data/files/cloud_provider.json",
	)
}

func dataFilesCloud_providerJson() (*asset, error) {
	bytes, err := dataFilesCloud_providerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/cloud_provider.json", size: 10043, mode: os.FileMode(420), modTime: time.Unix(1464258041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFilesClusterLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\x51\x6b\x83\x30\x10\xc7\xdf\xfd\x14\x92\xe7\x29\xb1\x65\x7b\xf0\xcd\x81\x14\x19\xed\x8a\x74\x0f\xdd\x18\x21\xd5\xa3\x0d\xad\x51\x92\xcb\x43\x29\x7e\xf7\x5d\xc4\xd6\x95\x16\x8e\x43\xee\xff\xbb\x1f\x67\x2e\x41\x18\xb2\xa3\xdb\x81\x90\x7b\xd0\xc8\xd2\x1f\x1a\x84\xe1\x65\xe8\x14\xd9\xa3\x63\x29\xfb\xf8\x7a\xcf\xa3\x6c\x91\xaf\x36\x51\x56\x7c\xb2\x97\x6b\x5a\x2b\xdb\x9d\xe4\x59\x68\xd9\x80\xc7\xc8\x63\x34\x20\xd8\x09\xe9\x8c\xaa\x94\xde\x8b\xa6\xad\xe1\x44\xcc\x3a\xdb\x2e\xa6\xd4\xba\x9d\xad\x8c\xea\x50\xb5\x5a\xe0\xb9\xf3\x96\x75\x59\x2c\xb3\x72\xfb\x44\x01\x48\x5f\x9e\xc8\x4b\xb1\x2c\x56\x13\xe1\xb4\x42\xe1\x31\x10\xce\xd6\x44\xf0\x38\x79\xbc\xf2\x1f\x70\xfd\x3f\x6f\x07\x23\x0e\xad\x33\xb4\x35\xbf\xed\x8c\xf3\xa6\xd5\x78\xa0\x60\x16\xbf\xb2\x31\xe9\x6f\x5a\x3a\x47\xd6\x12\xe5\x9d\xad\xea\xe8\xbd\x92\x07\x96\x38\x10\x16\xa5\x41\xf0\xe7\xcd\x78\xf2\x16\xf1\x84\x6a\xc3\x79\x3a\x54\xcc\x39\xff\x66\xf7\x0b\xa0\xeb\x11\x9f\xf3\xa7\xf8\x40\xf7\xd4\x7f\x83\x3e\xf8\x0b\x00\x00\xff\xff\xa3\xcf\x05\x06\xcb\x01\x00\x00")

func dataFilesClusterLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesClusterLatestJson,
		"data/files/cluster.latest.json",
	)
}

func dataFilesClusterLatestJson() (*asset, error) {
	bytes, err := dataFilesClusterLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/cluster.latest.json", size: 459, mode: os.FileMode(420), modTime: time.Unix(1464258041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFilesConfigsYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x4f\x6f\xe2\x30\x10\xc5\xef\xf9\x14\xa3\x70\xde\x44\xbb\x47\x6e\x88\x45\xbb\x48\xb4\xaa\x94\xb6\x77\x63\x0f\x78\x84\x33\x8e\xec\x31\x15\xdf\xbe\x0e\x88\x7f\xa5\x4a\x7a\xca\x4b\xde\xef\x39\x93\x97\x99\xc0\xdc\xf3\x86\xb6\x29\x28\x21\xcf\x11\x36\x3e\xc0\xec\x65\xb9\x60\xd3\x79\x62\x89\x55\x7f\xd7\x60\xd8\x63\x80\x0f\x72\xae\x98\x80\xd2\x02\x9e\x81\xd8\xd0\x9e\x4c\x52\x0e\x5a\x14\xeb\x4d\x84\xd4\xe5\xe7\x62\x11\xf4\xdd\xa1\x39\xd3\x05\x9f\x59\x34\x55\x51\xc8\xa1\xc3\x29\x94\x4b\x16\x0c\x1a\x3b\xf1\xe1\x34\x42\x59\xa8\x94\xb3\x2c\xa4\x8f\xb1\x1b\x60\x5a\x00\xf8\x60\x30\x4c\xe1\x77\x96\x89\x6f\x48\x34\xe7\x59\x7b\x0a\xe0\x17\x94\xb5\x45\xe5\xc4\x56\xff\x8f\x97\xba\x11\x25\x29\x96\x17\xb7\x0f\x57\xb3\xbb\x77\xd5\x2b\xbf\x25\x1e\x45\x7c\x92\x61\xe6\xd5\xef\xf0\xe6\x98\x56\x91\x23\xde\x3a\x8a\x52\x3d\x9d\xf4\x2a\xeb\xba\x41\x36\x8b\xde\xfc\x01\x9a\xd6\x51\x07\x5a\xe3\x38\xfa\xc6\xf1\x11\x66\xd5\x62\xec\x94\xc6\xea\xf9\xac\xea\xb9\x45\xbd\x1b\x41\x02\xe6\x6a\x87\x99\xaf\xbd\x7e\xc7\xe4\xd2\xae\x40\x8c\xb6\x6a\x50\xa7\x80\x8d\x45\xe7\xea\x7f\x78\xac\x33\x1e\xa2\x60\x3b\x1b\xfe\xa5\x9a\xf2\xb7\x66\x2e\x3c\x8c\x76\x75\xfe\xa2\xc3\xde\xd9\x2b\x47\x66\x60\x87\xfe\x64\xc9\xfe\xfd\x42\x8d\x2c\xce\x67\x00\x00\x00\xff\xff\x1c\x7f\xa9\xe6\x24\x03\x00\x00")

func dataFilesConfigsYamlBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesConfigsYaml,
		"data/files/configs.yaml",
	)
}

func dataFilesConfigsYaml() (*asset, error) {
	bytes, err := dataFilesConfigsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/configs.yaml", size: 804, mode: os.FileMode(420), modTime: time.Unix(1464258041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFilesDbLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\x5d\x0b\xda\x30\x14\x7d\xef\xaf\x90\x3c\xcf\x92\x5a\xdc\x98\x6f\xad\x75\xae\xa0\x4e\xea\x84\xb9\x31\x42\x6c\x33\x0d\xab\x6d\xc9\xc7\x98\x88\xff\x7d\xb7\x1a\xb1\x16\x15\x7c\xd8\xd6\x87\x96\x40\xd3\x7b\xcf\xbd\x37\xcd\xe1\x70\x0e\x56\xa7\x83\x8a\x5c\xaa\x8d\x60\x12\x0d\xca\x4f\x08\x64\x74\xc7\xd0\xe0\x1a\x7f\x73\x0e\xff\x62\x42\xf2\x3c\x03\xdc\xb7\x53\x00\x42\xef\xed\x3e\x3a\xed\xbf\x1b\x4c\xb2\xbe\x66\x0f\xe6\x0d\x61\xf9\x53\x43\xc3\xc0\xef\xce\xc7\x5d\xdf\x5b\x84\x43\xd3\xf3\x5c\xc3\x65\x91\xd2\x3d\x31\x63\x7d\x2a\x79\x5c\xcd\x17\x82\xc7\x3c\xdb\x90\x5d\x9e\xb0\x14\x00\x73\x6f\x35\xae\xe6\xa5\x5e\xcb\x58\xf0\x42\xc1\xe1\x88\xda\x17\x65\x93\x79\x14\x4e\xbd\x68\x75\xb7\x0d\x53\xb0\x2b\x31\xa3\x88\x4c\xc3\x59\x15\xa3\x33\xae\x48\x09\x64\x44\xcb\x04\x30\xd8\x76\xee\x1d\xb5\x02\xb9\xfe\x65\x39\x83\x09\xb2\xcd\xb5\x80\x4a\xb7\x52\x67\x32\xbb\x3c\x53\x5b\x48\xf5\xcd\xa5\x95\xcf\xb1\xda\x9d\x2a\x46\xa4\xa2\x42\x31\x68\xdc\x41\x3d\xec\xbc\xed\x62\x07\xd6\x67\x8c\x07\xa7\x65\x63\x8c\xbf\xa2\x7a\x0d\xcb\x92\x4b\x85\x8b\xef\x56\x58\xb5\x69\x0f\xb8\x09\x96\x91\xe7\x4f\x46\x4f\xd8\x09\xb4\xa0\xeb\x94\xb5\xfc\xfc\x17\x7e\x3e\x7a\x4f\xa8\xb9\x4d\xb6\xac\xfc\x33\x56\x86\x93\x4f\xcb\xe0\x09\x31\xc3\x34\xd7\x49\xcb\xcd\xdf\xe1\xe6\x6c\x3f\x96\x99\x8b\x58\x4a\xa5\xe2\xb1\x64\x54\xc4\xdb\xba\xa3\xdd\x26\x1f\xdb\x9a\x63\xbf\x03\x63\xbb\x9c\x09\xf5\x6c\x17\x2e\xed\x25\x9f\x1b\x2d\x5a\x9f\x6b\xaa\x62\x81\x9b\xd6\xe7\x9a\xa5\x5a\x9e\xfd\x48\xf5\xef\xba\x5c\x4d\xf4\xb1\x4e\xe1\x56\xb0\xed\xbe\xa6\xcc\x70\xf6\x61\xb2\xfc\xd2\xaa\xb3\xa9\xea\x34\xfc\xb4\x0a\x6d\x8c\x42\xad\xa3\xf5\x27\x00\x00\xff\xff\x71\xa9\xc9\xb4\x28\x0e\x00\x00")

func dataFilesDbLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesDbLatestJson,
		"data/files/db.latest.json",
	)
}

func dataFilesDbLatestJson() (*asset, error) {
	bytes, err := dataFilesDbLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/db.latest.json", size: 3624, mode: os.FileMode(420), modTime: time.Unix(1464258041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFilesPackageLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\x4b\x6f\xa3\x30\x10\xc7\xef\xf9\x14\x88\x73\x89\x0c\x55\xf6\xc0\x8d\x55\xa9\x8a\xb4\x69\x11\x65\x0f\xdd\xd5\x0a\x19\x70\x12\x2b\x84\x20\xdb\xac\xb6\xaa\xf2\xdd\x77\xdc\xbc\x08\x79\x29\x49\x0f\x09\x58\x42\xa8\xf5\x3c\x3c\x98\xff\xaf\x4c\xe7\xa3\xa3\x69\x7a\x81\x93\x31\x1e\x12\xdd\xfe\x0d\xbf\x69\xda\xc7\xe7\x1d\xd6\xf9\xb8\xd4\x6d\xfd\xc9\xf0\x9e\x1f\x3c\xd7\x08\x03\xcf\xf9\xa1\xdf\x2d\x8d\x29\xe5\x45\x86\xdf\xa3\x1c\x4f\x20\x52\xf7\xf2\x94\x92\xb5\xb5\x60\x34\xa1\xf9\x30\x9a\x4c\x53\x92\x81\xf9\x31\x70\xdd\xb5\x95\x97\x31\x4f\x18\x2d\x04\x9d\xe6\x91\x78\x2f\x64\x82\x5a\xfa\x55\x02\x22\xe0\x27\xb0\xfb\x6e\x10\x3d\x38\x6f\x6b\x8f\x32\xa7\x22\x92\x6e\x24\x2a\x79\x0a\x1e\xa8\x6b\x6e\x97\x57\x71\x58\x3e\x97\xcc\x4e\x18\x94\x96\x8b\x51\x34\xb1\x26\x10\x7a\xbf\x0a\xdc\x30\xe2\x42\x1a\xad\x6e\x4f\x5f\x58\x67\xab\xfc\x50\x17\x4e\xb1\xc0\x1b\x69\x4b\x4e\x98\x6e\x9b\xa8\x9a\x6c\x84\x63\x28\x01\x8b\x29\xab\xba\xce\x2b\x1c\x47\xc3\x58\xb7\x2d\xb4\x5a\x9e\x55\x22\x31\x13\x74\x80\x13\xb1\x2f\xcc\x44\xbb\xe3\x12\x5a\x8b\x88\x4b\x9a\xa5\x11\xbc\xe0\x1c\x72\x19\xe6\xee\xa0\xac\xe4\x82\xd4\x4b\x1c\x97\x31\x39\x12\x28\xcf\x20\xc6\x9c\xd4\x22\x8b\x29\x17\x43\x46\xb8\x6e\xf7\xee\xaa\xeb\x24\xc3\x5c\xd0\x84\x13\xcc\x92\x51\xdd\x48\xf3\x41\x56\xfe\x83\xd5\xf5\x4e\x5b\xe7\x0e\xfb\x91\x88\x0b\x38\x1c\x22\xdf\xb9\x85\xcc\x6f\x06\x32\xe1\x0a\x11\xb2\x3f\xaf\x2e\x42\xe8\x97\xbe\x19\x40\xf2\x74\xe1\x7e\x8f\x76\xba\x77\x2a\xbb\x6c\xe9\xff\x35\x74\x82\xf0\xa7\x7f\x84\x80\x57\x59\x54\x59\x28\x06\xa4\x8c\xac\xde\x69\x0c\xec\x13\xf3\x31\x08\x7a\xb7\x01\x41\xf5\x4f\xc2\x36\x05\x35\xeb\x12\x03\x13\x5d\x19\x07\x7e\xf0\x72\x84\x01\x9f\x4d\x07\x84\x73\x90\x34\xce\x14\x08\xf3\x8f\xc1\xa9\x5f\x83\x5e\xa3\x49\xb0\x0e\x92\x60\xed\x26\xc1\xba\x36\x12\xdc\xe7\xd0\xf8\xee\x86\xce\x5e\x10\xdc\x1c\x8e\x0f\xe4\xc6\x2f\xea\x89\x36\xb7\x68\x13\x05\x7b\xe4\x75\x0c\x82\x7d\x72\x56\x0c\x7c\x35\x03\xfe\x61\xfd\xfb\x8c\xfe\x85\x2d\x34\x2f\x87\xaa\xb2\x0c\x4b\x55\x2b\x12\x14\x09\xf3\xa7\x6e\x14\x09\x4f\xc6\xa6\x7c\x6b\x24\x3c\x32\x72\xe0\x1b\xd0\xb7\xfa\x87\x85\xef\x07\x5e\xdf\x09\xde\x1a\xa1\x7d\xeb\xc4\x4e\xe8\x2c\xe1\x9b\x37\xd1\x06\x1d\xd4\xfd\x6e\xd9\x5f\x9b\xea\x41\xba\xf3\xc9\xd0\xb9\x33\xa1\x56\x69\x5f\x0d\x85\x1a\x36\x14\x92\xf2\x5f\x0c\x86\xce\x1f\x09\xb5\x0a\x01\x35\x13\xd2\x9a\x37\x13\x92\x18\xf8\xc1\xcb\x85\x13\xa1\x56\x71\xa0\x46\x42\x5a\x03\xff\x09\x70\xfc\x0b\xdb\x21\x48\xd0\x26\x06\x54\x3b\xd4\xa8\x76\x48\xca\xff\xe2\x76\xa8\x55\x08\xa8\x76\xa8\x81\xed\x90\xc4\xe0\xf2\x76\xa8\x55\x1c\xa8\x76\x48\xbb\xe5\x76\x08\xee\x7f\x3a\xb3\xff\x01\x00\x00\xff\xff\xc5\x64\x70\x93\x48\x24\x00\x00")

func dataFilesPackageLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesPackageLatestJson,
		"data/files/package.latest.json",
	)
}

func dataFilesPackageLatestJson() (*asset, error) {
	bytes, err := dataFilesPackageLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/package.latest.json", size: 9288, mode: os.FileMode(420), modTime: time.Unix(1464258041, 0)}
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
	"data/files/ci.latest.json": dataFilesCiLatestJson,
	"data/files/cloud_provider.json": dataFilesCloud_providerJson,
	"data/files/cluster.latest.json": dataFilesClusterLatestJson,
	"data/files/configs.yaml": dataFilesConfigsYaml,
	"data/files/db.latest.json": dataFilesDbLatestJson,
	"data/files/package.latest.json": dataFilesPackageLatestJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"files": &bintree{nil, map[string]*bintree{
			"ci.latest.json": &bintree{dataFilesCiLatestJson, map[string]*bintree{}},
			"cloud_provider.json": &bintree{dataFilesCloud_providerJson, map[string]*bintree{}},
			"cluster.latest.json": &bintree{dataFilesClusterLatestJson, map[string]*bintree{}},
			"configs.yaml": &bintree{dataFilesConfigsYaml, map[string]*bintree{}},
			"db.latest.json": &bintree{dataFilesDbLatestJson, map[string]*bintree{}},
			"package.latest.json": &bintree{dataFilesPackageLatestJson, map[string]*bintree{}},
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

