package fs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"log"
	"os"
	"sync"
)

type HashMethod struct {
	Name string
	New  func() hash.Hash
	Use  bool
}

type HashWriter0 struct {
	Writer    hash.Hash
	Channel   chan []byte
	Category  string
	onceClose *sync.Once
	wg        *sync.WaitGroup
}

func (this *HashWriter0) write() {
	doWriteHash(this.Writer, this.Channel)
	fmt.Printf("%s exit", this.Category)
	this.wg.Done()
}

func (this *HashWriter0) Write(buf []byte) {
	tmpBuf := make([]byte, len(buf))
	copy(tmpBuf, buf)
	if this.Channel != nil {
		this.Channel <- tmpBuf
	}
}

func (this *HashWriter0) Close() {
	this.onceClose.Do(func() {
		if this.Channel != nil {
			close(this.Channel)
		}
	})
}

func (this *HashWriter0) Sum(b []byte) (sum []byte) {
	this.Close()

	this.wg.Wait()

	if this.Writer != nil {
		sum = this.Writer.Sum(b)
	}
	return
}
func NewSha256() *HashWriter0 {
	writer := new(HashWriter0)
	writer.onceClose = new(sync.Once)
	writer.wg = new(sync.WaitGroup)

	var hashCount int
	chanCount := 30

	hashCount++
	writer.Writer = sha256.New()
	writer.Channel = make(chan []byte, chanCount)
	go writer.write()
	writer.wg.Add(hashCount)
	return writer
}
func GetFileSha256Hash(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	hw := NewSha256()
	defer hw.Close()

	b := make([]byte, 32*1024)
	for {
		n, err := f.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil
		}
		hw.Write(b[:n])
	}

	return hw.Sum(nil)
}

func NewHashWriter0(hashes []HashMethod) []*HashWriter0 {
	var result []*HashWriter0
	var hashCount int
	chanCount := 30
	var writer *HashWriter0
	//once := new(sync.Once)
	wg := new(sync.WaitGroup)

	fmt.Printf("\nhashes: %v", hashes)
	for _, h := range hashes {
		fmt.Printf("\n%s: %+v", h.Name, h)
		if h.Use {
			writer = new(HashWriter0)
			writer.onceClose = new(sync.Once)
			writer.wg = wg
			hashCount++
			writer.Category = h.Name
			writer.Writer = h.New()
			writer.Channel = make(chan []byte, chanCount)
			go writer.write()
		} else {
			writer = nil
		}
		result = append(result, writer)
		fmt.Printf("\nresult1: %+v", result)
	}
	wg.Add(hashCount)
	fmt.Printf("\nresult2: %#v", result)

	return result
}

func GetFileHashs0(filename string) (result [][]byte) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var methods []HashMethod
	md5 := HashMethod{
		Name: "md5",
		New:  md5.New,
		Use:  true,
	}
	sha1 := HashMethod{
		Name: "sha1",
		New:  sha1.New,
		Use:  true,
	}
	sha256 := HashMethod{
		Name: "sha256",
		New:  sha256.New,
		Use:  true,
	}
	methods = append(methods, md5)
	methods = append(methods, sha1)
	methods = append(methods, sha256)

	hw := NewHashWriter0(methods)
	defer hw[0].Close()
	defer hw[1].Close()
	defer hw[2].Close()

	b := make([]byte, 32*1024)
	for {
		n, err := f.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil
		}
		if hw[0] == nil {
			break
		}
		hw[0].Write(b[:n])
		if hw[1] == nil {
			break
		}
		hw[1].Write(b[:n])
		if hw[2] == nil {
			break
		}
		hw[2].Write(b[:n])
	}
	result = append(result, hw[0].Sum(nil))
	result = append(result, hw[1].Sum(nil))
	result = append(result, hw[2].Sum(nil))

	return
}

type HashWriter struct {
	md5Writer hash.Hash
	md5Chan   chan []byte

	sha1Writer hash.Hash
	sha1Chan   chan []byte

	sha256Writer hash.Hash
	sha256Chan   chan []byte

	onceClose *sync.Once

	wg *sync.WaitGroup
}

func doWriteHash(writer hash.Hash, in chan []byte) {
	for {
		select {
		case buf, open := <-in:
			if !open {
				return
			}
			writer.Write(buf)
		}
	}
}

func (this *HashWriter) writeMd5() {
	doWriteHash(this.md5Writer, this.md5Chan)
	fmt.Println("md5 exit")
	this.wg.Done()
}

func (this *HashWriter) writeSha1() {
	doWriteHash(this.sha1Writer, this.sha1Chan)
	fmt.Println("sha1 exit")
	this.wg.Done()
}

func (this *HashWriter) writeSha256() {
	doWriteHash(this.sha256Writer, this.sha256Chan)
	fmt.Println("sha256 exit")
	this.wg.Done()
}

func NewHashWriter(useMd5, useSha1, useSha256 bool) *HashWriter {
	writer := new(HashWriter)
	writer.onceClose = new(sync.Once)
	writer.wg = new(sync.WaitGroup)

	var hashCount int
	chanCount := 30

	if useMd5 {
		hashCount++
		writer.md5Writer = md5.New()
		writer.md5Chan = make(chan []byte, chanCount)
		go writer.writeMd5()
	}
	if useSha1 {
		hashCount++
		writer.sha1Writer = sha1.New()
		writer.sha1Chan = make(chan []byte, chanCount)
		go writer.writeSha1()
	}
	if useSha256 {
		hashCount++
		writer.sha256Writer = sha256.New()
		writer.sha256Chan = make(chan []byte, chanCount)
		go writer.writeSha256()
	}

	writer.wg.Add(hashCount)

	return writer
}

func (this *HashWriter) Write(buf []byte) {
	tmpBuf := make([]byte, len(buf))
	copy(tmpBuf, buf)
	if this.md5Chan != nil {
		this.md5Chan <- tmpBuf
	}
	if this.sha1Chan != nil {
		this.sha1Chan <- tmpBuf
	}
	if this.sha256Chan != nil {
		this.sha256Chan <- tmpBuf
	}
}

func (this *HashWriter) Close() {
	this.onceClose.Do(func() {
		if this.md5Chan != nil {
			close(this.md5Chan)
		}
		if this.sha1Chan != nil {
			close(this.sha1Chan)
		}
		if this.sha256Chan != nil {
			close(this.sha256Chan)
		}
	})
}

func (this *HashWriter) Sum(b []byte) (md5Sum []byte, sha1Sum []byte, sha256Sum []byte) {
	this.Close()

	this.wg.Wait()

	if this.md5Writer != nil {
		md5Sum = this.md5Writer.Sum(b)
	}
	if this.sha1Writer != nil {
		sha1Sum = this.sha1Writer.Sum(b)
	}
	if this.sha256Writer != nil {
		sha256Sum = this.sha256Writer.Sum(b)
	}
	return
}

func GetFileHashs(filename string) ([]byte, []byte, []byte) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil
	}

	hw := NewHashWriter(true, true, true)
	defer hw.Close()

	b := make([]byte, 32*1024)
	for {
		n, err := f.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, nil, nil
		}
		hw.Write(b[:n])
	}

	return hw.Sum(nil)
}

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
