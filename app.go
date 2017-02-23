package main

import ("fmt"
	"os"
	"bufio"

	"strings"
	"strconv"
)


func assert(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	fmt.Println("hello world")

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//assert(err)

	dir := "inputs/"

	files := []string{"m.in"}

	for _,v  := range files {
		readGoogleHashcodeFile(dir+v)
		writeSolution("solution_"+v)
	}



}

type Video struct {
	id int
	size int
}

/*type Endpoint struct {
	id int
	videos []Video
}*/

type Endpoint struct{
	id int
	cacheInfo []CacheInfo
	latency int
}

type CacheInfo struct{

	id int
	latency int

}

type RequestInfo struct {
	numRequest int
	idVideo int
	endpointId int
}



const MAX_CACHED_VIDEOS_SIZE = 500000

type Cache struct {
	id int
	videos []Video
}

type Requests struct {
	latency int

}

type ResultOfParse struct {
	videos []Video
	caches []CacheInfo
	requestInfo []RequestInfo
	endpoints int
	numberOfCaches int
	size int
}


func readGoogleHashcodeFile( path string) (ResultOfParse){


	file,error := os.Open(path)

	assert(error)

	reader := bufio.NewReader(file)

	line, _, error := reader.ReadLine()

	counter := 0


	numberOfVideos := 0
	numberOfEnpoints := 0
	numberOfRequest := 0
	numberOfCaches := 0
	cacheSize:= 0

	var videos  = []Video{}

	endpointCacheNumber := 0
	enpointDetailIteration := 0
	var currentEndpoint Endpoint;
	enpointNumber := 0
	var cacheInfo  = []CacheInfo{}

	enpointIteration := 0

	var requests  = []RequestInfo{}
	requestsIteration := 0

	for line != nil {
		assert(error)
		//process line
		linesplited := strings.Split(string(line), " ")

		if (counter == 0){
			println("First line")
			println(string(line))
			numberOfVideos = convertToInt(linesplited[0])
			numberOfEnpoints = convertToInt(linesplited[1])
			numberOfRequest = convertToInt(linesplited[2])
			numberOfCaches = convertToInt(linesplited[3])
			cacheSize = convertToInt(linesplited[4])

		}else if (counter == 1){
			println("Second line")
			println(string(line))
			println(numberOfVideos)
			println(len(linesplited))
			//panic(numberOfVideos  != len(linesplited))
			total := numberOfVideos - 1
			for i := 0; i < total; i++ {
				println(i)
				println(total)
				videos = append(videos,Video{id:i,size:convertToInt(linesplited[i])})
			}
		}else {
			// ENDPOINTS
			//println("endpoits")
			println(string(line))
			if enpointIteration < numberOfEnpoints   {
				//print("enpointIteration")
				println("enpointIteration"+strconv.Itoa(enpointIteration))

				if  endpointCacheNumber != 0 && enpointDetailIteration < endpointCacheNumber  {
					//print("enpointDetailIteration")
					println("enpointIteration"+strconv.Itoa(enpointDetailIteration))

					enpointDetailIteration = enpointDetailIteration +1
					cacheInfo = append(cacheInfo,CacheInfo{id:convertToInt(linesplited[0]),latency:convertToInt(linesplited[1])})
				}else{

					if len(cacheInfo) > 0 {
						currentEndpoint.cacheInfo = cacheInfo
					}

					currentEndpoint = Endpoint{id:enpointNumber,latency:convertToInt(linesplited[0])}
					endpointCacheNumber = convertToInt(linesplited[1])
					enpointNumber = enpointNumber +1
					enpointDetailIteration = 0
					cacheInfo  = []CacheInfo{}

					enpointIteration = enpointIteration + 1
				}
			}else {

				if len(cacheInfo) > 0 {
					currentEndpoint.cacheInfo = cacheInfo
				}

				println("requjest")
				// REQUEST
				requests := append(requests,RequestInfo{numRequest:convertToInt(linesplited[0]),idVideo:convertToInt(linesplited[1]),endpointId:convertToInt(linesplited[2])})

				requestsIteration = requestsIteration + 1

				_ = requests

			}



		}

		counter = counter + 1
		line, _, _ = reader.ReadLine()
	}

	//panic(requestsIteration != numberOfRequest)

	if(requestsIteration != numberOfRequest){
		panic("not iqual ")
	}


	file.Close()

	return ResultOfParse{caches:cacheInfo,requestInfo:requests,videos:videos,endpoints:numberOfEnpoints,numberOfCaches:numberOfCaches,size:cacheSize}
}

func convertToInt(str string) (int){
	intRes,_ := strconv.Atoi(str)
	return intRes
}


func writeSolution(result string){

	file,err := os.Create(result)
	assert(err)

	defer file.Close()

	file.WriteString("writes\n")

	file.Sync()

}