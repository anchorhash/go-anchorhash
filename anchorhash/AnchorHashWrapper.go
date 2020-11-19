//this file implement functions from algorithm 2

package anchor

import (
	"errors"
	"fmt"
)

//HashWrapper is the anchor Hash Wrapper struct
type HashWrapper struct {
	anchor           *HashImp
	resourceToBucket map[string]uint32 //M (in algorithm 2)
	bucketToResource []string          //M^-1 (in algorithm 2)
}

//NewHashWrapper returns a new anchor HashWrapper.
func NewHashWrapper(maxSize uint32, w []string,
	seed uint32) (*HashWrapper, error) {
	wSize := uint32(len(w))
	if maxSize < wSize {
		return nil, fmt.Errorf("bad parameters: (%d=) |A|<|W| (=%d)",
			maxSize, wSize)
	}
	a := HashWrapper{
		anchor:           newHashImp(maxSize, wSize, seed),
		resourceToBucket: make(map[string]uint32),
		bucketToResource: make([]string, maxSize),
	}
	for i, name := range w {
		if w[i] == "" {
			return nil, fmt.Errorf("bad parameters: empty string at index:%d", i)
		}
		if _, exists := a.resourceToBucket[name]; exists {
			return nil, fmt.Errorf("bad parameters: duplicated string:'%v'", name)
		}
		a.bucketToResource[i] = w[i]
		a.resourceToBucket[name] = uint32(i)
	}
	return &a, nil
}

//GetResource receives a key as an input.
//returns working bucket string as an output.
func (a *HashWrapper) GetResource(k string) string {
	b := a.anchor.getBucket(k)
	return a.bucketToResource[b]
}

//AddResource receives a resource name `xi` as an input
//and add it as a resource to the working set.
func (a *HashWrapper) AddResource(xi string) error {
	if a.anchor.n >= a.anchor.maxSize {
		return errors.New("anchor is full")
	}
	if _, exists := a.resourceToBucket[xi]; exists {
		return fmt.Errorf("bad parameters:'%v' already exists", xi)
	}
	b := a.anchor.addBucket()
	a.bucketToResource[b] = xi
	a.resourceToBucket[xi] = b
	return nil
}

//RemoveResource receives a resource name as an input
//and remove it from the working set.
func (a *HashWrapper) RemoveResource(xi string) error {
	if a.anchor.n == 1 {
		return errors.New("anchor can't remove the last resource")
	}
	b, exists := a.resourceToBucket[xi]
	if !exists {
		return fmt.Errorf("bad parameters: resource to remove:'%v' not exists", xi)
	}
	delete(a.resourceToBucket, xi)
	a.bucketToResource[b] = ""
	a.anchor.removeBucket(b)
	return nil
}
