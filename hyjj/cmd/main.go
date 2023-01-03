package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"os/exec"
	"syscall"
)

//var chanFileHash = make(chan uint32)

func getHash(filename string, chanFileHash chan string) {
	bs, err := os.ReadFile(filename)
	result := ""
	if err != nil {
		result = fmt.Sprintf("%d,%s", 0, err)
		chanFileHash <- result
		//return result
	}
	x := crc32.NewIEEE()
	x.Write(bs)
	v := x.Sum32()
	result = fmt.Sprintf("%d,", v)
	chanFileHash <- result
	//return result
}

//func fileHashCollector(fileNames []string) []string {
//	filenames := []string{
//		"D:\\software\\.x\\2207\\5892dffa-916e-4c07-8292-b8a122d8d374",
//		"D:\\software\\.x\\c\\98bdadbd-4231-4a89-99b1-1ce533ca0665",
//		"D:\\software\\.x\\2207\\0a8cbe83-0fa1-41c1-ba21-30e971ab1728",
//		"D:\\software\\.x\\2207\\3dc8a4b8-2880-4cbf-aad9-311fb9a2828d",
//		"D:\\software\\.x\\2207\\4a6ef658-3d01-430d-ade5-b5eb89923840",
//		"D:\\software\\.x\\2207\\4bc067d9-cb76-405b-a768-cb4774d40777",
//		"D:\\software\\.x\\2207\\4e3ef3a0-3614-4278-bc89-c0b9727ad8e1",
//		"D:\\software\\.x\\2207\\5c4291b3-a241-4408-ad6d-415980539a5b",
//		"D:\\software\\.x\\2207\\5e65435d-8321-4017-8c3c-8ab3a0ffc7c1",
//		"D:\\software\\.x\\2207\\6fc01373-f08e-4cc8-adae-a000141fd300",
//	}
//	var result []string
//
//	wg := &sync.WaitGroup{}
//	chanLimiter := make(chan bool, 5)
//	defer close(chanLimiter)
//
//	chanFileHash := []chan { make(chan string, len(filenames))}
//	wgResponse := &sync.WaitGroup{}
//	for i := 0; i < len(filenames); i++ {
//		go func(i uint8) {
//			wgResponse.Add(1)
//			getHash(filenames[i], chanFileHash)
//		}(i)
//	}
//}

func get_sysinfo() {
	cmd := exec.Command("wmic", "csproduct", "get", "UUID")
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("result: %s\n", string(result))
		log.Fatal("cmd.Run failed: %s\n", err)
	}
	fmt.Printf("result: %s\n", string(result))
}

func bitsToDrives(bitMap uint32) (drives []string) {
	availableDrives := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i := range availableDrives {
		if bitMap&1 == 1 {
			drives = append(drives, availableDrives[i])
		}
		bitMap >>= 1
	}

	return
}

func getDrivers() (drivers []string) {
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getLogicalDrivesHandle, _ := syscall.GetProcAddress(kernel32, "GetLogicalDrives")
	if ret, _, err := syscall.SyscallN(uintptr(getLogicalDrivesHandle), 0, 0, 0, 0); err != 0 {
		log.Fatal("Error: ", err)
	} else {
		drivers = bitsToDrives(uint32(ret))
	}
	return
}

func main() {
	//start := time.Now()
	//result := getHash("D:\\software\\.x\\2207\\5892dffa-916e-4c07-8292-b8a122d8d374")
	//strs := strings.Split(result, ",")
	//h, err := strs[0], strs[1]
	//if err != "" {
	//	fmt.Println(err)
	//}
	//fmt.Println("HASH: ", h)
	//fmt.Println("Time elapsed: ", time.Since(start).Milliseconds())
	drivers := getDrivers()
	fmt.Println(drivers)
}
