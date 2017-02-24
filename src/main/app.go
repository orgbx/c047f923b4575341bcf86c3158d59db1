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

	files := []string{"me_at_the_zoo.in"}

	trans := Transformed{}
	trans.caches = []Cache{Cache{id:1,size:12,videos:[]Video{Video{size:2,id:3},Video{size:2,id:4}}},Cache{id:3,size:12,videos:[]Video{Video{size:2,id:5},Video{size:2,id:7}}}}
	for _,v  := range files {
		res := readGoogleHashcodeFile(dir+v)

		trs := transform(res)
		println(len(trs.endpoints))
		finaRes := process(trs)
		writeSolution("solution_"+v,finaRes)
	}



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
	var currentEndpoint EndpointInfo;
	enpointNumber := 0
	var cacheInfo  = []CacheInfo{}

	var endpoints  = []EndpointInfo{}

	enpointIteration := 0

	var requests  = []RequestInfo{}
	requestsIteration := 0

	notFirst := false

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
			if enpointIteration < numberOfEnpoints     {
				//print("enpointIteration")
				println("enpointIteration"+strconv.Itoa(enpointIteration))

				if   enpointDetailIteration < endpointCacheNumber  {
					//print("enpointDetailIteration")
					println("enpointDetailIteration"+strconv.Itoa(enpointDetailIteration))
					println("endpointCacheNumber"+strconv.Itoa(endpointCacheNumber))
					enpointDetailIteration = enpointDetailIteration +1
					currentEndpoint.cacheInfo  = append(currentEndpoint.cacheInfo ,CacheInfo{id:convertToInt(linesplited[0]),latency:convertToInt(linesplited[1])})
					cacheInfo   = append(cacheInfo ,CacheInfo{id:convertToInt(linesplited[0]),latency:convertToInt(linesplited[1])})
					println("Append result "+strconv.Itoa(len(currentEndpoint.cacheInfo)))
					notFirst = true
				}else{

					if(notFirst){
						endpoints = append(endpoints,currentEndpoint)
					}

					currentEndpoint = EndpointInfo{id:enpointNumber,latency:convertToInt(linesplited[0])}
					//currentEndpoint.cacheInfo = cacheInfo
					println(len(currentEndpoint.cacheInfo))


					endpointCacheNumber = convertToInt(linesplited[1])
					enpointNumber = enpointNumber +1
					enpointDetailIteration = 0
					cacheInfo   = []CacheInfo{}
					if(notFirst){
						enpointIteration = enpointIteration + 1
					}





					//if ( convertToInt(linesplited[0]) == 696 &&  convertToInt(linesplited[1]) == 2) {
		//						panic("s")
					//}



				}
			}else {

				if len(cacheInfo) > 0 {
					currentEndpoint.cacheInfo = cacheInfo
				}

				println("requjest")
				// REQUEST
				requests = append(requests,RequestInfo{numRequest:convertToInt(linesplited[0]),idVideo:convertToInt(linesplited[1]),endpointId:convertToInt(linesplited[2])})

				requestsIteration = requestsIteration + 1



			}



		}

		counter = counter + 1
		line, _, _ = reader.ReadLine()
	}

	//panic(requestsIteration != numberOfRequest)

	if(requestsIteration != numberOfRequest -1 ){
		panic("not iqual ")
	}


	file.Close()

	return ResultOfParse{requestInfo:requests,videos:videos,endpoints:numberOfEnpoints,numberOfCaches:numberOfCaches,size:cacheSize,endpointsInfo:endpoints}
}

func convertToInt(str string) (int){
	intRes,_ := strconv.Atoi(str)
	return intRes
}


func writeSolution(result string,val Transformed){

	file,err := os.Create(result)
	assert(err)

	defer file.Close()
	file.WriteString(strconv.Itoa(len(val.caches))+"\n")

	for i2, v := range val.caches {
		file.WriteString(strconv.Itoa(v.id)+" ")
		for i,v2 := range v.videos {
			if (i < len(v.videos) - 1) {
				file.WriteString(strconv.Itoa(v2.id)+" ")
			}else {
				file.WriteString(strconv.Itoa(v2.id))
				if (i2 < len(val.caches) - 1) {
					file.WriteString("\n")
				}
			}

		}

	}



	file.Sync()

}