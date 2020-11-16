/*
 *  @Author : huangzj
 *  @Time : 2020/11/12 10:52
 *  @Description：这边是linq包的使用操作，为之后的使用做一些示例
 *  包引用命令： go get github.com/ahmetalpbalkan/go-linq
 */

package linqUse

import (
	"Go-Tool/util/timed"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"testing"
	"time"
)

func Test(t *testing.T) {
	bookList := MakeBook()
	query := linq.From(bookList)

	//Aggregate(query) //自定义聚合操作
	//
	//All(query) //判断数组的所有元素是否都满足条件
	//
	//AnyWith(query) //判断数组是否有任意个元素满足条件

	//Calculate() //这个方法只能计算数值类型

	//Combine(query) //组合两个数组的两种方式

	fmt.Println(fmt.Sprintf("跳过三个元素，返回的数组中包含%d个元素", len(query.Skip(3).Results())))
	fmt.Println(fmt.Sprintf("跳过发布时间在2020.1.1之前的元素，返回的数组中包含%d个元素", len(query.SkipWhile(PublishTimeBeforeFunc).Results()))) //这个方法好像有问题，没办法过滤
	author := make([]string, 0)
	query.Select(func(book interface{}) interface{} {
		return book.(Book).Author
	}).ToSlice(&author)
	fmt.Println(author)
	query1 := linq.From([][]Book{bookList, bookList})
	query1.SelectMany(func(query interface{}) linq.Query {
		fmt.Println(query)
		return linq.From(query)
	}).Select(func(book interface{}) interface{} {
		return book.(Book).Author
	}).ToSlice(&author)
	fmt.Println(author)

	//Contains(query) //判断数组中是否包含某个元素
	//query.CountWith()
	//
	//query.Distinct()
	//query.DistinctBy()
	//
	//query.OrderBy()
	//query.OrderByDescending()

	//FirstAndLast(query) //获取第一个/最后一个元素

}

func Combine(query linq.Query) {
	query1 := query.Concat(query) //不会排除重复项
	query2 := query.Union(query)  //会排除重复项
	fmt.Println(fmt.Sprintf("不会排除重复项的方法返回的元素个数是%d", len(query1.Results())))
	fmt.Println(fmt.Sprintf("会排除重复项的方法返回的元素个数是%d", len(query2.Results())))
}

func FirstAndLast(query linq.Query) {
	fmt.Println(fmt.Sprintf("数组第一本书书名叫:%s", query.First().(Book).Name))
	fmt.Println(fmt.Sprintf("数组第一本书发布时间在2020年1月1号之前的书书名叫:%s", query.FirstWith(PublishTimeBeforeFunc).(Book).Name))

	fmt.Println(fmt.Sprintf("数组最后一本书书名叫:%s", query.Last().(Book).Name))
	fmt.Println(fmt.Sprintf("数组最后一本书发布时间在2020年1月1号之前的书书名叫:%s", query.LastWith(PublishTimeBeforeFunc).(Book).Name))
}

func Contains(query linq.Query) {
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

func AnyWith(query linq.Query) {
	result := query.AnyWith(PublishTimeBeforeFunc) //判断是否有任意个元素满足条件
	fmt.Println(fmt.Sprintf("是否有任意书的发表时间早于2020年1月1号，答案是%v", result))

	result = query.AnyWith(PublishTimeAfterFunc2) //判断是否有任意个元素满足条件
	fmt.Println(fmt.Sprintf("是否有任意书的发表时间晚于2021年1月1号，答案是%v", result))
}

func All(query linq.Query) {
	result := query.All(PublishTimeBeforeFunc)
	fmt.Println(fmt.Sprintf("是否所有的书的发表时间都早于2020年1月1号，答案是%v", result))

	result = query.All(PublishTimeAfterFunc)
	fmt.Println(fmt.Sprintf("是否所有的书的发表时间都晚于2018年1月1号，答案是%v", result))
}

func Aggregate(query linq.Query) {
	// 这边的聚合操作其实就是【自定义对应的比较方法】。
	// 根据注释，第一次拿到的是数组的第一个元素，然后跟第二个元素进行比较，如果第二个元素符合条件则【替换第一个元素】，否则当前聚合操作还是持有第一个元素。
	// 以此类推后面的元素
	result := query.Aggregate(AggregateFunc).(*Book)
	fmt.Println(fmt.Sprintf("聚合操作查询出的结果是【%s】这本书，它的发布时间最早为：%v", result.Name, timed.GetTimeDefaultFormatString(result.PublishTime.Unix())))
}
