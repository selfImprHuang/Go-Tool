/*
 *  @Author : huangzj
 *  @Time : 2020/11/12 10:52
 *  @Description：这边是linq包的使用操作，为之后的使用做一些示例
 *  包引用命令： go get github.com/ahmetb/go-linq
 *  示例参考地址：https://godoc.org/github.com/ahmetb/go-linq#
 *  在Idea中对应方法名上使用 Shift+F1 组合键操作可以直接跳转到方法对应的地址
 */

package linqUse

import (
	"Go-Tool/util/timed"
	"fmt"
	"github.com/ahmetb/go-linq"
	"strconv"
	"testing"
	"time"
)

func Test(t *testing.T) {

	Aggregate() //自定义聚合操作

	All() //判断数组的所有元素是否都满足条件

	Where() //根据条件查询对应的元素

	AnyWith() //判断数组是否有任意个元素满足条件

	Calculate() //这个方法只能计算数值类型

	Combine() //组合两个数组的两种方式

	Contains() //包含

	FirstAndLast() //第一个和最后一个

	Distinct() //返回数组中唯一的元素

	Except() //返回第一个序列没有出现在第二个序列中的成员

	GroupBy() //按照对应规则进行分组

	Intersect() //产生两个数组的交集

	OrderBy() //排序

	Skip() //略过指定条件的元素

	Count() //计算数组总长度

	Result() //对结果集进行操作

	Take() //获取对应数量的元素，一般多用在排序后前几个的元素

	Select() //筛选结构体中对应的元素

	SelectMany() //对多维数组进行合并或者其他处理 对结构体中的数组进行处理(与结构体结合操作)

	Join() //对两个有关联的数据进行处理

}

func Where() {
	bookList := MakeBook()
	query := linq.From(bookList)
	fmt.Println()
	fmt.Println()
	query.WhereT(func(book Book) bool {
		return book.Money > float64(70)
	}).ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", book.Author, book.Name))
	})
}

func Take() {
	bookList := MakeBook()
	query := linq.From(bookList)
	query.OrderByDescendingT(func(book Book) string {
		return book.Author
	}).Take(3).ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", book.Author, book.Name))
	})

	fmt.Println()
	fmt.Println()

	//超出对应数量的元素会不会报错（输出所有的元素）
	query.OrderByDescendingT(func(book Book) string {
		return book.Author
	}).Take(100).ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", book.Author, book.Name))
	})

	fmt.Println()
	fmt.Println()
	//负数会发生什么情况(直接返回，不做操作)
	query.OrderByDescendingT(func(book Book) string {
		return book.Author
	}).Take(-2).ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", book.Author, book.Name))
	})

	//TakeWhile系列的方法，虽然能够根据条件进行查询，但是有一个弊端就是一旦查询到的结果和条件不符合则直接返回，也就是说后面满足条件的元素不会再被做筛选
	//那么比如这边的这个大于的操作，就需要对整个数组进行排序之后然后进行筛选，如果不希望排序，则用原生的for循环可能会更好
	fmt.Println()
	fmt.Println()
	query.TakeWhileT(func(book Book) bool {
		return book.Money > float64(70)
	}).ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", book.Author, book.Name))
	})
}

func Result() {
	bookList := MakeBook()
	query := linq.From(bookList)

	bookMap := make(map[string]string, 0)
	query.SelectT(func(book Book) linq.KeyValue {
		return linq.KeyValue{
			Key:   book.Author,
			Value: book.Name,
		}
	}).ToMap(&bookMap)

	for key, value := range bookMap {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", key, value))
	}

	fmt.Println()
	fmt.Println()

	//如果key相同的话，前面的元素会被覆盖掉（只会保留最后一个）
	bookList1 := MakeBook1()
	query1 := linq.From(bookList1)
	bookMap1 := make(map[string]string, 0)
	query1.SelectT(func(book Book) linq.KeyValue {
		return linq.KeyValue{
			Key:   book.Author,
			Value: book.Name,
		}
	}).ToMap(&bookMap1)

	for key, value := range bookMap1 {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", key, value))
	}

	fmt.Println()
	fmt.Println()
	bookList2 := MakeBook()
	query2 := linq.From(bookList2)

	bookMap2 := make(map[string]Book, 0)
	query2.ToMapByT(&bookMap2,
		//这个方法决定key的字段是什么
		func(book Book) string {
			return book.Author
		}, //这个方法决定value的字段是什么
		func(book Book) Book {
			return book
		})

	for key, value := range bookMap2 {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", key, value.Name))
	}

	fmt.Println()
	fmt.Println()
	bookListResult := make([]Book, 0)
	linq.From(MakeBook()).OrderByDescendingT(func(book Book) string {
		return book.Author
	}).ToSlice(&bookListResult)
	for _, value := range bookListResult {
		fmt.Println(fmt.Sprintf("作者%v，写了一本书叫做%v", value.Author, value.Name))
	}
}

func Count() {
	bookList := MakeBook()
	query := linq.From(bookList)
	fmt.Println(fmt.Sprintf("数组中发布时间早于2019-1-1的书有%d本", query.CountWithT(PublishTimeBeforeFunc)))
	fmt.Println(fmt.Sprintf("数组中发布时间晚于2018-1-1的书有%d本", query.CountWithT(PublishTimeAfterFunc)))
}

func Skip() {
	bookList := MakeBook()
	query := linq.From(bookList)

	fmt.Println(fmt.Sprintf("跳过三个元素，返回的数组中包含%d个元素", len(query.Skip(3).Results())))
	fmt.Println(fmt.Sprintf("跳过发布时间在2018.1.1之后的元素，返回的数组中包含%d个元素", query.SkipWhile(PublishTimeAfterFunc).Count()))
}

func SelectMany() {
	input := [][]int{{1, 2, 3}, {4, 5, 6, 7}}

	//二位数组进行合并
	linq.From(input).SelectManyT(
		func(i []int) linq.Query {
			return linq.From(i)
		},
	).ForEachT(func(num int) {
		fmt.Println(num)
	})

	fmt.Println()
	fmt.Println()

	//三维数组进行合并
	input1 := [][][]int{{{1, 2, 3}}, {{4, 5, 6, 7}, {8, 9, 10}}}
	linq.From(input1).SelectManyT(
		func(i [][]int) linq.Query {
			return linq.From(i)
		},
	).SelectManyT(
		func(i []int) linq.Query {
			return linq.From(i)
		},
	).ForEachT(func(num int) {
		fmt.Print(num)
		fmt.Print(" ")
	})

	//SelectMany的主要作用是,结构体中包含数组元素，可以把数组元素取出来和结构体做相应的操作
	//如果是结构体中有多个数组的操作，可以把对应的多个数组放到一个新定义的结构体数组中
	fmt.Println()
	men := MakeInnerData()
	linq.From(men).
		//第一个方法筛选出结构体对应的数组
		SelectManyByT(func(man Man) linq.Query {
			return linq.From(man.Pets)
		},
			//第二个方法这边可以拿到数组中元素对应的结构体的数据.
			func(pet Pets, man Man) map[string]Pets {
				return map[string]Pets{
					man.Name: pet,
				}
			}).ForEachT(func(petWithMan map[string]Pets) {
		for key, value := range petWithMan {
			fmt.Println(fmt.Sprintf("%v拥有一只宠物，它的名字叫做%v", key, value.Name))
		}
	})

}

func Select() {
	bookList := MakeBook()
	query := linq.From(bookList)

	//筛选出结构体对应的元素
	query.SelectT(func(book Book) string {
		return book.Name
	}).ForEachT(func(bookName string) {
		fmt.Println(fmt.Sprintf("书名%v", bookName))
	})

	fmt.Println()
	fmt.Println()

	//根据下标进行返回，如果是跟下标有关的，则可以通过这个方法进行组合然后返回.
	query.SelectIndexedT(func(index int, book Book) []string {
		return []string{strconv.Itoa(index + 1), book.Name}
	}).ForEachT(func(book []string) {
		fmt.Println(fmt.Sprintf("第%v本书,书名%v", book[0], book[1]))
	})
}

//OrderBy可以按照对应的规则进行排序，可以实现多重排序，需要注意的就是OrderBy是从到达，orderByDescending是从大到小
func OrderBy() {
	bookList := MakeBook()
	query := linq.From(bookList)
	query = query.OrderByDescendingT(func(book Book) string {
		return book.Name
	}).ThenByDescendingT(func(book Book) string {
		return book.Author
	}).ThenByDescendingT(func(book Book) float64 {
		return book.Money
	}).Query

	query.ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
	})

	fmt.Println()
	fmt.Println()

	//reverse: 把结果逆序
	query.Reverse().ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
	})
}

func Join() {
	//任务和动物有关联，直接找到对应的关系，没有对应关系的数据不会处理，正常情况下应该用更单一的数据来做关联
	persons, pets := MakeJoinData()
	query1 := linq.From(persons)
	query2 := linq.From(pets)
	query1.JoinT(query2,
		func(person Person) string {
			return person.Name
		},
		func(pet Pet) string {
			return pet.OwnerName
		},
		func(person Person, pet Pet) map[string]Pet {
			petMap := make(map[string]Pet)
			petMap[person.Name] = pet
			return petMap
		},
	).ForEachT(func(petMap map[string]Pet) {
		for key, value := range petMap {
			fmt.Println(fmt.Sprintf("%v拥有一只宠物，它的名字叫做:%v", key, value.Name))
		}
	})
}

func Intersect() {
	bookList1 := MakeBook1()
	query1 := linq.From(bookList1)
	bookList2 := MakeBook2()
	query2 := linq.From(bookList2)

	fmt.Println("输出对应的书籍元素，按照结构体判断交集")
	query3 := query1.Intersect(query2)
	query3.ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
	})

	fmt.Println()
	fmt.Println("输出对应的书籍元素，按照元素判断交集")
	query3 = query1.IntersectByT(query2, func(book Book) string {
		return book.Author
	})
	query3.ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
	})

	//这边计算是按照结构体的一个元素进行比较。事实证明无论是否按照结构体的元素筛选，返回的都是完全相同的元素
	fmt.Println()
	fmt.Println("输出对应的书籍元素，按照元素判断交集2")
	bookList1 = MakeBook1()
	query1 = linq.From(bookList1)
	bookList3 := MakeBook3()
	query2 = linq.From(bookList3)
	query3 = query2.IntersectByT(query1, func(book Book) string {
		return book.Author
	})
	query3.ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
	})

	fmt.Println()
	fmt.Println("判断是否会不重复的处理")
	query1 = linq.From([]int{1, 2, 3, 4, 5, 3})
	query2 = linq.From([]int{1, 3, 3, 78, 99, 0, 111})
	query3 = query2.IntersectByT(query1, func(i int) int {
		return i
	})
	fmt.Println(query3.Results())
}

func GroupBy() {
	bookList1 := MakeBook1()
	query := linq.From(bookList1)

	fmt.Println("输出对应书名和作者")
	//第一个func是分组的依据，第二个func是分组后返回的元素.
	query = query.GroupByT(func(book Book) string {
		return book.Author
	}, func(book Book) Book {
		return book
	})

	query.ForEachT(func(bookGroup linq.Group) {
		fmt.Println(fmt.Sprintf("作者是%v", bookGroup.Key))
		for _, item := range bookGroup.Group {
			fmt.Println(fmt.Sprintf("书名%v,作者%v", item.(Book).Name, item.(Book).Author))
			fmt.Println()
		}

	})
	fmt.Println()
	fmt.Println("输出对应书名")

	bookList1 = MakeBook1()
	query = linq.From(bookList1)
	var nameGroups []linq.Group
	query.GroupByT(
		func(book Book) string { return book.Author },
		func(book Book) string { return book.Name },
	).ToSlice(&nameGroups)

	for _, item := range nameGroups {
		fmt.Println(fmt.Sprintf("作者是%v", item.Key))
		for _, row := range item.Group {
			fmt.Println(fmt.Sprintf("书名%v", row))
		}
	}
}

func Except() {
	bookList1 := MakeBook1()
	query1 := linq.From(bookList1)
	bookList2 := MakeBook2()
	query2 := linq.From(bookList2)

	fmt.Println("输出对应书名和作者")
	//判断结构结构体本身是否相等
	query1.Except(query2).
		ForEachT(func(book Book) {
			fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
		})

	fmt.Println()
	fmt.Println("输出对应书名和作者")
	//判断结构体的某个字段是否出现
	query1.ExceptByT(query2, func(book Book) string {
		return book.Author
	}).ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
	})
}

func Distinct() {
	bookList := MakeBook()
	query := linq.From(bookList)
	query = query.Append(Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	bookList1 := make([]Book, 0)
	query.ToSlice(&bookList1)
	fmt.Println(fmt.Sprintf("筛选之前query中总共有%d个数据", query.Count()))
	//这边的distinct如果是结构体，则判断的是结构体的所有值，如果是指针则判断的是指针地址.
	query = query.Distinct()
	fmt.Println(fmt.Sprintf("筛选之后query中总共有%d个数据", query.Count()))

	fmt.Println()
	fmt.Println("输出对应书名和作者")
	query = query.Append(Book{
		Name:        "Go语言(第二版)",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	//根据对应的规则来判断是否相同 -- 如果对应的字段一样，只会保留第一个找到的，后面的会被清除掉
	query.DistinctByT(func(book Book) string {
		return book.Author
	}).ForEachT(func(book Book) {
		fmt.Println(fmt.Sprintf("书名%v,作者%v", book.Name, book.Author))
	})
}

func Combine() {
	bookList := MakeBook()
	query := linq.From(bookList)
	query1 := query.Concat(query) //不会排除重复项
	query2 := query.Union(query)  //会排除重复项
	fmt.Println(fmt.Sprintf("不会排除重复项的方法返回的元素个数是%d", len(query1.Results())))
	fmt.Println(fmt.Sprintf("会排除重复项的方法返回的元素个数是%d", len(query2.Results())))
	query1 = query1.Append(Book{Name: "Last One"})   //加在原来的Query的最后面
	query2 = query2.Prepend(Book{Name: "First One"}) //加在原来的Query的第一个
	fmt.Println(fmt.Sprintf("最后一个元素是%v", query1.Last()))
	fmt.Println(fmt.Sprintf("第一个元素是%v", query2.First()))
}

func FirstAndLast() {
	bookList := MakeBook()
	query := linq.From(bookList)
	fmt.Println(fmt.Sprintf("数组第一本书书名叫:%s", query.First().(Book).Name))
	fmt.Println(fmt.Sprintf("数组第一本书发布时间在2018年1月1号之后的书书名叫:%s", query.FirstWith(PublishTimeAfterFunc).(Book).Name))

	fmt.Println(fmt.Sprintf("数组最后一本书书名叫:%s", query.Last().(Book).Name))
	fmt.Println(fmt.Sprintf("数组最后一本书发布时间在2018年1月1号之后的书书名叫:%s", query.LastWith(PublishTimeAfterFunc).(Book).Name))
}

func Contains() {
	bookList := MakeBook()
	query := linq.From(bookList)
	//这边需要注意是的是如果要判断两个元素是否相等，在go里面不能使用元素的地址
	// 所以这边加到query里面的必须不能是元素的地址，否则没办法比较结构体本身，而是比较结构体的地址
	fmt.Println(fmt.Sprintf("判断当前数组是否包含相同的元素,判断的是所有的值，而不是地址，结果是%v", query.Contains(Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})))
}

func Calculate() {
	query := linq.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(fmt.Sprintf("平均值为：%f", query.Average()))
	fmt.Println(fmt.Sprintf("最大的值为%d", query.Max()))
	fmt.Println(fmt.Sprintf("最小的值为%d", query.Min()))
}

func AnyWith() {
	bookList := MakeBook()
	query := linq.From(bookList)
	result := query.AnyWith(PublishTimeBeforeFunc) //判断是否有任意个元素满足条件
	fmt.Println(fmt.Sprintf("是否有任意书的发表时间早于2020年1月1号，答案是%v", result))

	result = query.AnyWith(PublishTimeAfterFunc2) //判断是否有任意个元素满足条件
	fmt.Println(fmt.Sprintf("是否有任意书的发表时间晚于2021年1月1号，答案是%v", result))
}

func All() {
	bookList := MakeBook()
	query := linq.From(bookList)
	result := query.All(PublishTimeBeforeFunc)
	fmt.Println(fmt.Sprintf("是否所有的书的发表时间都早于2020年1月1号，答案是%v", result))

	result = query.All(PublishTimeAfterFunc)
	fmt.Println(fmt.Sprintf("是否所有的书的发表时间都晚于2018年1月1号，答案是%v", result))
}

func Aggregate() {
	bookList := MakeBook()
	query := linq.From(bookList)
	// 这边的聚合操作其实就是【自定义对应的比较方法】。
	// 根据注释，第一次拿到的是数组的第一个元素，然后跟第二个元素进行比较，如果第二个元素符合条件则【替换第一个元素】，否则当前聚合操作还是持有第一个元素。
	// 以此类推后面的元素
	result := query.Aggregate(AggregateFunc).(Book)
	fmt.Println(fmt.Sprintf("聚合操作查询出的结果是【%s】这本书，它的发布时间最早为：%v", result.Name, timed.GetTimeDefaultFormatString(result.PublishTime.Unix())))
}
