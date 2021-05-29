package amazon

type LRUCache struct {
	capacity int
	head   *Node
	tail   *Node
	length int
	keys   []*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head:   nil,
		tail:   nil,
		length: 0,
		keys:   make([]*Node, MaxSize),
	}
}

func (this *LRUCache) Get(key int) int {
	found := this.keys[key]
	if found == nil {
		return -1
	}

	this.moveAhead(found)

	return found.value
}

func (this *LRUCache) Put(key int, value int) {
	cur := this.keys[key]

	if cur != nil {
		cur.value = value
	} else {

		if this.length == this.capacity {
			if this.head == this.tail {
				this.keys[this.head.key] = nil
				this.head = nil
				this.tail = nil
			} else {
				this.keys[this.tail.key] = nil
				this.tail.prev.next = nil
				this.tail = this.tail.prev
			}
		} else {
			this.length++
		}

		cur = &Node{
			key:   key,
			value: value,
		}
		this.keys[key] = cur
	}
	this.moveAhead(cur)
}

func (this *LRUCache) moveAhead(node *Node) {
	if this.head == nil && this.tail == nil {
		this.head = node
		this.tail = node
	} else {
		if this.head == node {
			return
		}
		if this.tail == node {
			this.tail.prev.next = nil
			this.tail = this.tail.prev
		}
		// adjust tail and move found node ahead
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

const MaxSize = 3001

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
