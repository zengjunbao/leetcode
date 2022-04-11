package main

func main() {

}

type s1 struct {
	id uint
}

type l1 []*s1

type any interface {
	getKey() uint
}

func (s *s1) getKey() uint {
	return 0
}

func getFunc(list l1) bool{
	mp := make(map[uint]bool)
	for _, v := range list {
		mp[v.getKey()] = true
	}
	
	return false
}