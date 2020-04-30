/*
 *  @Author : huangzj
 *  @Time : 2020/4/27 17:20
 *  @Description： 工具类只对ini读取进行一层封装。参考地址：https://ini.unknwon.cn/docs/howto/work_with_sections
 */

package ini

//
//import (
//	"go-ini/ini"
//	"time"
//)
//
//func NewIniFileObj(iniPath string) (*ini.File, error) {
//	return ini.Load(iniPath)
//}
//
///*
// * @param f ini文件对象
// * @return name->Section 的Map
// * @description 通过所有Section对应的name->Section的map
// */
//func GetSectionMap(f *ini.File) map[string]*ini.Section {
//	sectionMap := make(map[string]*ini.Section, 0)
//	sections := f.Sections()
//	for _, row := range sections {
//		sectionMap[row.Name()] = row
//	}
//
//	return sectionMap
//}
//
///*
// * @param f ini文件对象
// * @param sectionName 区间名字
// * @return key->value的map
// * @description 返回对应Section名字的key-value的map结构,如果没有这个section返回空map
// */
//func GetKeyValueBySection(f *ini.File, sectionName string) map[string]string {
//	keyValueMap := make(map[string]string, 0)
//	sec := f.Section(sectionName)
//	for _, row := range sec.Keys() {
//		keyValueMap[row.Name()] = row.Value()
//	}
//
//	return keyValueMap
//}
//
///*
// * @param f ini文件对象
// * @return {section -> [{key->value}]}
// * @description 返回每个区间对应的所有key、value的结构，具体返回结构如下：
// * {
// *   "sectionName1":[
//			"keyName1":KeyValue1
// * 			.....
// *    ],
// *    .....
// *	}
//*/
//func GetKeyValueByALLSection(f *ini.File) {
//	secMap := make(map[string][]map[string]string, 0)
//	for _, row := range f.Sections() {
//		keyList := make([]map[string]string, 0)
//		for _, r := range row.Keys() {
//			keyMap := make(map[string]string, 0)
//			keyMap[r.Name()] = r.Value()
//			keyList = append(keyList, keyMap)
//		}
//
//		secMap[row.Name()] = keyList
//	}
//}
//
///*
// * 使用例子，官网说明文档可以看到，这边补记录一下.
// */
//func example(cfg *ini.File) {
//	//获取键值时设定候选值，如果没有找到，则返回默认值
//	cfg.Section("SectionName").Key("KeyName").In("default", []string{"str", "arr", "types"})
//	cfg.Section("SectionName").Key("KeyName").InFloat64(1.1, []float64{1.25, 2.5, 3.75})
//	cfg.Section("SectionName").Key("KeyName").InInt(5, []int{10, 20, 30})
//	cfg.Section("SectionName").Key("KeyName").InInt64(10, []int64{10, 20, 30})
//	cfg.Section("SectionName").Key("KeyName").InUint(4, []uint{3, 6, 9})
//	cfg.Section("SectionName").Key("KeyName").InUint64(8, []uint64{3, 6, 9})
//	//这边的三个时间不一样，为了防止报错这么写...
//	cfg.Section("SectionName").Key("KeyName").InTimeFormat(time.RFC3339, time.Now(), []time.Time{time.Now(), time.Now(), time.Now()})
//	cfg.Section("SectionName").Key("KeyName").InTime(time.Now(), []time.Time{time.Now(), time.Now(), time.Now()}) // RFC3339
//
//	//验证获取的值是否在指定范围内,第一个参数是默认值，后面两个值分别是范围
//	cfg.Section("SectionName").Key("KeyName").RangeFloat64(0.0, 1.1, 2.2)
//	cfg.Section("SectionName").Key("KeyName").RangeInt(0, 10, 20)
//	cfg.Section("SectionName").Key("KeyName").RangeInt64(0, 10, 20)
//	cfg.Section("SectionName").Key("KeyName").RangeTimeFormat(time.RFC3339, time.Now(), time.Now(), time.Now())
//	cfg.Section("SectionName").Key("KeyName").RangeTime(time.Now(), time.Now(), time.Now()) // RFC3339
//
//	//自动分割键值到切片,当存在无效输入时，使用零值代替
//	// Input: 1.1, 2.2, 3.3, 4.4 -> [1.1 2.2 3.3 4.4]
//	// Input: how, 2.2, are, you -> [0.0 2.2 0.0 0.0]
//	cfg.Section("SectionName").Key("KeyName").Strings(",")
//	cfg.Section("SectionName").Key("KeyName").Float64s(",")
//	cfg.Section("SectionName").Key("KeyName").Ints(",")
//	cfg.Section("SectionName").Key("KeyName").Int64s(",")
//	cfg.Section("SectionName").Key("KeyName").Uints(",")
//	cfg.Section("SectionName").Key("KeyName").Uint64s(",")
//	cfg.Section("SectionName").Key("KeyName").Times(",")
//
//	//自动分割键值到切片,从结果切片中剔除无效输入
//	// Input: 1.1, 2.2, 3.3, 4.4 -> [1.1 2.2 3.3 4.4]
//	// Input: how, 2.2, are, you -> [2.2]
//	cfg.Section("SectionName").Key("KeyName").ValidFloat64s(",")
//	cfg.Section("SectionName").Key("KeyName").ValidInts(",")
//	cfg.Section("SectionName").Key("KeyName").ValidInt64s(",")
//	cfg.Section("SectionName").Key("KeyName").ValidUints(",")
//	cfg.Section("SectionName").Key("KeyName").ValidUint64s(",")
//	cfg.Section("SectionName").Key("KeyName").ValidTimes(",")
//
//	//自动分割键值到切片,当存在无效输入时，直接返回错误
//	// Input: 1.1, 2.2, 3.3, 4.4 -> [1.1 2.2 3.3 4.4]
//	// Input: how, 2.2, are, you -> error
//	cfg.Section("SectionName").Key("KeyName").StrictFloat64s(",")
//	cfg.Section("SectionName").Key("KeyName").StrictInts(",")
//	cfg.Section("SectionName").Key("KeyName").StrictInt64s(",")
//	cfg.Section("SectionName").Key("KeyName").StrictUints(",")
//	cfg.Section("SectionName").Key("KeyName").StrictUint64s(",")
//	cfg.Section("SectionName").Key("KeyName").StrictTimes(",")
//}
