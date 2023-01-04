package fs

import (
	"fmt"
	"hash/crc32"
	"os"
)

func getHash(filename string, ch chan uint32) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		ch <- 0
	}
	x := crc32.NewIEEE()
	_, err = x.Write(bs)
	if err != nil {
		ch <- 0
	}
	fmt.Println(x.Sum32())
	ch <- x.Sum32()
}

func FilesHashCollect(filenames []string) []uint32 {
	//filenames := []string{
	//	"D:\\software\\.x\\2207\\5892dffa-916e-4c07-8292-b8a122d8d374",
	//	"D:\\software\\.x\\c\\98bdadbd-4231-4a89-99b1-1ce533ca0665",
	//	"D:\\software\\.x\\2207\\0a8cbe83-0fa1-41c1-ba21-30e971ab1728",
	//	"D:\\software\\.x\\2207\\3dc8a4b8-2880-4cbf-aad9-311fb9a2828d",
	//	"D:\\software\\.x\\2207\\4a6ef658-3d01-430d-ade5-b5eb89923840",
	//	"D:\\software\\.x\\2207\\4bc067d9-cb76-405b-a768-cb4774d40777",
	//	"D:\\software\\.x\\2207\\4e3ef3a0-3614-4278-bc89-c0b9727ad8e1",
	//	"D:\\software\\.x\\2207\\5c4291b3-a241-4408-ad6d-415980539a5b",
	//	"D:\\software\\.x\\2207\\5e65435d-8321-4017-8c3c-8ab3a0ffc7c1",
	//	"D:\\software\\.x\\2207\\6fc01373-f08e-4cc8-adae-a000141fd300",
	//}
	var result []uint32
	//var wg sync.WaitGroup
	//
	//wg.Add(int(MAX_FILE_HASHER))

	filesLength := len(filenames)

	ch := make(chan uint32, filesLength)
	b := filesLength / MAX_FILE_HASHER

	for i := 0; i < b; i++ {
		for j := 0; j < MAX_FILE_HASHER; j++ {
			go getHash(filenames[i*MAX_FILE_HASHER+j], ch)
			result = append(result, <-ch)
		}
	}
	for k := b * MAX_FILE_HASHER; k < filesLength; k++ {
		go getHash(filenames[k], ch)
		result = append(result, <-ch)
	}
	return result
}

//
//func Ex() {
//	start := time.Now()
//	result := getHash("D:\\software\\.x\\2207\\5892dffa-916e-4c07-8292-b8a122d8d374")
//	strs := strings.Split(result, ",")
//	h, err := strs[0], strs[1]
//	if err != "" {
//		fmt.Println(err)
//	}
//	fmt.Println("HASH: ", h)
//	fmt.Println("Time elapsed: ", time.Since(start).Milliseconds())
//}
