# AnchorHash: A Scalable Consistent Hash

This repository contains supplementary code for the paper  "AnchorHash: A Scalable Consistent Hash" (add link here)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.


### Prerequisites

<!--get all remote dependencies.-->
* get all dependencies:
    ```
  go get github.com/golang-collections/collections/stack
  go get github.com/spaolacci/murmur3
    ```
* if you went to use the benchmarks analysis `benchstat` you can download it useing:
    ```
    go get golang.org/x/perf/cmd/benchstat
    ```
  and make sure it is installed using command:
  ```
  benchstat
  ```
### Installation
* Clone the repository (inside your local GOPATH or GOROOT) and cd into it:
  ```
  git clone https://github.com/anchorhash #inside your local GOPATH or GOROOT
  cd anchorhash
  ```
<!--
### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be
(all the git clone and go get and etc..)

```
Give the example
```

And repeat

```
until finished
```
-->


## Running the tests
<!--Explain how to run the automated tests for this system-->
you can run the tests by executing the following commands from the root directory.
### rate benchmarks
* running the rate benchmarks: (WIP this part should be according to new directories structure)
  * ratio benchmark
    ```
    go test ./benchmarks/rate -parallel 1 -benchtime=50x -cpu 1 -count 5 -bench=BenchmarkRatioRate >ratio_benchmark.out
    benchstat ratio_benchmark.out
    ```
  * load benchmark
    ```
    go test ./benchmarks/rate -parallel 1 -benchtime=50x -cpu 1 -count 5 -bench=BenchmarkLoadRate >load_benchmark.out
    benchstat load_benchmark.out
    ```
* running the consistency test:
  ```
  go run tests/consistency/ManualConsistencyUnitTest.go tests/consistency/TestConsistency.go
  ```
* running the balance test:
  ```
  go run tests/balance/ManualBalanceUnitTest.go tests/balance/TestBalance.go
  ```
## code
* [anchor hash implementation](anchorhash)
  * [anchor hash documentation](doc/anchorhash/readme.md)
* tests and benchmarks:
  * [Consistency test](tests/consistency) [WIP]
  * [Balance test](tests/balance)
  * [Rate benchmarks](benchmarks/rate)

## Authors

add Authors

## License

add License

## Acknowledgments
