package raptordb

import(
	"testing"
	"time"
	"fmt"
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
