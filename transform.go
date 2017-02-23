package main

import (
	"fmt"
)

func transform(input ResultOfParse) Transformed {
	fmt.Println("transform parsed")

	var trans = Transformed{}

	trans.caches = getCaches(input.caches)
	trans.endpoints = getEndpoints(input)
	trans.maxCachedVideosSize = input.size
	return trans
}

func getCaches(input []CacheInfo) []Cache {
	var result []Cache
	for _, cache := range input {
		result = append(result, Cache{id: cache.id, videos: make([]Video, 0), latency: cache.latency})
	}
	return result
}

func getEndpoints(input ResultOfParse) []Endpoint {
	// var videos map[int]Video
	var result []Endpoint

	for _, endpointsInfo := range input.endpointsInfo {
		var caches []int
		for _, cache := range endpointsInfo.cacheInfo {
			caches = append(caches, cache.id)
		}
		var endpoint = Endpoint{id: endpointsInfo.id, cacheIds: caches}
		result = append(result, endpoint)
	}

	for _, requestInfo := range input.requestInfo {
		var video = Video{id: requestInfo.idVideo, numRequest: requestInfo.numRequest}
		for _, videoInfo := range input.videos {
			if video.id == videoInfo.id {
				video.numRequest = videoInfo.size
			}
			break
		}
		for _, endpoint := range result {
			if endpoint.id == requestInfo.endpointId {
				endpoint.videos = append(endpoint.videos, video)
			}
			break
		}
	}
	// println(string(result))

	return result
}
