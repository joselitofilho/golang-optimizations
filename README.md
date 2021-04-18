# golang-optimization

**Table of Contents**

- [Overview](#overview)
  - [Clousures](clousures/README.md)
  - [Arrays and Slices](memoryblock/README.md)
  - [Memory Leaking](memoryleak/README.md)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
  - [Using Docker](#using-docker)
- [Run Profile](#run-profile)

## Overview

Building systems with high performance, very low response times, supporting the heavy load on little hardware has never been easier. It is necessary to a little bit more about the language, some concepts, and optimize our own code to extract more performance.

I separated some concepts for us to analyze and study together.
- [Pointers](pointers/README.md)
- [Arrays and Slices](memoryblock/README.md)
- [Clousures](clousures/README.md)
- [Memory Leaking](memoryleak/README.md)
- ...

Enjoy!

## Prerequisites

- Go 1.15.8
- Git
- Docker 19.x

## Setup

### Using Docker

#### Build an image
```bash
docker build -t golang-dev --build-arg GO_VERSION=1.15.8 .
```
#### Create a new container

Linux and MacOS:
```bash
docker run -v `pwd`:/src -w /src --label com.docker.compose.project=golang-optimizations -it --name ${PWD##/*} golang-dev
```

Windows:
```bash
docker run -v $pwd\:/src -w /src --label com.docker.compose.project=golang-optimizations -it --name golang-optimizations golang-dev
```

## Run profile

We are using [pprof](https://github.com/google/pprof) for visualization and analysis of profiling data.

Generating memory profile:
```bash
go test -memprofile="mem.prof" -benchtime=1000x -bench .
```

Seeing profile on text mode:
```bash
go tool pprof -text mem.prof
```

Saving inuse_space profile on png:
```bash
go tool pprof -sample_index=inuse_space -png mem.prof
```

Seeing profile on graph mode:
```bash
apt-get install -y graphviz gv
go tool pprof -http=localhost:8080 mem.prof
```

Seeing Garbage Collector steps:
```bash
GODEBUG=gctrace=1 go test -memprofile="mem.prof" -benchtime=10x -bench .

# gc 1 @6.068s 11%: 0.058+1.2+0.083 ms clock, 0.70+2.5/1.5/0+0.99 ms cpu, 7->11->6 MB, 10 MB goal, 12 P

# // General
# gc 1        : The 1 GC run since the program started
# @6.068s     : Six seconds since the program started
# 11%         : Eleven percent of the available CPU so far has been spent in GC

# // Wall-Clock
# 0.058ms     : STW        : Mark Start       - Write Barrier on
# 1.2ms       : Concurrent : Marking
# 0.083ms     : STW        : Mark Termination - Write Barrier off and clean up

# // CPU Time
# 0.70ms      : STW        : Mark Start
# 2.5ms       : Concurrent : Mark - Assist Time (GC performed in line with allocation)
# 1.5ms       : Concurrent : Mark - Background GC time
# 0ms         : Concurrent : Mark - Idle GC time
# 0.99ms      : STW        : Mark Term

# // Memory
# 7MB         : Heap memory in-use before the Marking started
# 11MB        : Heap memory in-use after the Marking finished
# 6MB         : Heap memory marked as live after the Marking finished
# 10MB        : Collection goal for heap memory in-use after Marking finished

# // Threads
# 12P         : Number of logical processors or threads used to run Goroutines
```
For more information about `gctrace` in Golang visit [here](https://golang.org/pkg/runtime/).
