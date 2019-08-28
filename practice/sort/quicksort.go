package sort

func quickSort(a *[100]int,low int ,high int)  {
	mid := a[low]
	flow := low
	fhigh := high
	lowTurn := true
	for flow < fhigh{
		if lowTurn {
			for flow < fhigh && mid < a[fhigh] {
				fhigh --
			}
			if flow < fhigh {
				a[flow] = a[fhigh]
				flow ++
				lowTurn = false
			}
		} else {
			for flow < fhigh && a[flow] < mid {
				flow ++
			}
			if flow < fhigh {
				a[fhigh] = a[flow]
				fhigh --
				lowTurn = true
			}
		}
	}
	a[flow] = mid
	if low < flow {
		quickSort(a,low,flow-1)
	}
	if flow < high {
		quickSort(a,flow+1,high)
	}
}