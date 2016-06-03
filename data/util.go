package data

import (
	"encoding/json"
	"fmt"
	"github.com/appscode/api/data/files"
	"strings"
)

func DBGenericName(dbName, sku string) (string, error) {
	bytes, err := files.Asset("data/files/db.latest.json")
	if err != nil {
		return "", err
	}

	var dbs DBProduct
	err = json.Unmarshal(bytes, &dbs)
	if err != nil {
		return "", err
	}
	if pg, ok := dbs[dbName]; ok {
		for _, dt := range pg.DB {
			if strings.ToUpper(dt.Sku) == strings.ToUpper(sku) {
				return dt.DisplayName, nil
			}
		}
	}
	return "", fmt.Errorf("Unknown SKU provided")
}


func DBGenericNameMap() (map[string]string, error) {
	bytes, err := files.Asset("data/files/db.latest.json")
	if err != nil {
		return nil, err
	}

	var dbs DBProduct
	err = json.Unmarshal(bytes, &dbs)
	if err != nil {
		return nil, err
	}

	dbType := []string{"postgres", "elasticsearch", "influx"}
	skuMap := make(map[string]string)

	for _, dType := range dbType {
		if db, ok := dbs[dType]; ok {
			for _, dt := range db.DB {
				skuMap[dt.Sku] = dt.DisplayName
			}
		}
	}

	return skuMap, nil
}

func DBSku(dbName, mode string) (string, error) {
	bytes, err := files.Asset("data/files/db.latest.json")
	if err != nil {
		return "", err
	}

	var dbs DBProduct
	err = json.Unmarshal(bytes, &dbs)
	if err != nil {
		return "", err
	}
	if pg, ok := dbs[dbName]; ok {
		for _, dt := range pg.DB {
			if strings.ToLower(dt.DisplayName) == strings.ToLower(mode) {
				return dt.Sku, nil
			}
		}
	}
	return "", nil
}

func DBVersion(dbName, dbVersion string) (string, error) {
	bytes, err := files.Asset("data/files/db.latest.json")
	if err != nil {
		return "", err
	}

	var dbs DBProduct
	err = json.Unmarshal(bytes, &dbs)
	if err != nil {
		return "", err
	}

	version := ""
	if pg, ok := dbs[dbName]; ok {
		for _, v := range pg.Versions {
			if v == dbVersion {
				return dbVersion, nil
			}
			version = v
		}
	}
	if dbVersion == "" {
		return version, nil
	}
	return "", nil
}

/*
curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer <token>" "https://api.digitalocean.com/v2/sizes"
*/
func CIBuildAgent(sku string) (*BuildAgent, error) {
	bytes, err := files.Asset("data/files/ci.latest.json")
	if err != nil {
		return nil, err
	}

	var ci CIProduct
	err = json.Unmarshal(bytes, &ci)
	if err != nil {
		return nil, err
	}
	for _, agent := range ci.BuildAgent {
		if agent.Sku == strings.ToUpper(sku) {
			return agent, nil
		}
	}
	return nil, fmt.Errorf("Can't find build agent for CIProduct sku %v", sku)
}

func ClusterMachineType(cloud, sku string) (*InstanceType, error) {
	bytes, err := files.Asset("data/files/cloud_provider.json")
	if err != nil {
		return nil, err
	}

	var provider ClusterProvider
	err = json.Unmarshal(bytes, &provider)
	if err != nil {
		return nil, err
	}
	for _, instance := range provider.Provider[cloud].InstanceTypes{
		if instance.Sku == sku {
			return instance, nil
		}
	}
	return nil, err
}