/*
 *  @Author : huangzj
 *  @Time : 2020/12/14 16:15
 *  @Description：前缀树\字典树
 */

package trie

import "github.com/ahmetb/go-linq"

type Trie struct {
	Child map[rune]*Trie //存储当前字符（Rune）下面对应的Rune的集合，使用Map可以提高查找效率
	Word  string         //存储字典树上有的字符串信息,如果Word是空字符串("")，则代表没有在该节点找到对应的字符串信息
}

const (
	emptyTag = ""
)

type WordFunc func(sentence []rune, start int, end int) string

func New() *Trie {
	return initTrie()
}

//初始化Trie
func initTrie() *Trie {
	return &Trie{
		Child: make(map[rune]*Trie, 0),
		Word:  emptyTag,
	}
}

//往前缀树中添加对应的字符串(build模式)
func (trie *Trie) Insert(word string) *Trie {
	traverse := trie
	for _, v := range []rune(word) {
		if _, ok := traverse.Child[v]; !ok {
			newTrie := initTrie()
			traverse.Child[v] = newTrie
		}
		traverse = traverse.Child[v]
	}
	traverse.Word = word
	return trie
}

//删除对应的字符串，如果查找不到则直接返回(build模式)
func (trie *Trie) Delete(word string) *Trie {
	traverse := trie
	parent := trie
	var value rune
	for _, v := range []rune(word) {
		if _, ok := traverse.Child[v]; !ok {
			return trie
		}
		parent = traverse
		traverse = traverse.Child[v]
		value = v
	}
	if len(parent.Child) == 0 {
		delete(parent.Child, value)
		return trie
	}
	parent.Child[value].Word = emptyTag
	return trie
}

//查找对应的字符串，找到则返回true，没找到则返回false
func (trie *Trie) Search(word string) bool {
	traverse := trie
	for _, v := range []rune(word) {
		if _, ok := traverse.Child[v]; !ok {
			return false
		}
		traverse = traverse.Child[v]
	}
	if traverse.Word != emptyTag {
		return true
	}
	return false
}

//遍历前缀树存储的所有字符串
func (trie *Trie) Traverse() []string {
	result := make([]string, 0)
	result = traverse(trie, result)
	return result
}

func traverse(trie *Trie, result []string) []string {
	for _, value := range trie.Child {
		if value.Word != emptyTag {
			result = append(result, value.Word)
		}
		//进行子的路径遍历
		if len(value.Child) != 0 {
			result = traverse(value, result)
		}
	}

	return result
}

//查找字符串中匹配的子字符串，进行自定义处理
func (trie *Trie) DealFunc(sentence string, wordFunc WordFunc) string {
	traverse := trie
	for index, v := range []rune(sentence) {
		//从句子的字符开始往下查找,找到最后一个对应的字符(如果uuck和uucku都是敏感词汇的话，会把uucku设置为*,而不只是uuck.)
		if _, ok := traverse.Child[v]; ok {
			traverse = traverse.Child[v]
			continue
		}
		if traverse.Word != emptyTag {
			//传入到wordFunc里面的是字符串的rune数组、查询到的词汇在字符串中的开始位置、查询到的词汇在字符串中的结束位置
			sentence = wordFunc([]rune(sentence), index-len([]rune(traverse.Word)), index)
		}
		//trie从头开始
		traverse = trie
	}

	return sentence
}

//搜索前缀匹配的字符串集合(精准匹配)
func (trie *Trie) SearchTrie(word string) ([]string, bool) {
	traverse := trie
	for _, value := range []rune(word) {
		if _, ok := traverse.Child[value]; !ok {
			return []string{}, false
		}
		traverse = traverse.Child[value]
	}
	return traverse.Traverse(), true
}

//搜索前缀匹配的字符串集合(模糊匹配)
func (trie *Trie) SearchTrieVague(word string) ([]string, bool) {
	var tries []*Trie
	traverse := trie
	for _, value := range []rune(word) {
		if _, ok := traverse.Child[value]; !ok {
			break
		}
		traverse = traverse.Child[value]
		tries = append(tries, traverse)
	}
	result := make([]string, 0)
	for _, item := range tries {
		result = append(result, item.Traverse()...)
	}

	if len(result) == 0 {
		return []string{}, false
	}
	realResult := make([]string, 0)
	linq.From(result).Distinct().Sort(func(i, j interface{}) bool {
		if i.(string) > j.(string) {
			return true
		}
		return false
	}).ToSlice(&realResult) //去重
	return realResult, true
}
