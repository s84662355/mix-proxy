// // 定义无锁队列结构
package Queue


import(
 
 "fmt"
	"unsafe"
	 
		"sync/atomic"
		"errors"
)

type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	status atomic.Bool
	count atomic.Int64
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

// 新建队列，返回一个空队列
func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	qq:=&LKQueue{head: n, tail: n}
	qq.status.Swap(true)
	return  qq
}

func (q *LKQueue) Close( ) {
   q.status.CompareAndSwap(true,false)
}

// 插入，将给定的值v放在队列的尾部
func (q *LKQueue) Enqueue(v interface{}) error {
	if !q.status.Load() {
		return  errors.New("is close")
	}
	n := &node{value: v}
	for {
		if !q.status.Load() {
			return errors.New("is close")
		}

		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) {
			if next == nil {
				if cas(&tail.next, next, n) {
					cas(&q.tail, tail, n) // 排队完成, 尝试将tail移到插入的节点
					q.count.Add(1)
					return nil
				}
			} else { // tail没有指向最后一个节点
				// 将Tail移到下一个节点
				cas(&q.tail, tail, next)
			}
		}
	}
}

// 移除，删除并返回队列头部的值,如果队列为空，则返回nil
func (q *LKQueue) Dequeue() (interface{},error) {

    var err error = nil
	for {
 		if !q.status.Load() {
			err=errors.New("is close")
 		}

		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) {
			if head == tail {
				if next == nil {
					return nil,  err
				}

				cas(&q.tail, tail, next)
			} else {
				// 在CAS之前读取值，否则另一个出队列可能释放下一个节点
				v := next.value
				if cas(&q.head, head, next) {
					q.count.Add(-1)
					if q.count.Load() == -1 {
						fmt.Println("dasdsadsa " ,v)
					}
					return v,err
				}
			}
		}
	}
}

func (q *LKQueue) Count() int64 {
return q.count.Load()
}

func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// CAS算法
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
