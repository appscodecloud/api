// Code generated by go-bindata.
// sources:
// data/files/ci.latest.json
// data/files/cloud_provider.json
// data/files/cluster.latest.json
// data/files/configs.yaml
// data/files/db.latest.json
// data/files/pkg.latest.json
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

var _dataFilesCiLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x96\x51\x6b\xdb\x30\x10\xc7\xdf\xf3\x29\x8c\x9e\xd7\x20\x39\x8e\x31\x79\xcb\x9a\x52\x02\x5d\x3b\x42\xfa\xb0\x8d\x21\x14\xfb\x96\x8a\x5a\x52\x2a\x4b\x23\xa5\xf4\xbb\xef\x6c\xd2\xd4\xae\x5a\x30\x14\xb2\x3c\x04\x8c\x1f\xee\xff\xbf\x3b\x9d\xef\x87\xd1\xd3\x20\x8a\xc8\xca\xcb\xb2\xe0\x62\x0d\xda\x91\xc9\x2f\x8c\x44\xd1\x53\xf3\x46\xad\xba\xf7\x64\x42\xce\xe7\x67\xe7\x6c\xc1\xce\x66\x37\xe4\xcb\x8b\xe2\x1e\x37\xd0\x48\xaf\xa1\xca\xaf\xf8\x2e\xfc\xf5\x76\x7e\x35\xe3\xd3\xcb\x8b\xeb\xe5\xab\xfe\xe0\x8d\x13\xdc\xc2\x83\x97\x16\x0a\x74\xe5\x72\xd8\x6e\xbe\x37\x16\xb2\xda\x94\xe2\x91\x6b\xa1\x9a\x62\xb5\x27\x9a\x76\x3d\x1b\x2b\x73\xa9\xd7\x5c\x99\x02\x4a\x34\x7d\x9f\xfe\xb8\x6c\x9d\x4e\x2a\xe0\x5e\x4b\x9c\x88\x7c\x9b\x5f\xdf\x2e\x2f\xc2\xcc\x3f\xc6\x2a\x5f\x0a\x74\xd0\x21\x0b\x9b\xd7\x36\xac\x51\xe1\x41\x5f\x3e\x47\x9d\x0c\x96\xdf\x19\x6f\x31\x6b\xb4\xcf\xd9\xc5\x95\xd1\xee\x0e\x85\xf1\x98\xec\x84\xe7\x7d\x55\x05\x4e\x14\xc2\x89\x6e\x31\x6b\xfe\xca\x02\xea\x62\x33\xb9\x96\x4e\x94\x37\x39\x08\xdd\xae\x6b\x2a\x14\xaf\xa4\xf6\xdb\x76\x34\xdf\xe0\x5a\x58\x2b\x60\x85\xea\x06\x70\x8a\x7b\x32\x19\xd1\xb6\x07\x72\xa3\x14\xe8\x02\x0a\x0e\x5b\xc8\xbd\x33\xb6\xea\x66\xc1\xd6\x81\xd5\xa2\xe4\xb2\x5e\x0f\x5b\xaf\xda\x4d\xf7\xeb\x0b\x37\x17\x8c\x8b\xa3\x02\xaf\x9c\xb0\xae\x59\x74\x4c\x59\x7a\x46\x19\x3e\x4b\x4a\x27\xcd\x33\xa4\x94\xfe\x24\xdd\x84\xe6\x6c\x8d\x7d\x44\xdf\xb5\x0f\x5a\x5d\xde\x61\x34\x5e\xc4\x27\x46\xfb\x31\xca\x28\x3d\x08\xa4\xf1\x5b\x48\xe3\x00\xd2\xa4\x07\xa4\xf1\x87\x90\xc6\x08\xe9\xd1\xe0\x97\x9c\xf0\xeb\x89\x5f\xfa\x9f\xf0\x4b\x02\xfc\xd2\x4f\xe1\x97\x1c\x0f\x7e\xc9\x22\x3b\xe1\xd7\x0f\xbf\x78\x7c\x18\xfc\x92\xb7\xf8\x65\x01\x7e\xd9\xa7\xf0\xcb\x8e\x07\xbf\x6c\x81\x6d\x4e\xfc\xf5\xe2\x6f\x74\xa0\xdf\x5f\x16\x5c\x11\xd3\x00\x40\xd6\xe7\x07\x98\x7c\x7c\x49\x4c\x0f\x8e\x20\xbe\x7f\x0f\x9e\x07\xff\x02\x00\x00\xff\xff\xf8\x0d\x7e\x9d\xbd\x0c\x00\x00")

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

	info := bindataFileInfo{name: "data/files/ci.latest.json", size: 3261, mode: os.FileMode(420), modTime: time.Unix(1464354088, 0)}
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

var _dataFilesClusterLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\xcf\x6a\xf2\x40\x14\xc5\xf7\x3e\xc5\x30\xcb\xef\xab\x61\xa2\xd8\x45\x76\x69\x09\x22\x6d\x6d\x69\xe3\xa2\x2d\x65\x18\x93\x5b\x1d\xcc\x4c\xe2\xfc\x59\x88\xe4\xdd\x7b\x33\x44\x63\xb1\x30\x5c\xc2\x3d\xbf\x73\x38\xb9\xc7\x11\x21\x74\xe7\xd7\xc0\xc5\x06\xb4\xa3\xc9\x27\x2e\x08\x39\x86\x89\x92\xdd\x79\x9a\xd0\x87\xd5\x5d\x36\x4e\xe7\xd9\x32\x1f\xa7\x8b\x67\x7a\x73\x52\xdd\xa1\x01\x94\xef\x1f\xdf\xf2\xd7\x61\x6b\xfd\x9a\xf7\x4a\x67\xe4\xc1\x38\xc8\x7b\x5f\x3b\xc1\x0d\xec\xbd\x34\x50\x22\x54\x54\xde\x3a\x30\xd1\x45\x8d\x33\x5c\x4a\xdb\x54\xe2\xc0\xb5\x50\x21\x0f\x11\xa3\xc1\x81\x1d\x90\xc6\xc8\x42\xea\x0d\x57\x75\x09\x15\x32\x2f\xe9\xfb\xfc\xa2\xa2\x54\xc0\xbd\x96\xf8\x6b\xf4\x69\xb1\x5c\xe5\xd9\xb5\xf3\xbb\x36\xca\x57\x02\x09\x16\xcd\xc8\x3f\x52\x34\x9e\xfc\x27\x2c\x8a\xf1\xdb\x08\x75\xdd\xa6\x33\x62\xaa\xc5\xf6\xa7\x4b\x75\x71\x60\xf8\xb6\xf6\x06\x73\xa6\x67\x4f\xbf\x57\xb5\x76\x5b\x14\x26\xd1\x8c\xf6\x4a\x3b\xc4\x0a\x07\xdc\x3a\x61\x5c\xb8\xc7\x84\xc5\xb7\x63\x16\xe3\xcb\x19\x4b\xc2\x8b\x18\x63\x1f\xf4\xb7\x01\x74\xd9\xe3\x53\xf6\x27\x1e\xe8\x16\xe7\xd7\xa8\x1d\xfd\x04\x00\x00\xff\xff\x94\x83\x36\x08\xe9\x01\x00\x00")

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

	info := bindataFileInfo{name: "data/files/cluster.latest.json", size: 489, mode: os.FileMode(420), modTime: time.Unix(1464354088, 0)}
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

var _dataFilesDbLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\xdb\x6e\xda\x40\x10\x7d\xf7\x57\x44\x7e\x2e\xd6\x12\x94\x56\xe5\xcd\x18\x87\x20\xd1\x04\x61\x90\x7a\x51\x65\x2d\x78\x08\xab\xfa\x96\xbd\x54\x8d\xa2\xfc\x7b\x87\x65\x51\x8c\x65\x5a\x9b\x24\xe2\xc5\xc8\x12\x30\x67\x66\x76\x3c\xc7\x3e\xe7\xc9\xba\xb8\xb0\xf3\x4c\xc8\x7b\x0e\xc2\xee\x6f\xff\x62\x20\xa5\x09\xd8\xfd\x97\xf8\x87\x5d\xf8\x37\x70\xc1\xb2\x14\xf3\x7e\xe8\x00\x86\x3e\x3b\x57\xb6\xfe\xfd\xd3\xe4\x44\xcb\x17\xf4\xc9\x7c\x63\x58\xfc\x52\xd8\x70\x38\xe8\x4c\x47\x9d\x81\x1b\x8c\x3d\xd3\x53\x83\xf2\x31\x07\x8d\x16\x83\x42\x2d\x43\x03\x4c\xef\x82\xf9\x68\xe6\x07\x45\xf8\x41\x65\x92\x86\x1c\x1e\x14\xe3\x10\x61\x52\xb4\x74\xd6\x1c\xc0\x29\x0d\xbd\x1b\x8a\x89\x3c\xa6\x8f\xa1\xb9\xaf\x01\x15\x6c\x55\xc4\x73\xce\x56\x2c\xbd\x0f\x93\x2c\x82\x18\x13\xae\x67\xbe\x7f\x30\x20\x4b\x20\x54\x29\x93\x88\x7d\x19\xdf\x2e\xe6\x7e\x55\xf5\x3a\xe3\x89\x8a\x29\xe6\x10\xa7\x5b\x75\xfc\x36\x11\xfb\x88\x68\xbf\xe8\x7d\x03\xe0\xe1\x26\x53\x1c\x2b\x7b\x85\x3a\x83\x24\x59\x2a\x37\x08\x5d\x99\x4d\x6f\x3f\xcf\xc5\xee\x54\x42\x28\x24\xe5\x52\xef\xe1\x92\x74\x3f\x76\x48\x17\xaf\x39\x21\x7d\x7d\x39\x84\x90\xef\x76\xb9\x04\xd2\xc8\x14\xf4\x48\x65\x81\x55\x3a\xeb\x08\x9d\xc3\xc5\xcc\x1d\x4c\xfc\xf7\x20\x34\xa7\x2c\xaa\x43\xe8\x50\x71\xba\x8c\xe1\x5f\x94\x4e\xdd\x6f\xa3\x96\xd2\x9a\x94\xde\xb8\xe7\x64\xf3\xf0\xf4\x96\xc8\x57\x10\xe9\x4d\xee\x16\xc3\x73\x72\xe9\xc5\x99\x8a\x5a\x3a\xeb\xd2\xb9\xb3\x52\xcb\x9c\x6a\x43\x4c\x85\x64\x2b\x01\x94\xaf\x36\x65\x77\x3e\x04\x8f\x5b\x74\xd7\xf9\x84\x26\xbd\x1f\xc9\xbe\x74\x7a\xb8\xb2\x46\x9e\xed\x07\xcd\x3d\xdb\x9f\xb8\xc1\x7c\xec\x05\xbe\x3b\xf3\x6e\x6a\x19\x77\xd5\xfd\x54\x3d\x52\xad\x7b\xbf\x99\x42\x20\xb1\x27\xb8\x77\x23\x6a\xb5\x50\xd4\xa5\xb6\xf5\xf1\xd7\xe9\x05\x4b\xd7\xb1\xfa\x53\x16\x0a\x13\x3d\xae\x10\xb8\x13\xe2\xf4\x9a\x69\xc2\xf8\xf6\x7a\xb2\xf8\xda\x5c\x17\x76\x75\xb5\x04\xe1\x60\xee\xaa\xc7\xa5\x55\x82\x37\x53\x02\x43\xe7\x09\x6a\x50\x8f\x50\x2d\x03\xff\x25\xb4\x7d\xff\x4f\x7e\xff\xad\x67\xeb\x6f\x00\x00\x00\xff\xff\x4f\xcc\xb8\xbc\xcc\x0f\x00\x00")

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

	info := bindataFileInfo{name: "data/files/db.latest.json", size: 4044, mode: os.FileMode(420), modTime: time.Unix(1464354088, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFilesPkgLatestJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\xd8\x5d\x8b\x9b\x4c\x14\x07\xf0\xfb\x7c\x8a\xe0\xf5\x13\x19\xb3\xec\x73\xe1\x9d\x65\xdd\x56\xda\xec\x8a\x71\x2f\xb6\xa5\xc8\xa8\x93\x64\xc8\x38\xba\xf3\x52\xba\x2c\xfb\xdd\x7b\x4c\x93\x26\x6e\x42\xa1\xac\x09\x5e\x1c\x08\x12\xfc\x9f\x19\x8f\xc3\x0f\x19\xe6\x65\x34\x1e\x3b\xcd\x7a\xe9\xf8\xdf\xe0\xdf\x78\xfc\xb2\xb9\xc2\x3d\xbd\xb6\x8e\xef\x7c\x9a\x44\x77\x37\x51\x38\x49\x93\x28\xf8\xe2\xfc\xb7\x0b\xcd\x73\xc3\x20\x8d\x3f\x7f\xdc\xdf\xd3\x36\xcf\xb6\xf7\x1f\xe6\x61\xb2\x0f\x9e\x6c\x6d\x68\xa6\xd8\x93\xe5\x8a\x95\x10\xc3\xe3\x5c\xab\x99\xda\x97\x94\x5c\x37\x82\x3e\x67\x92\x56\xed\xf8\x48\x96\x9c\xed\xd3\x46\xf1\x82\xcb\x65\x56\xd5\x25\x13\x10\xdf\x26\x61\xd8\x79\xae\x2e\x14\x6f\x0c\xaf\xe5\xae\x81\xb7\xed\xf2\x8a\x65\x56\x72\x03\xd1\x4d\xf0\x78\x3c\xf3\xa2\x56\x95\x15\x14\x62\xe2\x7a\xc7\x6d\xb5\x65\x30\x81\x86\xe6\x77\xeb\xd3\x0e\x66\x0a\x5a\x92\x66\x95\x55\xd3\x0a\x86\x5e\xfd\x19\xd8\x09\x69\xd3\x86\x53\xf7\xda\xd9\xa6\xaf\xdd\x95\xe9\xce\xb9\x5b\x1a\xdf\x23\x87\xb3\xad\x68\x0e\x3d\x50\x53\x2b\x17\x9a\x5a\x67\xcb\xdc\xf1\xa7\x87\x15\x54\x19\xbe\xa0\x85\xd9\xc7\x1e\x39\xcc\x0b\xee\xe6\x96\x8b\x32\xa3\x4b\x26\x61\x1d\x26\xde\x61\x28\xac\x36\x4c\xb9\x6b\x9b\xb3\x93\x05\x65\xee\x36\x94\x97\x6e\x53\x6b\xb3\x54\x4c\x3b\xfe\xf5\x89\x94\x09\xaa\x0d\x2f\x34\xa3\xaa\x58\x9d\x2e\xe1\x72\x21\xec\x4f\xc8\x8e\xd6\xa2\xa4\x86\x65\xda\xc0\x7b\x6c\x8c\x4c\x89\xf7\xff\x84\x78\xf0\x4b\x09\xf1\x37\x3f\x97\x10\xf2\xd5\xe9\x0e\x60\xb2\xdc\x96\x5f\x91\x93\xe5\xa3\x83\xa7\x1c\xd9\x9e\xa7\x41\x92\x3e\xc4\x17\xd6\x3d\x6f\x5f\xd2\x36\xe8\x1b\x7d\x9f\xd9\x77\x9c\xdc\x5f\xd8\x76\xac\xea\x05\xd3\x1a\xa8\x52\x81\xc0\x11\xf8\x99\x81\x87\x77\xe9\xe4\x43\x98\x06\x17\xf3\x1d\x4a\x58\x66\xd0\xa4\xdf\xb5\x3d\x79\xd3\x32\xe2\x46\xdc\x47\xb8\xe3\xcb\xc2\x8e\x15\xff\x01\x2d\x8f\x23\x09\x6f\x29\x04\x6d\xb9\x22\x71\x24\x7e\xde\xef\x77\xd7\xd1\x99\x89\xdf\x2a\xf6\xae\xaf\x76\x9c\x44\xb3\x20\x79\x44\xd5\xa8\xfa\xef\xaa\x67\xd3\xd9\xef\x63\x93\xa1\x1c\x98\x40\x43\x48\x1b\x69\xf7\x44\x7b\x7b\x6a\x32\x9c\xf3\x12\xe4\x8d\xbc\xfb\xe3\x1d\x27\xf7\x03\x3b\x2e\x41\xdf\xe8\xbb\x1f\xdf\x41\x3c\xb0\x9d\x09\x34\x84\xb4\x91\x76\x4f\xb4\x07\xb7\x33\x41\xde\xc8\xbb\x3f\xde\xc3\xdb\x99\xa0\x6f\xf4\xfd\x8f\xbe\xe1\xfa\x7d\xf4\x3a\xfa\x15\x00\x00\xff\xff\x09\xaf\x41\x3a\x78\x22\x00\x00")

func dataFilesPkgLatestJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFilesPkgLatestJson,
		"data/files/pkg.latest.json",
	)
}

func dataFilesPkgLatestJson() (*asset, error) {
	bytes, err := dataFilesPkgLatestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/files/pkg.latest.json", size: 8824, mode: os.FileMode(436), modTime: time.Unix(1464908933, 0)}
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
	"data/files/pkg.latest.json": dataFilesPkgLatestJson,
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
			"pkg.latest.json": &bintree{dataFilesPkgLatestJson, map[string]*bintree{}},
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

