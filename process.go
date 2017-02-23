package main

import (
	"fmt"
	"sort"
)

func process(input Transformed) Transformed {
	fmt.Println("process transformed")

	final := initialOrder(input)
	for _,endpoint := range final.endpoints {
		cache := getCacheWithSpace(final)
		for _,video := range endpoint.videos {
			cache.size = cache.size - video.size
			cache.videos = append(cache.videos, video)
		}
	}

	return Transformed{}
}

func getCacheWithSpace(input Transformed) Cache {

	for _,v := range input.caches {
		if v.size > 0 {
			return v
		}
	}

}

func initialOrder(result Transformed) Transformed {
	for _, endpoint := range result.endpoints {
		sort.Sort(ByRequest(endpoint.videos))
	}
	sort.Sort(ByLatency(result.caches))
	return result
}

type ByRequest []Video
type ByLatency []Cache

func (a ByRequest) Len() int           { return len(a) }
func (a ByRequest) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRequest) Less(i, j int) bool { return a[i].numRequest < a[j].numRequest }

func (a ByLatency) Len() int           { return len(a) }
func (a ByLatency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLatency) Less(i, j int) bool { return a[i].latency > a[j].latency }
