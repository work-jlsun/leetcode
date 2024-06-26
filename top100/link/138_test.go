package link

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 方案： 通过 slice、以及map来在扫描过程中暂存关系
//  next 订正，直接进入到 slice一并更新
// random 怎么更新：存储random对应的index,在push的时候保存每一个 random丢应在
// index中的值

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var nodeSlice []*Node
	var nodeIndexMap = make(map[*Node]int, 0)
	for i := 0; head != nil; i++ {
		// new node
		newNode := &Node{
			Val:    head.Val,
			Next:   head.Next,
			Random: head.Random,
		}
		nodeSlice = append(nodeSlice, newNode)
		nodeIndexMap[head] = i
		head = head.Next
	}

	for i := 0; i < len(nodeSlice); i++ {
		if i != len(nodeSlice)-1 {
			nodeSlice[i].Next = nodeSlice[i+1]
		}
		if nodeSlice[i].Random != nil {
			index, _ := nodeIndexMap[nodeSlice[i].Random]
			nodeSlice[i].Random = nodeSlice[index]
		}
	}
	return nodeSlice[0]
}
