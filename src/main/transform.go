package main

import (
	"fmt"
)

func transform(input ResultOfParse) Transformed {
	fmt.Println("transform parsed")

	var trans = Transformed{}

	trans.caches = getCaches(input)
	trans.endpoints = getEndpoints(input)
	trans.maxCachedVideosSize = input.size
	return trans
}

func getCaches(input ResultOfParse) []Cache {
	var result []Cache
	var cahesInfo []CacheInfo
	for _,v := range input.endpointsInfo {
		for _,v2 := range v.cacheInfo{
			cahesInfo = append(cahesInfo,v2)
		}

	}

	println(len(cahesInfo ))
	for _, cache := range cahesInfo {
		println("asd")
		result = append(result, Cache{id: cache.id, videos: make([]Video, 0), latency: cache.latency,size:input.size})
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
