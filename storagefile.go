package raptordb

import (
   "os"
   "bufio"
   "bytes"
   "sync"
)

type SyncedBuffer struct {
   lock sync.Mutex
   buffer bytes.Buffer
}

type StorageFile struct {
   Filename string
   MaxKeyLen int
   file *os.File
   writer *bufio.Writer
}

func NewStorageFile(name string, maxKeyLen int) (store *StorageFile){
   store = &StorageFile{Filename: name, MaxKeyLen: maxKeyLen}
   store.initFileWriter()
   
   return;
}

func (s *StorageFile) initFileWriter() {
   if file, err := os.OpenFile(s.Filename, os.O_CREATE, 0777); err == nil {
      s.file = file
      s.writer = bufio.NewWriter(file)
      println("Success allocating a new Writer")
   }
}

func (s *StorageFile) Close() {
   if s.file != nil {
      s.file.Close()
      println("storagefile: Call Close()")
   }
}

func (s *StorageFile) Flush() {
   s.writer.Flush()
}

func (s *StorageFile) Write(b []byte) {
   s.writer.Write(b)
}
