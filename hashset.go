package main

type hashset struct {
	data map[interface{}]struct{}
}

func newhashset() *hashset {
	var hs hashset
	hs.data = make(map[interface{}]struct{})
	return &hs
}

func (hs *hashset) insert(v interface{}) {
	hs.data[v] = struct{}{}
}

func (hs *hashset) has(v interface{}) bool {
	_, c := hs.data[v]
	return c
}
