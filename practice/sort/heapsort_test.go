package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHeapSort(t *testing.T)  {
	arr := [100]int{}
	for i,_ := range arr{
		arr[i] = rand.Int()%100
		fmt.Printf("%d ",arr[i])
	}
	heapSort(&arr)
	println()
	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
}