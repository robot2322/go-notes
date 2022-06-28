package observer

import "fmt"

// 观察目标类实现

// Subject 观察目标类
type Subject struct {
	observers []Observer
	context string
}

// NewSubject 实例化观察目标
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// Attach 将观察者绑定到观察目标
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// notify 通知所有观察者
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// UpdateContext更新观察目标内容
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

// 观察者类实现

// Observer 观察者抽象接口
type Observer interface {
	Update(*Subject)
}

// Reader 观察者类
type Reader struct {
	name string
}

// NewReader 实例化观察者类
func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

// Update 观察者接受观察目标的更新内容
func (r *Reader) Update(s *Subject){
	fmt.Printf("%s receive %s\n", r.name, s.context)
}