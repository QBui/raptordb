package raptordb

import(
	"testing"
)

// Test IndexType const value
func TestIndexTypeConstValues(t *testing.T) {
   if  BTREE != 0 || HASH != 1 {
	t.Fail()
   }
}
