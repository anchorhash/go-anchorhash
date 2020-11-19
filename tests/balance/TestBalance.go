package main

import (
	"fmt"
	anchorhash "go-anchorhash/anchorhash"
	helpers "go-anchorhash/helpers"
)

//balanceTesting checks the balance of the anchor hash implemantaion
func balanceTesting(maxSize uint32, wSize uint32, seed uint32, numOfKeys int) error {
	keyLen := 5
	buckets := helpers.CreateStringSlice("b", 0, int(wSize), "")
	keys := helpers.CreateKeys(numOfKeys, int64(seed), keyLen)
	hist := make(map[string]int)
	anchor, err := anchorhash.NewHashWrapper(maxSize, buckets, seed)
	if err != nil {
		return err
	}
	for _, key := range keys {
		hist[anchor.GetResource(key)]++
	}
	fmt.Printf(
		"balancing test with:\n"+
			"buckets: %v , keys: %v , seed: %v\n"+
			"keys/buckets ratio is %v:\n",
		wSize, numOfKeys, seed, ((float64(numOfKeys)) / float64(wSize)))
	fmt.Println(hist)
	return nil
}
