package chat

import "errors"

// Deque 基于环形缓冲区的双端队列（泛型实现）
type Deque[T any] struct {
	buffer []T // 缓冲区
	head   int // 队头指针
	tail   int // 队尾指针
	size   int // 当前队列大小
	cap    int // 缓冲区容量
}

// NewDeque 创建一个新的双端队列
func NewDeque[T any](capacity int) *Deque[T] {
	return &Deque[T]{
		buffer: make([]T, capacity),
		head:   0,
		tail:   0,
		size:   0,
		cap:    capacity,
	}
}

// IsEmpty 判断队列是否为空
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// IsFull 判断队列是否已满
func (d *Deque[T]) IsFull() bool {
	return d.size == d.cap
}

// PushFront 在队头插入元素
func (d *Deque[T]) PushFront(item T) error {
	if d.IsFull() {
		return errors.New("Deque is full")
	}
	d.head = (d.head - 1 + d.cap) % d.cap // 计算新的队头位置
	d.buffer[d.head] = item
	d.size++
	return nil
}

// PushBack 在队尾插入元素
func (d *Deque[T]) PushBack(item T) error {
	if d.IsFull() {
		return errors.New("Deque is full")
	}
	d.buffer[d.tail] = item
	d.tail = (d.tail + 1) % d.cap // 计算新的队尾位置
	d.size++
	return nil
}

// PopFront 从队头删除元素
func (d *Deque[T]) PopFront() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("Deque is empty")
	}
	item := d.buffer[d.head]
	d.head = (d.head + 1) % d.cap // 计算新的队头位置
	d.size--
	return item, nil
}

// PopBack 从队尾删除元素
func (d *Deque[T]) PopBack() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("Deque is empty")
	}
	d.tail = (d.tail - 1 + d.cap) % d.cap // 计算新的队尾位置
	item := d.buffer[d.tail]
	d.size--
	return item, nil
}

// PeekFront 查看队头元素
func (d *Deque[T]) PeekFront() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("Deque is empty")
	}
	return d.buffer[d.head], nil
}

// PeekBack 查看队尾元素
func (d *Deque[T]) PeekBack() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("Deque is empty")
	}
	return d.buffer[(d.tail-1+d.cap)%d.cap], nil
}

// Size 返回队列当前大小
func (d *Deque[T]) Size() int {
	return d.size
}

// Cap 返回队列容量
func (d *Deque[T]) Cap() int {
	return d.cap
}
