package kettle

var title = `
    给你一个装满水的 8 升满壶和两个分别是 5 升、3 升的空壶，请想个优雅的办法，
使得其中一个水壶恰好装 4 升水，每一步的操作只能是倒空或倒满。
`

type node struct{
	arr [3]int
	parent *node
}

func kettle()  {
	vol := [3]int{8,5,3}
	doing := make([]node,0)
	done := make(map[int]string,0)
	root := node{[3]int{8,0,0},nil}
	doing = append(doing, root)
	for len(doing) != 0 {
		cur := doing[0]
		for i:=0;i< len(vol);i++{
			if cur.arr[i] != 0 {
				for j:=1;j<len(vol) ;j++  {
					next := (i+j)%len(vol)
					diff := vol[next] - cur.arr[next]
					if diff > cur.arr[i] {
						diff = cur.arr[i]
					}
					if diff > 0 {
						child := node{[...]int{cur.arr[0],cur.arr[1],cur.arr[2]},&cur}
						child.arr[i] -= diff
						child.arr[next] += diff
						if child.arr[0] == 4 || child.arr[1] == 4 || child.arr[2] == 4{
							p(&child)
							return
						}
						sum := child.arr[0]*100 + child.arr[1] *10 + child.arr[2]
						if done[sum] == ""{
							doing = append(doing, child)
						}
					}
				}
			}
		}
		doing = doing[1:]
		done[cur.arr[0]*100+cur.arr[1]*10 + cur.arr[2]] = "done"
	}
}

func p(node *node)  {
	if node.parent != nil {
		p(node.parent)
	}
	println(node.arr[0],node.arr[1],node.arr[2])
}
