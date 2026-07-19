type Node struct {
        val  int
        next *Node
        prev *Node
}

func (n *Node) display() {
        fmt.Printf("Node(%d)",n.val)
}

type LList struct {
        tail *Node
        head *Node
        size int
}

func (ll *LList) add(val int) {
        if ll.head == nil {
                ll.head = &Node{val: val}
                ll.tail = ll.head
                ll.size = 1
                return
        }

        next := ll.head
        ll.head = &Node{val: val, next: next}
        next.prev = ll.head
        ll.size++
}

func (ll *LList) pop() (int, bool) {
        if ll.tail == nil {
                return 0, false
        }

        if ll.tail.prev == nil {
                r := ll.tail.val
                ll.head = nil
                ll.tail = nil
                ll.size = 0
                return r, true
        }

        r := ll.tail.val
        ll.tail = ll.tail.prev
        ll.tail.next = nil
        ll.size--

        return r, true
}

func (ll *LList) updateHead(val int) {


        current := ll.head
        if current == nil {
                return
        }
        var node *Node
        for {
                if current.val == val {
                        node = current
                        break
                }
                current = current.next
        }

        // node is head
        if node == ll.head {
                return
        }

        // node is tail
        if node == ll.tail {
                // repoint tail
                node.prev.next = nil
                ll.tail = node.prev

                // add to head
                ll.head.prev = node
                node.next = ll.head
                node.prev = nil
                ll.head = node
                return
        }

        // node in the middle
        // remove from chain, linking the previous to the next
        node.prev.next = node.next
        node.next.prev = node.prev

        //  update head
        ll.head.prev = node
        node.prev = nil
        node.next = ll.head
        ll.head = node

}

func (ll *LList) display() {
        current := ll.head
        for {
                current.display()
                current = current.next
                if current == nil {
                        break
                } 
                fmt.Printf(" -> ")

        }

        fmt.Println()
}

type LRUCache struct {
        cache map[int]int
        llist LList
        cap   int
}

func Constructor(capacity int) LRUCache {
        return LRUCache{cap: capacity, cache: make(map[int]int)}

}

func (this *LRUCache) Get(key int) int {
        if v, ok := this.cache[key]; ok {
                this.llist.updateHead(key)
                return v
        }
        return -1
}

func (this *LRUCache) Put(key int, value int) {

        if _, ok := this.cache[key]; ok {
                this.cache[key] = value
                this.llist.updateHead(key)
                return
        }


        this.cache[key] = value
        this.llist.add(key)
        if this.llist.size > this.cap {
                if rkey, ok := this.llist.pop(); ok {
                        delete(this.cache, rkey)
                }
        }

}
