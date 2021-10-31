package complicated

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAvlTree_Insert(t *testing.T) {
	avl := &AvlTree{NodeAmount: 0}
	for i := 0; i <= 20; i++ {
		avl.Insert(rand.Int())
	}
	seq := avl.DfsAvl()
	fmt.Println("\n", seq)
	seq = avl.BfsAvl()
	fmt.Println("\n", seq)
}
