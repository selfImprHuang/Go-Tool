# 作用
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;本工程仅供个人学习记录，刚开始学习go（之前是java），本工程积累一些反射的用法。后期查询更加快捷

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;

# 更新日志
2020/4/24 ：创建FiledValueTool.go四个方法，分别处理【结构体对象】、【结构体指针对象】的字段名-字段值的map获取以及 【结构体类型】、【结构体指针类型】的字段名获取

2020/4/27 ：创建FileUtil.go、FileReadTool.go、FileWriteTool.go文件，添加文件文件夹操作、读写文件操作方法

2020/4/27 ：创建ColorUtil：颜色输出工具、修改目录层级、添加数据库表字段处理工具（需要已连接数据库）

2020/4/28 : 创建IniTool.go:包装int工具类，进行文件配置数据结构的读取，包括对应的使用例子(工具地址：https://ini.unknwon.cn/docs/howto/work_with_sections)

2020/4/30 ：添加FileUtil.go、FileReadTool.go单测类代码，修改相应bug-----IniTool.go依赖获取不到，注释掉代码

2020/4/30 : 创建ArrayUtil.go，添加数组对应方法

2020/5/6 : 创建ArrayUtils测试类，修改对应方法

2020/5/7 : 创建timed包相关时间操作工具，实现时间戳、时间格式化操作。创建timed测试类

2020/6/15: 创建Set，实现相关方法

2020/6/30: ArrayList添加深度复制方法以及单测类、添加Rand随机获取参数方法、RandByWeight根据权重随机获取奖励方法

2020/7/8: 增加读取json文件方法及其单测类

2020/7/13: 增加计算获取字符串哈希值的方法(这个是之前用于做分表系统的方法，搬到这边来看一下)

2020/7/13: 修改FileUtil的报错

2020/7/20: 添加深度优先搜索算法工具类以及变种，分别对应以下场景

- 在N个道具中，提取所有M个道具的组合方案,这种方式可以用在 权重一样的 N个道具中选择一定的数量的道具组合(这种方案就是权重一样)
- 在N个道具中，提取超过一定价值V的所有组合方案，道具的数量没有限制，但是必须超过最低价值，也就是我们说的保底方案
- 在N个道具中，提取超过一定价值VMin,但是不超过一定价值VMax的组合方案，道具的数量没有限制，必须在最低价值和最高价值之间.
- 在N个道具中，提取超过一定价值V并且个数大小为M的所有组合方案，这个就是保证 保底的一定数量的道具

类似上面的方案,在游戏中可以在活动创建的时候，初始化所有的方案放到redis，然后每一次抽取都筛选下标来确认方案.这样就可以保证在一定的范围或者数量内确认对应的方案。


2020/8/6:  添加求背包问题最优解三种方式(KnapsackOptimizationUtil.go)及其单测类，添加求背包问题最优解过程路径(KnapsackSearchAnswerUtil.go)及其单测类.添加背包问题的readme相关说明

2020/8/18: 修改单测位置，添加lomuto查找第k小元素算法代码及其单测（分治法）

2020/8/20：新增二分法查找算法及其单测、插入查询及其单测、lomuto划分查询第k小元素及其单测、归并排序及其单测、快速排序及其单测.

2020/8/20: 新增选择排序及其单测、冒泡排序及其单测、堆排序及其单测


# mod vendor模式加载包
通过go mod的方式加载的github上面的包会有报红的问题，但是包本身是可以运行的，这样就是会有一个问题，如果你想要点击去看方法的内容，没办法做到

查阅了网上提供的处理方式，通过设置GOPATH 和Go Modules(vgo)可以解决相应的问题，但是我在这样处理之后依然会出现报红的问题

后面通过使用vendor包的方式来处理，就避免了报红的问题

处理步骤如下 ：
- 通过 go get github.com/kardianos/govendor 命令下载govendor命令
- 通过 go mod vendor 切换到vendor管理
