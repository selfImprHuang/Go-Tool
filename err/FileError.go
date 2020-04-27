/*
 *  @Author : huangzj
 *  @Time : 2020/4/26 15:09
 *  @Description：
 */

package err

type FileError struct {
	Err error
	Mes string
}

func ENewFileError(err error) *FileError {
	return &FileError{
		Err: err,
		Mes: "",
	}
}

func NewFileError(mes string) *FileError {
	return &FileError{
		Err: nil,
		Mes: mes,
	}
}

func (e *FileError) Error() string {
	if e == nil {
		return "<nil>"
	}

	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Mes
}

/*
 * 默认Error方法直接返回Err信息，所以这边只需要加一个方法返回Mes就可以
 */
func (e *FileError) ErrorMes() string {
	return e.Mes
}
