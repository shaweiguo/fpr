package fs

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var fileTypeMap map[string]string

func init() {
	fileTypeMap = make(map[string]string)
	fileTypeMap["00000020"] = "mp4"
	fileTypeMap["2E524D46"] = "rm"
	fileTypeMap["000001BA"] = "mpg"
	fileTypeMap["6D6F6F76"] = "mov"
	fileTypeMap["41564920"] = "avi"
	fileTypeMap["464C5601"] = "flv"
}

func bytesToHexString(b []byte) string {
	res := bytes.Buffer{}
	if b == nil || len(b) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	i, length := 100, len(b)
	if length < i {
		i = length
	}
	for j := 0; j < i; j++ {
		sub := b[j] & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

func getFileType(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	buf := make([]byte, 20)
	n, _ := f.Read(buf)

	fileCode := bytesToHexString(buf[:n])
	for k, v := range fileTypeMap {
		if strings.HasPrefix(fileCode, k) {
			return v
		}
	}
	return ""
}

func fileType(head *[]byte) string {
	fileCode := bytesToHexString(*head)
	for k, v := range fileTypeMap {
		if strings.HasPrefix(fileCode, k) {
			return v
		}
	}
	return ""
}

func FileCrawl(ctx context.Context, root string, vfCh chan string, exitCh chan bool) (err error) {
	if ctx.Err() != nil {
		return err
	}

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-exitCh:
		fmt.Println("exit")
	default:
		err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			ext := getFileType(path)
			if ext != "" {
				vfCh <- path + "," + ext
				return nil
			} else {
				return errors.New("file type unknown")
			}
		})
	}
	return
}
