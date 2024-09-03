package main

type set map[Book]struct{}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}
