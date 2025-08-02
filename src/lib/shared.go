package lib

import "os"

var CACHE_DIR = "cache"
var CACHE_SUBDIRS = []string{
	"cache/heads",
	"cache/leather",
	"cache/potions",
}

func init() {
	if _, err := os.Stat(CACHE_DIR); os.IsNotExist(err) {
		err := os.MkdirAll(CACHE_DIR, 0755)
		if err != nil {
			panic("Failed to create cache directory: " + err.Error())
		}
	}
	for _, dir := range CACHE_SUBDIRS {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				panic("Failed to create subdirectory: " + dir + ": " + err.Error())
			}
		}
	}
}
