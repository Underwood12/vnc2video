// Code generated by "stringer -type=Button"; DO NOT EDIT.

package vnc2video

import "fmt"

const (
	_Button_name_0 = "BtnNoneBtnLeftBtnMiddle"
	_Button_name_1 = "BtnRight"
	_Button_name_2 = "BtnFour"
	_Button_name_3 = "BtnFive"
	_Button_name_4 = "BtnSix"
	_Button_name_5 = "BtnSeven"
	_Button_name_6 = "BtnEight"
)

var (
	_Button_index_0 = [...]uint8{0, 7, 14, 23}
	_Button_index_1 = [...]uint8{0, 8}
	_Button_index_2 = [...]uint8{0, 7}
	_Button_index_3 = [...]uint8{0, 7}
	_Button_index_4 = [...]uint8{0, 6}
	_Button_index_5 = [...]uint8{0, 8}
	_Button_index_6 = [...]uint8{0, 8}
)

func (i Button) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _Button_name_0[_Button_index_0[i]:_Button_index_0[i+1]]
	case i == 4:
		return _Button_name_1
	case i == 8:
		return _Button_name_2
	case i == 16:
		return _Button_name_3
	case i == 32:
		return _Button_name_4
	case i == 64:
		return _Button_name_5
	case i == 128:
		return _Button_name_6
	default:
		return fmt.Sprintf("Button(%d)", i)
	}
}
