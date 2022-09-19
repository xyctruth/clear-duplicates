package main

import (
	"github.com/m-manu/go-find-duplicates/entity"
	"github.com/m-manu/go-find-duplicates/fmte"
	"os"
	"sort"
)

func clearDuplicates(duplicates *entity.DigestToFiles) error {
	var err error
	for iter := duplicates.Iterator(); iter.HasNext(); {
		_, paths := iter.Next()
		sort.Strings(paths)
		removePaths := paths[1:]
		for _, path := range removePaths {
			err = os.Remove(path)
			if err != nil {
				return err
			}
			fmte.Printf("remove duplicate file: %s successful\n", path)
		}
	}
	return err
}
