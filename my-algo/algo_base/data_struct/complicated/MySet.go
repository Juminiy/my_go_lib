package complicated

import (
	"fmt"
	"reflect"
)

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

func (set *MySet) DeepExist(value interface{}) bool {
	for ele, _ := range set.ImmutableMap {
		if reflect.DeepEqual(ele, value) {
			return true
		}
	}
	return false
}

func (set *MySet) Erase(value interface{}) {
	if set.Exist(value) {
		delete(set.ImmutableMap, value)
	}
}

func (set *MySet) Len() int {
	if set == nil || set.ImmutableMap == nil {
		return 0
	}
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

func SliceToSet(dataSlice []interface{}) *MySet {
	mySet := &MySet{}
	mySet.Construct()
	for _, ele := range dataSlice {
		mySet.Insert(ele)
	}
	return mySet
}
func (set *MySet) Slice() []interface{} {
	eleArr := make([]interface{}, 0)
	for ele, _ := range set.ImmutableMap {
		eleArr = append(eleArr, ele)
	}
	return eleArr
}
func (set *MySet) SliceBatchInsert(slice []interface{}) {
	if slice == nil || len(slice) == 0 {
		return
	}
	for _, ele := range slice {
		set.Insert(ele)
	}
}
func (set *MySet) SortSetToSlice() []interface{} {
	setIntSlc := make([]interface{}, 0)
	for ele, _ := range set.ImmutableMap {
		setIntSlc = append(setIntSlc, ele)
	}
	return setIntSlc
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
	if set != nil && set.Len() > 0 {
		for ele, _ := range set.ImmutableMap {
			uniteSet.Insert(ele)
		}
	}
	if cSet != nil && cSet.Len() > 0 {
		for ele, _ := range cSet.ImmutableMap {
			uniteSet.Insert(ele)
		}
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
func (set *MySet) Equal(cSet *MySet) bool {
	if set == nil && cSet == nil {
		return true
	}
	if set != nil && set.Len() > 0 && cSet != nil && cSet.Len() > 0 {
		unionSet := set.Union(cSet)
		if unionSet != nil &&
			unionSet.Len() != 0 &&
			unionSet.Len() == set.Len() &&
			unionSet.Len() == cSet.Len() {
			return true
		} else {
			return true
		}
	}
	return false
}
