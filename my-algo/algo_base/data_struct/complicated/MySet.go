package complicated

type MySet struct {
	ImmutableMap map[interface{}]bool
}

func (set *MySet) Construct() {
	set.ImmutableMap = make(map[interface{}]bool)
}

func (set *MySet) Insert(value interface{}) {
	set.ImmutableMap[value] = true
}

func (set *MySet) Exist(value interface{}) bool {
	return set.ImmutableMap[value] == true
}

func (set *MySet) Erase(value interface{}) {
	if set.Exist(value) {
		delete(set.ImmutableMap, value)
	}
}
