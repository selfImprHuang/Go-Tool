/*
 *  @Author : huangzj
 *  @Time : 2020/12/11 16:16
 *  @Description：
 */

package testValidator

type Ower struct {
	Age int `validate:"min=10"`
}

type Cat struct {
	Age  int `validate:"min=10"`
	Ower Ower
}

type Man struct {
	Name string `validate:"ipe"`
}

type Woman struct {
	Name string `validate:"diy=5"`
}

type Person struct {
	Name    string
	M       map[string]int `validate:"min=10"`
	H       int            `validate:"max=20"`
	Exist   bool           `validate:"exists"`
	Require bool           `validate:"required"`
	Cat     Cat
}
