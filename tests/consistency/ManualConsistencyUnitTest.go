package main

//hash import

//consistency test for anchor hash
//todo: parse parameters from the user
func main() {
	var numOfTests, maxSize, maxKeysNum int = 30, 2000, 10000
	var seed uint32 = 1
	consistencyTestLoop(numOfTests, maxSize, maxKeysNum, seed)
}
