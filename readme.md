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