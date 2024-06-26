package link

import "testing"

//  lru 如何实现，时间福再度o(1)
// lru 访问会导致网上走到顶部，不能使用 array来实现，否者会有空洞
// 只能使用链表。
// map指向链表?朝招对应的元素
// 如何表示先后访问关系? len(?): 双向链表？才能做到O1
// 还有别的结构么？党项链表。需要删除节点，得找到前一个节点，应该是不太行的

type LruNode struct {
	key   int
	value int
	Pre   *LruNode
	Next  *LruNode
}

type LRUCache struct {
	NodeMap          map[int]*LruNode
	DoubleLinkHeader *LruNode // next for head, prev for tailer
	// DoubleLinkTailer *LruNode
	Len     int
	Capcity int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		NodeMap: make(map[int]*LruNode),
		DoubleLinkHeader: &LruNode{
			key:   0,
			value: 0,
			Pre:   nil,
			Next:  nil,
		},
		Len:     0,
		Capcity: capacity,
	}
	lruCache.DoubleLinkHeader.Pre = lruCache.DoubleLinkHeader
	lruCache.DoubleLinkHeader.Next = lruCache.DoubleLinkHeader
	return lruCache
}

// 注意：生产新节点，删除老节点，然后统一作为新插入处理
func (this *LRUCache) Put(key int, value int) {
	// if new node
	var node *LruNode
	var exist bool
	//  生成新节点，删除老节点
	if node, exist = this.NodeMap[key]; exist {
		node.value = value
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	} else {
		node = &LruNode{
			key:   key,
			value: value,
		}
		if this.Len == this.Capcity {
			delete(this.NodeMap, this.DoubleLinkHeader.Pre.key)
			this.DoubleLinkHeader.Pre = this.DoubleLinkHeader.Pre.Pre
			this.DoubleLinkHeader.Pre.Next = this.DoubleLinkHeader
		} else {
			this.Len++
		}
		this.NodeMap[key] = node
	}
	// 插入新节点
	node.Next = this.DoubleLinkHeader.Next
	node.Pre = this.DoubleLinkHeader
	this.DoubleLinkHeader.Next.Pre = node
	this.DoubleLinkHeader.Next = node
}

func (this *LRUCache) Get(key int) int {
	// just reposition
	if node, exist := this.NodeMap[key]; exist {

		// delete from position
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre

		// add to the new
		node.Next = this.DoubleLinkHeader.Next
		node.Pre = this.DoubleLinkHeader
		this.DoubleLinkHeader.Next.Pre = node
		this.DoubleLinkHeader.Next = node

		return node.value
	} else {
		return -1
	}
}

func Test_LRU(t *testing.T) {

	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)              // 缓存是 {1=1}
	lRUCache.Put(2, 2)              // 缓存是 {1=1, 2=2}
	t.Logf("%d\n", lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)              // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	t.Logf("%d\n", lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)              // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	t.Logf("%d\n", lRUCache.Get(1)) // 返回 -1 (未找到)
	t.Logf("%d\n", lRUCache.Get(3)) // 返回 3
	t.Logf("%d\n", lRUCache.Get(4)) // 返回 4
	/**
	 * Your LRUCache object will be instantiated and called as such:
	 * obj := Constructor(capacity);
	 * param_1 := obj.Get(key);
	 * obj.Put(key,value);
	 */
}
