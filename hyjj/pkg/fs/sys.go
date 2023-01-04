package fs

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func GetSysinfo() {
	cmd := exec.Command("wmic", "csproduct", "get", "UUID")
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("result: %s\n", string(result))
		log.Fatal("cmd.Run failed: %s\n", err)
	}
	fmt.Printf("result: %s\n", string(result))
}

func bitsToDrives(bitMap uint32) (drives []string) {
	availableDrives := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i := range availableDrives {
		if bitMap&1 == 1 {
			drives = append(drives, availableDrives[i])
		}
		bitMap >>= 1
	}

	return
}

func GetDrivers() (drivers []string) {
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getLogicalDrivesHandle, _ := syscall.GetProcAddress(kernel32, "GetLogicalDrives")
	if ret, _, err := syscall.SyscallN(uintptr(getLogicalDrivesHandle), 0, 0, 0, 0); err != 0 {
		log.Fatal("Error: ", err)
	} else {
		drivers = bitsToDrives(uint32(ret))
	}
	return
}
