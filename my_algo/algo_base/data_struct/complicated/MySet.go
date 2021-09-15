package complicated

import (
	"fmt"
	"reflect"
	"strconv"
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

func TypeForcedTransfer(value interface{}) string {
	typeName, ValueStr := reflect.TypeOf(value).Name(), ""
	switch typeName {
	case "int":
		ValueStr = strconv.Itoa(value.(int))
		break
	case "bool":
		ValueStr = strconv.Itoa(value.(int))
		break
	case "float":
		ValueStr = strconv.FormatFloat(value.(float64), 8, 0, 64)
		break
	case "int32":
		ValueStr = strconv.FormatInt(int64(value.(int32)-49), 10)
	default:
		ValueStr = value.(string)
	}
	return ValueStr
}

// Multiple 乘积
func (set *MySet) Product(cSet *MySet) *MySet {
	if set == nil || cSet == nil || set.Len() < 1 || cSet.Len() < 1 {
		return nil
	}
	multipleSet := &MySet{}
	multipleSet.Construct()
	for ele1, _ := range set.ImmutableMap {
		for ele2, _ := range cSet.ImmutableMap {
			tEle1, tEle2 := TypeForcedTransfer(ele1), TypeForcedTransfer(ele2)
			multipleSet.Insert(tEle1 + tEle2)
		}
	}
	return multipleSet
}
func (set *MySet) Power(n int) *MySet {
	if n < 1 || set == nil || set.Len() < 1 {
		return nil
	}
	powerSet := &MySet{}
	powerSet.Construct()
	for i := 1; i <= n; i++ {
		powerSet = powerSet.Product(set)
	}
	return powerSet
}

// PositiveClosure 因为没有边界所以需要约束N
func (set *MySet) PositiveClosure(n int) *MySet {
	closureSet := &MySet{}
	closureSet.Construct()
	for i := 1; i <= n; i++ {
		closureSet = closureSet.Unite(set.Power(i))
	}
	return closureSet
}

// KleeneClosure 在PositiveClosure加一个空边即可
func (set *MySet) KleeneClosure(n int) *MySet {
	kleeneClosure := set.PositiveClosure(n)
	kleeneClosure.Insert(Epsilon)
	return kleeneClosure
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
