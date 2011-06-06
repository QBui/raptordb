package raptordb

import(
	"testing"
	"time"
	"fmt"
        crand "crypto/rand"
	"io"
	"log"
//	"os"
        "rand"
//	"./storage"
)

// Test IndexType const value
func TestIndexTypeConstValues(t *testing.T) {
   if  BTREE != 0 || HASH != 1 {
	t.Fail()
   }
}

// Test writing 1 mil records
func TestStorageFileForWriting1MilRecords(t *testing.T) {
   ns := -time.Nanoseconds()

//   file := io.File.Create("one_million.dat")
//   storage := storage.StorageFile.Create(file, 16)
//   storage.SkipDateTime = true;

   

   ns += time.Nanoseconds()
   fmt.Printf("Duration: %.2f seconds\n", float64(ns)/1e9)
}

// Test timer
func TestTimeModule(t *testing.T) {
   ns := -time.Nanoseconds()
  
   time.Sleep(1000000000) // 1 second

   ns += time.Nanoseconds()
   fmt.Printf("Duration: %.2f seconds\n", float64(ns)/1e9)
}

// Test bytes module
func TestBytesModule(t *testing.T) {
   
}

// Test rand module
func TestRandModule(t *testing.T) {
   rand.Seed(1234567890)
   var p int
   for i := 0; i < 10; i ++ {
      p = rand.Int()
      println("Rand.Int() returns ", p)
   }
}

func TestStringToByteArray(t *testing.T){
   s := "Hello World"
   byteArray := []byte(s)
   fmt.Printf("%s\n",byteArray)
}

func uuid() string {
	b := make([]byte, 16)
	_, err := io.ReadFull(crand.Reader, b)
	if err != nil {
		log.Fatal(err)
	}
	b[6] = (b[6] & 0x0F) | 0x40
	b[8] = (b[8] &^ 0x40) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}


func TestUuid(t *testing.T) {
   for i := 0; i < 10; i++ {
      fmt.Println(uuid())
   }
   
}
