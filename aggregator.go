package goassets

import (
	"os"
	"fmt"
	"regexp"
	"path/filepath"
)

type IAggregator interface {
	Aggregate(nameOrPath string) (string, error)
}

type AssetAggregator struct {
	assetFormat string
	assetNameRegex string
}

func NewAssetAggregator(assetNameRegex string, assetFormat string) IAggregator {
	assetAggregator := &AssetAggregator{}
	assetAggregator.Init(assetNameRegex, assetFormat)
	return assetAggregator
}

func (this *AssetAggregator) Init(assetNameRegex string, assetFormat string) {
	this.assetNameRegex = assetNameRegex
	this.assetFormat = assetFormat
}

func (this *AssetAggregator) Aggregate(nameOrPath string) (string, error) {
	fi, err := os.Stat(nameOrPath)
	
	if err != nil { return "", err }

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return this.processDir(nameOrPath)
	case mode.IsRegular():
		return this.processFile(nameOrPath, fi)
	}

	return "", err
} 

// private methods
func (this *AssetAggregator) processDir(nameOrPath string) (string, error) {
	scripts := ""

	err := filepath.Walk(nameOrPath, func (path string, info os.FileInfo, err error) error {
		if fileMode := info.Mode(); err == nil && (fileMode.IsRegular() && this.isValidAsset(path)) {
			scripts += fmt.Sprintf(this.assetFormat, path, fmt.Sprint(info.ModTime().Unix()))
		}
		return err
	})
	
	return scripts, err
}

func (this *AssetAggregator) processFile(path string, fileInfo os.FileInfo) (string, error) {
	scripts := ""
	
	if this.isValidAsset(path) {
		scripts = fmt.Sprintf(this.assetFormat, path, fmt.Sprint(fileInfo.ModTime().Unix()))
	}
	
	return scripts, nil
}

func (this *AssetAggregator) isValidAsset(path string) bool {
	valid, err := regexp.MatchString(this.assetNameRegex, path)
	if err != nil { panic(err) }
	return valid
}


