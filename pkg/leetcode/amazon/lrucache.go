package amazon

type LRUCache struct {
	capacity int
	head *Node
	tail *Node
	length int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head: nil,
		tail: nil,
		length: 0,
	}
}

func (this *LRUCache) Get(key int) int {
	cur := this.head
	for cur != nil {
		if cur.key == key {
			ret := cur.value

			if this.head == cur {
				return cur.value
			} else if this.tail == cur {
				this.tail = cur.prev
				this.tail.next = nil
			}

			this.head.prev = cur
			cur.next = this.head
			this.head = cur

			return ret
		}
		cur = cur.next
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	cur := this.head
	for cur != nil {
		if cur.key == key {
			cur.value = value
			if this.head == cur {
				return
			}
			if this.tail == cur {
				this.tail = this.tail.prev
				this.tail.next = nil
				cur.prev = nil
			}
			this.head.prev = cur
			cur.next = this.head
			this.head = cur
			return
		}
		cur = cur.next
	}
	cur = &Node{
		key: key,
		value: value,
	}
	if this.length == this.capacity {
		if this.head == this.tail {
			this.head = nil
			this.tail = nil
		} else {
			prev := this.tail.prev
			prev.next = nil
			this.tail = prev
		}
	} else {
		this.length++
	}
	if this.head == nil && this.tail == nil {
		this.head = cur
		this.tail = cur
	} else {
		this.head.prev = cur
		cur.next = this.head
		this.head = cur
	}
}

type Node struct {
	key int
	value int
	next *Node
	prev *Node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */