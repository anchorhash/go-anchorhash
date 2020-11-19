package main

import (
	"fmt"
	anchor "go-anchorhash/anchorhash"
	helpers "go-anchorhash/helpers"
	"math"
	"math/rand"
	"testing"
)

//BenchmarkRate check the anchor performance rate with diffrent ratios and fixed Load (=|A|)
func BenchmarkRatioRate(b *testing.B) {
	//hyper parameters
	seed := 1
	keyLen := 5
	numOfKeys := 100000
	numOfBuckets := int(math.Pow(2, 10))
	keys := helpers.CreateKeys(numOfKeys, int64(2), keyLen)
	//benchmarks
	merges := []struct {
		name string
		fun  func(ratio int, numOfBuckets int, seed int, b *testing.B) *anchor.HashWrapper
	}{
		{"testAnchor", CreateAnchorForRateTest},
		//one can add here diffrent tests
	}
	for _, merge := range merges {
		fmt.Println("start test ")
		//run the test ten times with ten incrising hyper parameters
		for k := 0; k <= 5; k++ {
			//setup for the test
			ratio := int(math.Pow(2, float64(k)))
			b.Run(fmt.Sprintf("%s with ratio %d", merge.name, ratio), func(b *testing.B) {
				anchor := merge.fun(ratio, numOfBuckets, seed, b)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for _, key := range keys {
						anchor.GetResource(key)
					}
				}
				b.StopTimer()
				b.ReportMetric(float64(numOfKeys), "num_of_keys")
				b.ReportMetric(float64(numOfBuckets), "num_of_total_buckets")
			})
		}
	}

}

//BenchmarkRate check the anchor performance rate with diffrent loads (=|A|) and fixed ratio
func BenchmarkLoadRate(b *testing.B) {
	//hyper parameters
	seed := 1
	ratio := 2
	keyLen := 5
	numOfKeys := 100000
	keys := helpers.CreateKeys(numOfKeys, int64(2), keyLen)
	//benchmarks
	merges := []struct {
		name string
		fun  func(ratio int, numOfBuckets int, seed int, b *testing.B) *anchor.HashWrapper
	}{
		{"testAnchor", CreateAnchorForRateTest},
		//one can add here diffrent tests
	}
	for _, merge := range merges {
		fmt.Println("start test ")
		//run the test ten times with ten incrising hyper parameters
		for k := 2; k <= 10; k++ {
			//setup for the test
			maxSize := int(math.Pow(2, float64(k)))
			b.Run(fmt.Sprintf("%s with |A|=%d,ratio=%d", merge.name, maxSize, ratio), func(b *testing.B) {
				anchor := merge.fun(ratio, maxSize, seed, b)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for _, key := range keys {
						anchor.GetResource(key)
					}
				}
				b.StopTimer()
				b.ReportMetric(float64(numOfKeys), "num_of_keys")
				b.ReportMetric(float64(maxSize), "num_of_total_buckets")
			})
		}
	}

}
func CreateAnchorForRateTest(ratio int, maxSize int, seed int, b *testing.B) *anchor.HashWrapper {
	//test parameters
	buckets := helpers.CreateStringSlice("b", 0, maxSize, "")
	anchor, err := anchor.NewHashWrapper(uint32(len(buckets)), buckets, (uint32(seed))) //create anchor with full working set //TODO:uncomment that
	if err != nil {
		b.Fatal("bad parameters:error creating anchorhash")
	}
	//removing buckets:
	permutation := rand.Perm(len(buckets))
	bucketsToRemove := len(buckets) - int(float64(len(buckets))/float64(ratio))
	if bucketsToRemove < 0 || bucketsToRemove > len(buckets)-1 {
		b.Fatal("bad parameters:ratio error")
	}
	for _, value := range permutation[:bucketsToRemove] {
		anchor.RemoveResource(buckets[value])
	}
	return anchor
}
