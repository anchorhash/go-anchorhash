package main

import (
	"errors"
	"fmt"
	anchorhash "go-anchorhash/anchorhash"
	helpers "go-anchorhash/helpers"
	"math/rand"
	//hash import
)

type keyDiffer struct {
	key                 string
	FirstAnchorHashing  string
	SecondAnchorHashing string
}

func consistencyTestLoop(numOfTests int, maxSize int, maxKeysNum int, seed uint32) {
	if numOfTests <= 0 {
		fmt.Println("numOfTests needs to be positive number")
	}
	s := rand.NewSource(int64(seed))
	r := rand.New(s)
	var A, W, n, numOfKeys int
	var res string
	var err error
	for i := 0; i < numOfTests; i++ {
		A, W, n, numOfKeys, err = CreateTestHyperParameters(i, maxSize, maxKeysNum, r)
		if err != nil {
			fmt.Println("wrong parameters")
		}
		if consistencyTest(A, W, n, seed, numOfKeys, r) {
			res = "pass"
		} else {
			res = "failed"
		}
		fmt.Printf("test num : %v \t consistencyTest with: \tA: %v  \t,W: %v  \t,steps: %v  \t,numOfKeys : %v  \t, result : %v\n", i, A, W, n, numOfKeys, res)
	}
}

//CreateTestHyperParameters create the Test Hyper Parameters
func CreateTestHyperParameters(i int, maxSize int, maxKeysNum int, r *rand.Rand) (A int, W int, n int, numOfKeys int, err error) {
	A, errA := positiveModuloOutput(r.Int(), maxSize)
	W, errW := positiveModuloOutput(r.Int(), A)
	n, errN := positiveModuloOutput(r.Int(), min(W, A-W))
	numOfKeys, errNumOfKeys := positiveModuloOutput(r.Int(), maxKeysNum)
	if errA != nil || errW != nil || errN != nil || errNumOfKeys != nil {
		err = errors.New("wrong parameters")
	}
	return
}

//consistencyTest check the consistency of the anchor hashing
/*
@input: A : num of max size of capacity of buckets in anchors
@input: W : num of initial capacity of buckets in anchors (numInitiallyBuckets: initial working set)
@input: n : num of step during test note that to avoid errors it is required that n+W<=A and W-n>0
@input: seed : seed of random actions of the test
@input: r : pseudorandom src from type hash.Hash64
@output: res : boolean flag : true if the consistencyTest passed and false otherwise
*/
func consistencyTest(A int, W int, n int, seed uint32, numOfKeys int, r *rand.Rand) (res bool) {
	res = true
	if W-n <= 0 || W+n > A {
		fmt.Println("wrong parameters entered!")
		res = false
		return
	}
	keys := helpers.CreateStringSlice("key:", 0, int(numOfKeys), "")
	initialWorkingSet := helpers.CreateStringSlice("initialBucket:", 0, W, "")
	currentWorkingSet := helpers.CreateStringSlice("initialBucket:", 0, W, "") //TODO: not sure if i need to mange two diffrent slices or one . validate that the anchor create it's own copy
	actions := [2]string{"AddResource", "RemoveResource"}                      //TODO:find a way to make this an enum or const
	FirstAnchor, errFirst := anchorhash.NewHashWrapper(uint32(A), initialWorkingSet, seed)
	SecondAnchor, errSecond := anchorhash.NewHashWrapper(uint32(A), initialWorkingSet, seed)
	//error handling
	if errFirst != nil {
		fmt.Printf("error with FirstAnchor init")
		return
	}
	if errSecond != nil {
		fmt.Printf("error with SecondAnchor init")
		return
	}
	var targetBucket string
	if !anchorIsConsistent(FirstAnchor, SecondAnchor, keys, true, "") {
		res = false
		return res
	}
	for i := 0; i < n; i++ {
		curAction := ChooseAction(r, actions)
		//first anchor movement:
		if curAction == actions[0] { // curAction==AddResource
			targetBucket = fmt.Sprintf("newBucketNum:%v", i)
			FirstAnchor.AddResource(targetBucket)
		} else { // curAction==RemoveResource
			targetBucket = ChooseBucketToRemove(r, currentWorkingSet, i)
			FirstAnchor.RemoveResource(targetBucket)
		}
		//fmt.Println("iter", i, "action :", curAction, "on FirstAnchor :", "targetbucket", targetBucket)
		//anchorComparePrint(FirstAnchor, SecondAnchor, keys, targetBucket)
		if !anchorIsConsistent(FirstAnchor, SecondAnchor, keys, false, targetBucket) {
			res = false
			break
		}
		//second anchor movement:
		if curAction == actions[0] { // curAction==AddResource
			SecondAnchor.AddResource(targetBucket)
			currentWorkingSet = append(currentWorkingSet, targetBucket)
		} else { // curAction==RemoveResource
			SecondAnchor.RemoveResource(targetBucket)
			removeBucketFromSlice(currentWorkingSet, targetBucket)
		}
		//fmt.Println("iter", i, "action :", curAction, "on SecondAnchor :", "targetbucket", targetBucket)
		//anchorComparePrint(FirstAnchor, SecondAnchor, keys, targetBucket)
		if !anchorIsConsistent(FirstAnchor, SecondAnchor, keys, true, targetBucket) {
			res = false
			break
		}
	}
	return res
}

//anchorIsConsistent checks if the anchors are consistent
func anchorIsConsistent(FirstAnchor *anchorhash.HashWrapper, SecondAnchor *anchorhash.HashWrapper, keys []string, shouldBeSame bool, targetBucket string) bool {
	IsSame, diffSlice := anchorCompare(FirstAnchor, SecondAnchor, keys)
	if IsSame || ((!shouldBeSame) && anchorDifferIsValid(diffSlice, targetBucket)) {
		return true
	}
	return false
}

//print the result of anchor Comapre func
func anchorComparePrint(FirstAnchor *anchorhash.HashWrapper, SecondAnchor *anchorhash.HashWrapper, keys []string, targetBucket string) {
	IsSame, diffSlice := anchorCompare(FirstAnchor, SecondAnchor, keys)
	if !IsSame {
		fmt.Println("two anchors are diffrent:")
		fmt.Println("anchor differ is Valid?:", anchorDifferIsValid(diffSlice, targetBucket))
	} else {
		fmt.Println("two anchors are the same")
	}
}

//anchorDifferIsValid checks if the change occur on two anchors is a valid change (the only keys that differ are the ones that's belong to the targetBucket)
func anchorDifferIsValid(differences []keyDiffer, targetBucket string) bool {
	res := true
	for _, element := range differences {
		if element.FirstAnchorHashing != targetBucket && element.SecondAnchorHashing != targetBucket {
			res = false
			break
		}
	}
	return res
}

//anchorCompare checks if two anchors are the same.
/*
@input: a, b : the two anchor hashers to comapre with
@input: keys : the keys to comapre the two hashers with
@output: IsSame : true if two anchor are same in terms of hashing the set of keys and false otherwise
@output: differences : a slice of all the keys that the both anchors are differ on
*/
func anchorCompare(FirstAnchor *anchorhash.HashWrapper, SecondAnchor *anchorhash.HashWrapper, keys []string) (IsSame bool, differences []keyDiffer) {
	IsSame = true
	differences = make([]keyDiffer, 0)
	for _, key := range keys {
		if FirstAnchor.GetResource(key) != SecondAnchor.GetResource(key) {
			IsSame = false
			differences = append(differences, keyDiffer{key, FirstAnchor.GetResource(key), SecondAnchor.GetResource(key)})
		}
	}
	return
}

//ChooseAction choose if to remove bucket or add one
func ChooseAction(r *rand.Rand, actions [2]string) string {
	return actions[r.Int()%2] //TODO: make this without the magic number 2
}

//ChooseBucketToRemove choose if to remove bucket or add one
func ChooseBucketToRemove(r *rand.Rand, currentWorkingSet []string, num int) string {
	var index uint64
	index = (uint64(r.Int()) % uint64(len(currentWorkingSet)))
	return currentWorkingSet[index]
}

/*
removeBucketFromSlice
this function remove a bucket from a working set slice.
@input: s : currentWorkingSet
*/
func removeBucketFromSlice(s []string, targetBucket string) error {
	i, err := indexOfBucket(targetBucket, s)
	if err != nil {
		return errors.New("bucket don't found on WorkingSet")
	}
	//removing bucket from working set , Order is not important
	s[len(s)-1], s[i] = s[i], s[len(s)-1] //swap the last element with the removed element
	s = s[:len(s)-1]
	return nil
}

func indexOfBucket(element string, data []string) (int, error) {
	for k, v := range data {
		if element == v {
			return k, nil
		}
	}
	return -1, errors.New("element don't found on slice") //not found.
}

//min return min num
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//calc a%n by the defnition that the output is a number between {0,1,...,a-1}
func positiveModuloOutput(a int, n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("wrong parameters") //not found.
	}
	res := a % n
	if a%n < 0 {
		res = -res
	}
	return res, nil
}
