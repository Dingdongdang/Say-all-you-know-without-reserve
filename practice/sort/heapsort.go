package sort

func heapSort(a *[100]int)  {
	makeMaxHeap(a)
	for i:= len(*a);i>1;i --{
		a[i-1],a[1] = a[1],a[i-1]
		adjust(a,1,i-1)
	}
}

func adjust(a *[100]int,n int ,length int)  {
	for i:=n;i<length; {
		p := 2*i
		if p +1 < length{
			if a[p] < a[p+1]{
				p ++
			}
			if a[i] < a[p] {
				a[i],a[p]=a[p],a[i]
			}else {
				break
			}
		} else if p < length {
			if a[i] < a[p]{
				a[i],a[p] = a[p],a[i]
			}
		}
		i = p
	}
}

func makeMaxHeap(a *[100]int)  {
	for i:=len(*a)/2+1 ;i>0 ;i-- {
		adjust(a, i,len(*a))
	}
}
