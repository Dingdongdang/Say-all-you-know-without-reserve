package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T)  {
	arr := [100]int{}
	for i,_ := range arr{
		arr[i] = rand.Int()%100
		fmt.Printf("%d ",arr[i])
	}
	quickSort(&arr,0,99)
	println()
	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
}
