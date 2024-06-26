package stack

// 方案：双堆栈方法
type MinStack struct {
	dataSlice    []int
	minDataSlice []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.dataSlice = append(this.dataSlice, val)
	if len(this.minDataSlice) >= 1 {
		lastVal := this.minDataSlice[len(this.minDataSlice)-1]
		var minVal int
		if lastVal > val {
			minVal = val
		} else {
			minVal = lastVal
		}
		this.minDataSlice = append(this.minDataSlice, minVal)
	} else {
		this.minDataSlice = append(this.minDataSlice, val)
	}
}

func (this *MinStack) Pop() {
	//  注意，slice切割，第二各是 从开始元素开始切割的总长度
	this.dataSlice = this.dataSlice[:len(this.dataSlice)-1]
	this.minDataSlice = this.minDataSlice[:len(this.minDataSlice)-1]
}

func (this *MinStack) Top() int {
	return this.dataSlice[len(this.dataSlice)-1]
}

func (this *MinStack) GetMin() int {
	return this.minDataSlice[len(this.minDataSlice)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
