package level9

type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		isWord:   false,
		children: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {

		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}

type LRUCache struct {
	Cap        int
	Keys       map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*Node),
	}
}
func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.Keys[key]; ok {
		node.val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node := &Node{
			key:  key,
			val:  value,
			prev: nil,
			next: nil,
		}
		this.Keys[key] = node
		this.Add(node)
	}

	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.key)
		this.Remove(this.tail)
	}
}
func (this *LRUCache) Add(node *Node) {

	node.prev = nil
	node.next = this.head
	if node.next != nil {
		node.next.prev = node
	}
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}
func (this *LRUCache) Remove(node *Node) {

	if node == this.head {

		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return

	}
	if node == this.tail {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev

}

/*
func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		//更新head 为node，所以node的next是旧的head
		this.head = node.next
		//如果node的next不是空的，说明原来有head，需要切断联系
		if node.next != nil {
			//现在新的head是node.next,所以它的prev指针要指向空
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == this.tail {
		//将现在的节点之前一个赋给tail，更新tail的值
		this.tail = node.prev
		//更新后最后一个是node.prev，那么node.prev 的next指针要为nil
		node.prev.next = nil
		//切断node与现在新的tail的联系
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev

}*/
