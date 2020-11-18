package anchor

import (
	"hash"

	"github.com/golang-collections/collections/stack" //stack import
	"github.com/spaolacci/murmur3"                    //hash import
)

//HashImp is the anchor hash implementation struct
type HashImp struct {
	a, l, w, k []uint32
	r          stack.Stack
	n          uint32 //num of the current working set size.
	maxSize    uint32 //the max bucket capacity |A|
	seed       uint32
	hasher     hash.Hash32
}

//newHashImp returns a new HashImp object
func newHashImp(maxSize uint32, wSize uint32, seed uint32) *HashImp {
	anchor := HashImp{
		a:       make([]uint32, maxSize),
		l:       make([]uint32, maxSize),
		w:       make([]uint32, maxSize),
		k:       make([]uint32, maxSize),
		maxSize: maxSize,
		n:       wSize,
		seed:    seed,
		hasher:  murmur3.New32WithSeed(seed),
	}
	for i := uint32(0); i < maxSize; i++ {
		anchor.k[i], anchor.w[i], anchor.l[i] = i, i, i
	}
	for i := maxSize - 1; i >= wSize; i-- {
		anchor.r.Push(i)
		anchor.a[i] = i
	}
	return &anchor
}

//getBucket receives a key k as an input
//and returns a working bucket number as an output.
func (i *HashImp) getBucket(k string) uint32 {
	b := i.digest(k) % i.maxSize
	for i.a[b] > 0 {
		h := murmur3.Sum32WithSeed([]byte(k), i.seed+uint32(b)) % i.a[b]
		for i.a[h] >= i.a[b] {
			h = i.k[h]
		}
		b = h
	}
	return b
}

//addBucket adds a bucket to the working set
//and returns the added bucket number.
func (i *HashImp) addBucket() (b uint32) {
	b = i.r.Pop().(uint32)
	i.a[b] = uint32(0)
	i.l[i.w[i.n]] = i.n
	i.w[i.l[b]], i.k[b] = b, b
	i.n++
	return b
}

//removeBucket receives as an input a bucket
//to remove and removes it from the working set.
func (i *HashImp) removeBucket(b uint32) {
	i.r.Push(b)
	i.n--
	i.a[b] = i.n
	i.w[i.l[b]], i.k[b] = i.w[i.n], i.w[i.n]
	i.l[i.w[i.n]] = i.l[b]
}

//digest receives a string and returns a digested uint64 number
func (i *HashImp) digest(s string) uint32 {
	i.hasher.Write([]byte(s))
	res := i.hasher.Sum32()
	i.hasher.Reset()
	return res
}
