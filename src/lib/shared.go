package lib

import "os"

var CACHE_DIR = "cache/heads"

func init() {
	if _, err := os.Stat(CACHE_DIR); os.IsNotExist(err) {
		err := os.MkdirAll(CACHE_DIR, 0755)
		if err != nil {
			panic("Failed to create cache directory: " + err.Error())
		}
	}
}
