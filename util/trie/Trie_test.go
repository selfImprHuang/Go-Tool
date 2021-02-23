/*
 *  @Author : huangzj
 *  @Time : 2020/12/14 16:15
 *  @Description：
 */

package trie

import (
	"fmt"
	"testing"
)

//模糊查询前缀匹配的所有字符串
func TestSearchTrieVague(t *testing.T) {
	fmt.Println()
	fmt.Println("模糊查询 查找前缀匹配的所有字符串，输入\"新年快乐\" ")
	trie := New()
	trie.Insert("abcuuid").Insert("abcuued").Insert("abcMxns").Insert("abssde").Insert("avddfa")
	trie.Insert("我爱中国").Insert("我爱祖国").Insert("我爱学习").Insert("我是谁").Insert("我在哪里")
	trie.Insert("我是你").Insert("我").Insert("我喜欢学习").Insert("我在厦门").Insert("我是一个程序员")
	trie.Insert("新年快乐，你好呀").Insert("新年新气象").Insert("新的一年新的开始")

	fmt.Println()
	fmt.Println()
	result3, _ := trie.SearchTrieVague("新年快乐")
	for _, r := range result3 {
		fmt.Println(r)
	}
}

//查找前缀匹配的所有字符串
func TestSearchTrie(t *testing.T) {
	fmt.Println()
	fmt.Println("查找前缀匹配的所有字符串")
	trie := New()
	trie.Insert("abcuuid").Insert("abcuued").Insert("abcMxns").Insert("abssde").Insert("avddfa")
	trie.Insert("我爱中国").Insert("我爱祖国").Insert("我爱学习").Insert("我是谁").Insert("我在哪里")

	fmt.Println()
	fmt.Println()
	result, _ := trie.SearchTrie("abc")
	for _, r := range result {
		fmt.Println(r)
	}

	fmt.Println()
	fmt.Println()
	result1, _ := trie.SearchTrie("我爱")
	for _, r := range result1 {
		fmt.Println(r)
	}

	fmt.Println()
	fmt.Println()
	result2, _ := trie.SearchTrie("我爱她")
	for _, r := range result2 {
		fmt.Println(r)
	}

}

//敏感词汇屏蔽测试
func TestSensitiveWordsTest(t *testing.T) {
	trie := New()
	trie.Insert("粗话").Insert("fuck").Insert("fuckk").Insert("脏话")
	s := "我是脏话，我是粗话，fuckk,fuck,fffuccck"
	word := trie.DealFunc(s, func(sentence []rune, start int, end int) string {
		var replace string
		for i := 0; i < end-start; i++ {
			replace = replace + "*"
		}
		runes := append(sentence[0:start], []rune(replace)...)
		runes = append(runes, sentence[end:]...)
		return string(runes)
	})
	fmt.Println(fmt.Sprintf("原来的字符串： \"%v\"", s))
	fmt.Println(fmt.Sprintf("转换后的字符串： \"%v\"", word))
}

//常规功能测试
func TestCommonTest(t *testing.T) {
	trie := New()
	trie.Insert("ABC").Insert("AB").Insert("ABE").Insert("ABEX").Insert("XYZ").Insert("你").Insert("你好").Insert("你是谁")
	result := trie.Traverse()
	for _, r := range result {
		fmt.Println(r)
	}

	fmt.Println()
	fmt.Println("进行查找测试")
	fmt.Println(trie.Search("Abc"))
	fmt.Println(trie.Search("ABC"))
	fmt.Println(trie.Search("ABEX"))
	fmt.Println(trie.Search("ABEx"))
	fmt.Println(trie.Search("你好"))
	fmt.Println(trie.Search("你好吗"))

	fmt.Println("进行删除测试")
	trie.Delete("你").Delete("ABC")
	fmt.Println(trie.Search("你"))
	fmt.Println(trie.Search("你好"))
	fmt.Println(trie.Search("你是谁"))
	fmt.Println(trie.Search("AB"))
	fmt.Println(trie.Search("ABE"))
	fmt.Println(trie.Search("ABC"))

	fmt.Println()
	fmt.Println("再次遍历所有的字典")
	result = trie.Traverse()
	for _, r := range result {
		fmt.Println(r)
	}
}
