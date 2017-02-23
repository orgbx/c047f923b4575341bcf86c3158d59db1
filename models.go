package main

type Video struct {
	id         int
	size       int
	numRequest int
}

type Endpoint struct {
	id       int
	videos   []Video
	cacheIds []int
}

type EndpointInfo struct {
	id        int
	cacheInfo []CacheInfo
	latency   int
}

type CacheInfo struct {
	id      int
	latency int
}

type RequestInfo struct {
	numRequest int
	idVideo    int
	endpointId int
}

const MAX_CACHED_VIDEOS_SIZE = 500000

type Cache struct {
	id      int
	videos  []Video
	latency int
}

type Requests struct {
	latency int
}

type ResultOfParse struct {
	videos         []Video
	caches         []CacheInfo
	requestInfo    []RequestInfo
	endpoints      int
	numberOfCaches int
	size           int
}

type Transformed struct {
	endpoints           []Endpoint
	caches              []Cache
	maxCachedVideosSize int
}
