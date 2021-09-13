package complicated

import "fmt"

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

func (set *MySet) Len() int {
	return len(set.ImmutableMap)
}
func (set *MySet) String() string {
	if set.Len() == 0 {
		return "nil"
	}
	setStr := "["
	for ele, _ := range set.ImmutableMap {
		setStr += fmt.Sprintf("%v,", ele)
	}
	setStr = setStr[:len(setStr)-1]
	setStr += "]"
	return setStr
}
func (set *MySet) Slice() []interface{} {
	eleArr := make([]interface{}, 0)
	for _, ele := range set.ImmutableMap {
		eleArr = append(eleArr, ele)
	}
	return eleArr
}

// Union 交集
func (set *MySet) Union(cSet *MySet) *MySet {
	unionSet := &MySet{}
	unionSet.Construct()
	for ele, _ := range set.ImmutableMap {
		if cSet.Exist(ele) {
			unionSet.Insert(ele)
		}
	}
	return unionSet
}

// Unite 并集
func (set *MySet) Unite(cSet *MySet) *MySet {
	uniteSet := &MySet{}
	uniteSet.Construct()
	for ele, _ := range set.ImmutableMap {
		uniteSet.Insert(ele)
	}
	for ele, _ := range cSet.ImmutableMap {
		uniteSet.Insert(ele)
	}
	return uniteSet
}

// Diff 差集
func (set *MySet) Diff(cSet *MySet) *MySet {
	diffSet, unionSet := &MySet{}, set.Union(cSet)
	diffSet.Construct()
	if unionSet != nil && unionSet.Len() != 0 {
		for ele, _ := range set.ImmutableMap {
			if !cSet.Exist(ele) {
				diffSet.Insert(ele)
			}
		}
	}
	return diffSet
}
