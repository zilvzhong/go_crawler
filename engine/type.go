package engine

import "sync"

type Request struct {
	Url string
	Region string
	ParserFunc func([]byte, string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items string
	//Items []map[string]string
}


func NilParser([]byte, string) ParseResult {
	return ParseResult{}
}



type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}


// New: 返回一个Set实例
func New() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

// Add: 增加一个元素
func (s *Set)Add(item interface{})  {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// Remove: 移除一个元素
func (s *Set)Remove(item interface{})  {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// Has: 是否存在指定的元素
func (s *Set)Has(item interface{}) bool {
	// 允许读
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// List: 获取Map转化成的list
func (s *Set)List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var l []interface{}
	for value := range s.m {
		l = append(l, value)
	}
	return l
}

// Len: 返回元素个数
func (s *Set)Len() int {
	return len(s.List())
}

// Clear: 清除Set
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

// Empty: Set是否是空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}
