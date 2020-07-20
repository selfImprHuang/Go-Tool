/*
 *  @Author : huangzj
 *  @Time : 2020/7/19 20:46
 *  @Description：深度搜索工具类，在现实问题中，用于比如一组道具，抽取一定数量，必须超过总价值多少，有几种组合方式
 *  剪枝策略，参考：https://blog.csdn.net/qq_38790716/article/details/88051412
 */

package deepSearch

type DeepFirstSearch struct {
	length        int           //道具数组长度
	vis           []bool        //标识 --true表示这个节点走过了，false表示这个节点没走过
	limitValue    int           //最小限制值(即抽中的道具需要至少最低保底多少价值)
	limitMaxValue int           //最大限制值(最大值区间，只用于计算在保底和区间之间的方法中,不超过)
	LimitNum      int           //限制数字的数量(最多只能有几个数字)
	sum           int           //总数值
	valueList     []ItemValue   //道具价值数组
	tempList      []int         //下标数组(用来存储当前对应的道具下标)
	itemGroup     [][]ItemValue //最终的组合数组
}

func NewDeepSearch() *DeepFirstSearch {
	return &DeepFirstSearch{
		sum: 0,
	}
}

//计算n个道具的所有组合，这个是最原始的深度优先搜索算法
func (srv *DeepFirstSearch) GetItemMatchNum(list []ItemValue, num int) [][]ItemValue {
	//不校验数量，如果超过数组的最大数量,就是失败了.返回空数组

	srv.length = len(list)
	srv.vis = make([]bool, srv.length)
	srv.valueList = list
	srv.tempList = make([]int, srv.length)
	srv.itemGroup = make([][]ItemValue, 0)
	srv.LimitNum = num

	srv.dfs(0, 0)

	return srv.itemGroup
}

//valueList：物品价值数组 一维数组的组成是：[编号,价值]
//limitValue：物品价值限制
//题目可以理解为：有一个长度为n的数组valueList，从其中选择出1~n个元素，保证这些元素相加的值能够至少超过limitValue.
//找到所有对应的组合即可
func (srv *DeepFirstSearch) GetItemMatchValue(list []ItemValue, limit int) [][]ItemValue {
	//这边不校验所有的价值总和是否要高于limit，如果没有超过数组就是0，也是符合条件的.

	srv.length = len(list)
	srv.vis = make([]bool, srv.length)
	srv.limitValue = limit
	srv.valueList = list
	srv.tempList = make([]int, srv.length)
	srv.itemGroup = make([][]ItemValue, 0)

	srv.dfsBeyondValue(0, 0)

	return srv.itemGroup
}

//获取限制N个数量的保底计算方法,全面两个参数合和 #GetItemMatchValue# 一致
func (srv *DeepFirstSearch) GetItemMatchValueLimitNum(list []ItemValue, limit int, num int) [][]ItemValue {
	//不校验，没有结果返回空数组
	srv.length = len(list)
	srv.vis = make([]bool, srv.length)
	srv.limitValue = limit
	srv.valueList = list
	srv.tempList = make([]int, srv.length)
	srv.itemGroup = make([][]ItemValue, 0)
	srv.LimitNum = num

	srv.dfsLimitNum(0, 0)

	return srv.itemGroup
}

func (srv *DeepFirstSearch) GetItemMatchValueInSection(list []ItemValue, min int, max int) [][]ItemValue {
	//不校验，查不到就返回空数组
	srv.length = len(list)
	srv.vis = make([]bool, srv.length)
	srv.limitValue = min
	srv.valueList = list
	srv.tempList = make([]int, srv.length)
	srv.itemGroup = make([][]ItemValue, 0)
	srv.limitMaxValue = max

	srv.dfsInSection(0, 0)

	return srv.itemGroup
}

func (srv *DeepFirstSearch) dfs(step int, pos int) {
	//满足条件，添加组合方案
	srv.generateGroupDfs(step)

	for i := pos; i < srv.length; i++ {
		//vis为false表示该位置没有被进入过 &&
		if srv.vis[i] == false {
			srv.tempList[step] = i
			srv.sum = srv.sum + srv.valueList[i].ItemValue //计算总数值,这个是需要到下一个step的dfs中进行判断的
			srv.vis[i] = true                              //该位置已经被进入了，设置为true

			srv.dfs(step+1, i+1) //进行下一步的深度搜索

			srv.vis[i] = false                             //搜索结束后返回，标识置为false，表示可以再进入
			srv.sum = srv.sum - srv.valueList[i].ItemValue //总数要减掉之前的值，才是现在的值
		}
	}
	return
}

func (srv *DeepFirstSearch) dfsBeyondValue(step int, pos int) {
	//满足条件，添加组合方案
	srv.generateGroup(step)

	for i := pos; i < srv.length; i++ {
		//vis为false表示该位置没有被进入过 &&
		if srv.vis[i] == false {
			srv.tempList[step] = i
			srv.sum = srv.sum + srv.valueList[i].ItemValue //计算总数值,这个是需要到下一个step的dfs中进行判断的
			srv.vis[i] = true                              //该位置已经被进入了，设置为true

			srv.dfsBeyondValue(step+1, i+1) //进行下一步的深度搜索

			srv.vis[i] = false                             //搜索结束后返回，标识置为false，表示可以再进入
			srv.sum = srv.sum - srv.valueList[i].ItemValue //总数要减掉之前的值，才是现在的值
		}
	}
	return
}

func (srv *DeepFirstSearch) dfsInSection(step int, pos int) {
	//满足条件，添加组合方案
	srv.generateGroupInSection(step)

	for i := pos; i < srv.length; i++ {
		//vis为false表示该位置没有被进入过 &&
		if srv.vis[i] == false {
			srv.tempList[step] = i
			srv.sum = srv.sum + srv.valueList[i].ItemValue //计算总数值,这个是需要到下一个step的dfs中进行判断的
			srv.vis[i] = true                              //该位置已经被进入了，设置为true

			srv.dfsInSection(step+1, i+1) //进行下一步的深度搜索

			srv.vis[i] = false                             //搜索结束后返回，标识置为false，表示可以再进入
			srv.sum = srv.sum - srv.valueList[i].ItemValue //总数要减掉之前的值，才是现在的值
		}
	}
	return
}

func (srv *DeepFirstSearch) dfsLimitNum(step int, pos int) {
	//满足条件，添加组合方案
	srv.generateGroupLimitNum(step)

	for i := pos; i < srv.length; i++ {
		//vis为false表示该位置没有被进入过 &&
		if srv.vis[i] == false {
			srv.tempList[step] = i
			srv.sum = srv.sum + srv.valueList[i].ItemValue //计算总数值,这个是需要到下一个step的dfs中进行判断的
			srv.vis[i] = true                              //该位置已经被进入了，设置为true

			srv.dfsLimitNum(step+1, i+1) //进行下一步的深度搜索

			srv.vis[i] = false                             //搜索结束后返回，标识置为false，表示可以再进入
			srv.sum = srv.sum - srv.valueList[i].ItemValue //总数要减掉之前的值，才是现在的值
		}
	}
	return
}

//满足条件，添加组合方案
func (srv *DeepFirstSearch) generateGroupDfs(step int) {
	//这边不直接return是因为可能后面还有元素没进组合，这个也是允许的，所有的组合,如果直接return的话，我搜索到[1,2,3]满足就不会再去搜索[1,2,3,4].
	if step == srv.LimitNum {
		itemList := make([]ItemValue, 0)
		for i := 0; i < step; i++ {
			itemList = append(itemList, srv.valueList[srv.tempList[i]])
		}
		srv.itemGroup = append(srv.itemGroup, itemList)
	}
}

//满足条件，添加组合方案
func (srv *DeepFirstSearch) generateGroup(step int) {
	//这边不直接return是因为可能后面还有元素没进组合，这个也是允许的，所有的组合,如果直接return的话，我搜索到[1,2,3]满足就不会再去搜索[1,2,3,4].
	if srv.sum >= srv.limitValue {
		itemList := make([]ItemValue, 0)
		for i := 0; i < step; i++ {
			itemList = append(itemList, srv.valueList[srv.tempList[i]])
		}
		srv.itemGroup = append(srv.itemGroup, itemList)
	}
}

//满足条件，添加组合方案
func (srv *DeepFirstSearch) generateGroupLimitNum(step int) {
	//这边不直接return是因为可能后面还有元素没进组合，这个也是允许的，所有的组合,如果直接return的话，我搜索到[1,2,3]满足就不会再去搜索[1,2,3,4].
	if srv.sum >= srv.limitValue && step == srv.LimitNum {
		itemList := make([]ItemValue, 0)
		for i := 0; i < step; i++ {
			itemList = append(itemList, srv.valueList[srv.tempList[i]])
		}
		srv.itemGroup = append(srv.itemGroup, itemList)
	}
}

//满足条件，添加组合方案
func (srv *DeepFirstSearch) generateGroupInSection(step int) {
	//这边不直接return是因为可能后面还有元素没进组合，这个也是允许的，所有的组合,如果直接return的话，我搜索到[1,2,3]满足就不会再去搜索[1,2,3,4].
	if srv.sum >= srv.limitValue && srv.sum <= srv.limitMaxValue {
		itemList := make([]ItemValue, 0)
		for i := 0; i < step; i++ {
			itemList = append(itemList, srv.valueList[srv.tempList[i]])
		}
		srv.itemGroup = append(srv.itemGroup, itemList)
	}
}
