package stack

type Stack[T any] interface {
	Put(t T)
	PutAll(t ...T)
	Take(n int) []T
	TakeOne() (value T, ok bool)
	Peek() (value T, ok bool)
}

type element[T any] struct {
	v        T
	previous *element[T]
}

type stack[T any] struct {
	l *element[T]
}

func New[T any]() Stack[T] {
	return &stack[T]{}
}

func NewWith[T any](arr ...T) Stack[T] {
	s := stack[T]{}
	s.PutAll(arr...)
	return &s
}

func (s *stack[T]) Put(t T) {
	e := element[T]{
		v:        t,
		previous: nil,
	}

	if s.l == nil {
		s.l = &e
	} else {
		e.previous = s.l
		s.l = &e
	}
}

func (s *stack[T]) PutAll(t ...T) {
	for _, v := range t {
		s.Put(v)
	}
}

func (s *stack[T]) Take(n int) []T {
	temp := make([]T, n)
	for i := 0; i < n; i++ {
		v, ok := s.TakeOne()
		if ok {
			temp[i] = v
		}
	}

	result := make([]T, len(temp))
	for i := range result {
		result[i] = temp[len(temp)-i-1]
	}

	return result
}

func (s *stack[T]) TakeOne() (value T, ok bool) {
	if s.l == nil {
		var empty T
		return empty, false
	} else {
		temp := s.l
		s.l = temp.previous
		return temp.v, true
	}
}

func (s *stack[T]) Peek() (v T, ok bool) {
	if s.l == nil {
		var empty T
		return empty, false
	} else {
		return s.l.v, true
	}
}
