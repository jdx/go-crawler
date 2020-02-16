package crawler

// Set is a set data structure
type set map[interface{}]bool

// Add a key to the set
func (s set) Add(key interface{}) {
	s[key] = true
}

// Has returns true if key in set
func (s set) Has(key interface{}) bool {
	_, has := s[key]
	return has
}
