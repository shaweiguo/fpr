package main

import (
	"fmt"
	crawl "hyjj/pkg/crawl"
	"hyjj/pkg/fs"
	"os"
)

//var chanFileHash = make(chan uint32)

func ex_crawl() {
	sites := []string{
		"https://bing.com/",
		"https://godoc.org",
		"https://github.com/",
	}

	resps, err := crawl.Crawl(sites)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resps received:", resps)
}

func get_files(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	var result []string
	var file string
	for _, f := range files {
		//fmt.Printf("%s\\%s\n", dir, f.Name())
		file = fmt.Sprintf("%s\\%s\n", dir, f.Name())
		result = append(result, file)
	}
	fmt.Println(result)
	var hashs []uint32
	hashs = fs.FilesHashCollect(result)
	fmt.Println(hashs)
}

func ex_get_file_hashs() {
	filename := "D:\\sync\\temp\\Organization - 副本.txt"
	//hashes := fs.GetFileHashs0(filename)
	sha256 := fs.GetFileSha256Hash(filename)
	if sha256 == nil {
		fmt.Println("error")
		return
	}
	//for _, h := range hashes {
	fmt.Printf("%x\n", sha256)
	//}
}

func main() {
	//get_files("D:\\sync\\temp")
	ex_get_file_hashs()
}
