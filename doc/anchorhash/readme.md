# Implementation of AnchorHash

This repository contains an implementation of algorithms 2 and 3 from  the paper  "AnchorHash: A Scalable Consistent Hash"
<!---(add link here)-->
## algorithm 2 — AnchorHash Wrapper
### Pseudocode:
![Alt text](algorithm2-wrapper.png?raw=true "Title")
<!---
implementation of functions from algorithm 2:
-->
### implementation:
#### Init wrapper(A,W)
<!---
### $Init wrapper(\mathcal{A},\mathcal{S})$ (`NewHashWrapper`)-->
<!---
#### Pseudocode:
Enter Pseudocode here
-->
<!---
>**function** INITWRAPPER$\left(\mathcal{A},\mathcal{W}\right)$: \
>&ensp;$M\leftarrow\emptyset,\mathcal{W}\leftarrow\emptyset$  \
>&ensp;**for** i $\in\left(0,1,2,\ldots,\left|\mathcal{S}\right|-1\right)$ **do** \
>&ensp;&ensp;$M\leftarrow M\cup\left\{ \left(\mathcal{A}\left[i\right],\mathcal{S}\left[i\right]\right)\right\}$ \
>&ensp;&ensp;$\mathcal{W}\leftarrow\mathcal{W}\cup\mathcal{W}\left\{ \mathcal{A}\left[i\right]\right\}$  \
>&ensp;INITANCHOR$\left(\mathcal{A},\mathcal{W}\right)$

#### implementation:
-->
```go
func NewHashWrapper(maxSize uint32, w []string,
	seed uint32) (*HashWrapper, error)
```
`NewHashWrapper` returns a new `HashWrapper` object (anchor hash wrapper object) or an error, if any.\
`NewHashWrapper` receives as an input :
* `maxSize` - the Anchor max size.in other words `maxSize` is the upper bound on the number of buckets that anchor hash allow (=|A|).
* `w` - initial working set of buckets (=W).
  * `w` should contain at most `maxSize` elements and at least 1 element.
  * all the elements on `w` must be non-empty strings.
  * all the elements on `w` must be unique(`w` doesn't allow duplicated elements).
* `seed` - the seed of the anchor

#### GetResource(k) <!---(`GetResource`)-->
<!---
#### Pseudocode:
Enter Pseudocode here

#### implementation:
-->
```go
func (a *HashWrapper) GetResource(k string) string
```
`GetResource` receives a key `k` as an input and returns a working bucket string as an output.

#### AddResource(xi) <!---(`AddResource`)-->
<!---
#### Pseudocode:
Enter Pseudocode here

#### implementation:
-->
```go
func (a *HashWrapper) AddResource(xi string) error
```
`AddResource` receives a resource name `xi` as an input and add it as a resource to the working set.\
`AddResource` will fail and return an error if:
* `HashWrapper` is full. <!---$(|\mathcal{W}|=|\mathcal{A}|)$-->
* `HashWrapper` has `xi` as a working bucket already.<!---$(\xi\in\mathcal{W})$-->

#### RemoveResource(xi) <!---(`RemoveResource`)-->
<!---
#### Pseudocode:
Enter Pseudocode here

#### implementation:
-->
```go
func (a *HashWrapper) RemoveResource(xi string) error
```
`RemoveResource` receives a resource name `xi` as an input
and remove it from the working set.\
`RemoveResource` will fail and return error if:
* `xi` isn't in the working bucket already. <!---$(\xi\notin\mathcal{W})$-->
* `HashWrapper` has only one bucket in the working set. <!---$(|\mathcal{W}|=1)$-->

## algorithm 3 — AnchorHash Implementation
### Pseudocode:
![Alt text](algorithm3-implementation.png?raw=true "Title")
### implementation:
#### InitAnchor(a,w) <!---(`newHashImp`)-->
<!---
#### Pseudocode:
Enter Pseudocode here

#### implementation:
-->
```go
func newHashImp(maxSize uint32, wSize uint32, seed uint32) *HashImp 
```
`newHashImp` returns a new `HashImp` object (anchor hash implementation object).\
`newHashImp` receives as an input :
* `maxSize` - the Anchor max size  (=|A|).
* `wSize` - initial working buckets size (|w|).
  * `wSize` should be smaller or equal to `maxSize`.
* `seed` -the seed of the anchor
<!---
Enter implementation here
-->

#### GetBucket(k) <!---(`getBucket`)-->
<!---
#### Pseudocode:
Enter Pseudocode here

#### implementation:
-->
```go
func (i *HashImp) getBucket(k string) uint32
```
`getBucket` receives a key `k` as an input and returns a working bucket number as an output.
<!---
Enter implementation here
-->
#### AddBucket( )
<!---
#### Pseudocode:

Enter Pseudocode here

#### implementation:
-->
```go
func (i *HashImp) addBucket() (b uint32)
```
`addBucket` adds a bucket to the working set and returns the added bucket number.
<!---
Enter implementation here
-->
#### RemoveBucket(b)
<!---
#### Pseudocode:

Enter Pseudocode here

#### implementation:
-->
```go
func (i *HashImp) removeBucket(b uint32)
```
`removeBucket` receives as an input a bucket to remove `b` and removes it from the working set. <!--- TODO: fix grammar her-->
<!---
Enter implementation here
-->