/*
 *  @Author : huangzj
 *  @Time : 2020/6/15 16:16
 *  @Description：
 */

package set

import "sync"

type Set struct {
	setMap sync.Map //用来存储元素的map
}

func (set *Set) Add(i interface{}) {
	set.setMap.Store(i, nil)
}

func (set *Set) Remove(i interface{}) {
	set.setMap.Delete(i)
}

func (set *Set) Contains(i interface{}) bool {
	_, has := set.setMap.Load(i)
	return has
}

func (set *Set) IsEmpty() bool {
	length := 0
	set.setMap.Range(func(key, value interface{}) bool {
		length = length + 1
		return true
	})
	return length == 0
}

func (set *Set) GetAllSet() []interface{} {
	list := make([]interface{}, 0)
	set.setMap.Range(func(key, value interface{}) bool {
		list = append(list, key)
		return true
	})
	return list
}
