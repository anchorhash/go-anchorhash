package main

//hash import

//balance test for anchor hash
//todo: parse parameters from the user
//todo: change the type's to int at balance test
func main() {
	var maxSize, wSize, seed uint32 = 10, 10, 1
	var numOfKeys int = 10000
	balanceTesting(maxSize, wSize, seed, numOfKeys)
}
