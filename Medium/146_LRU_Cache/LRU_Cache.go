package question146

/*
	146. LRU Cache
	Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

	get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
	put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

	The cache is initialized with a positive capacity.
	Follow up:
	Could you do both operations in O(1) time complexity?

	146. LRU缓存机制
	运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

	获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
	写入数据 put(key, value) - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

 	进阶:
	你是否可以在 O(1) 时间复杂度内完成这两种操作？
 */

/*
	思路: 这很明显得用个map存key value 以及使用次数和上限 =。= 干 超时了
	题解: 采用hashmap和双向链表，超时的原因应该每次新添一个key-value就释放再加一次 目的是让map有序
		 map存链表中的位置，链表存值。。这样删值的时候替换一下链表就好
		 注意下Java的linkedhashmap
*/

//https://studygolang.com/articles/27121?fr=sidebar map嵌套问题

type LinkedListNode struct {
	value int
	key int
	prev *LinkedListNode
	next *LinkedListNode
}

type LRUCache struct {
	capacity int
	length int //记一下当前链表长度
	cache map[int]*LinkedListNode
	head *LinkedListNode //头节点最近使用的
	tail *LinkedListNode //尾节点最久没使用的
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		length:   0,
		cache:    make(map[int]*LinkedListNode, capacity),
		head:     nil,
		tail:     nil,
	}
}

func deleteNode(delnode *LinkedListNode, head *LinkedListNode, tail *LinkedListNode) (*LinkedListNode, *LinkedListNode, *LinkedListNode) {

	if delnode.prev != nil {
		delnode.prev.next = delnode.next
	}

	if delnode.next != nil {
		delnode.next.prev = delnode.prev
	}

	if head == delnode {
		head = delnode.next
	}

	if tail == delnode {
		tail = delnode.prev
	}

	return delnode, head, tail
}

func addHead(addnode *LinkedListNode, head *LinkedListNode, tail *LinkedListNode) (*LinkedListNode, *LinkedListNode) {

	addnode.prev = nil
	addnode.prev = nil

	if head == nil || tail == nil {
		head = addnode
		tail = addnode
	} else {
		head.prev = addnode
		addnode.next = head
		head = addnode
	}

	return head, tail
}

func (this *LRUCache) Get(key int) int {

	//cache里没有返回-1
	if _, ok := (*this).cache[key]; !ok {
		return -1
	} else {
		//找到就是最近使用的，挪到头节点，即删除这个点再挪到开头
		curnode := (*this).cache[key]
		curnode, (*this).head, (*this).tail = deleteNode(curnode, (*this).head, (*this).tail)
		(*this).head, (*this).tail = addHead(curnode, (*this).head, (*this).tail)

		return curnode.value
	}
}


func (this *LRUCache) Put(key int, value int)  {

	//有位置直接放
	if _, ok := (*this).cache[key]; ok {
		//cache里有这个值就更新一下，像get一样挪到最前面 先判断这个！！！
		curnode := (*this).cache[key]
		curnode, (*this).head, (*this).tail = deleteNode(curnode, (*this).head, (*this).tail)
		(*this).head, (*this).tail = addHead(curnode, (*this).head, (*this).tail)

		curnode.value = value
	} else if (*this).length < (*this).capacity {
		curnode := LinkedListNode{
			value: value,
			key: key,
			prev: nil,
			next: nil,
		}
		(*this).cache[key] = &curnode
		(*this).head, (*this).tail = addHead(&curnode, (*this).head, (*this).tail)
		(*this).length += 1
	} else {
		//删除尾节点，替换值
		var tailnode *LinkedListNode
		tailnode, (*this).head, (*this).tail = deleteNode((*this).tail, (*this).head, (*this).tail)
		delete((*this).cache, tailnode.key)
		curnode := LinkedListNode{
			value: value,
			key: key,
			prev: nil,
			next: nil,
		}
		(*this).cache[key] = &curnode
		(*this).head, (*this).tail = addHead(&curnode, (*this).head, (*this).tail)
	}
}


//type LRUCache struct {
//	maxcapacity int
//	keyval map[int][]int //int[0] value int[1] count
//}
//
//func addOpt(cachemap *LRUCache) {
//
//	for ckey := range (*cachemap).keyval {
//		(*cachemap).keyval[ckey][1] += 1
//	}
//}
//
//func Constructor(capacity int) LRUCache {
//
//	ret := LRUCache{
//		maxcapacity: capacity,
//		keyval: make(map[int][]int, capacity),
//	}
//
//	return ret
//}
//
//
//func (this *LRUCache) Get(key int) int {
//
//	if _, ok := (*this).keyval[key]; !ok {
//		return -1
//	} else {
//		addOpt(this)
//		(*this).keyval[key][1] = 0
//		return (*this).keyval[key][0]
//	}
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//
//	addOpt(this)
//	if _, ok := (*this).keyval[key]; ok {
//		(*this).keyval[key][0] = value
//		(*this).keyval[key][1] = 0
//		return
//	} else if len((*this).keyval) >= (*this).maxcapacity {
//		var delkey int
//		delkeycount := 0
//		for ckey := range (*this).keyval {
//			if (*this).keyval[ckey][1] > delkeycount {
//				delkey = ckey
//				delkeycount = (*this).keyval[ckey][1]
//			}
//		}
//		delete((*this).keyval, delkey)
//	}
//	(*this).keyval[key] = make([]int, 2)
//	(*this).keyval[key][0] = value
//	(*this).keyval[key][1] = 0
//
//}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */