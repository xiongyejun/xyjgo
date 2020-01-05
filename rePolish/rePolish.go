// 中缀表达式转后缀表格是，逆波兰
package rePolish

import (
	"errors"
	"strconv"

	"github.com/xiongyejun/xyjgo/stack"
)

const (
	OP_L_括号 int = 0
	OP_R_括号 int = 1

	OP_ADD int = 2
	OP_SUB int = OP_ADD

	OP_MUL int = 3
	OP_除   int = OP_MUL
)

var mOp map[string]int = map[string]int{
	"+": OP_ADD,
	"-": OP_SUB,
	"*": OP_MUL,
	"/": OP_除,
	"(": OP_L_括号,
	")": OP_R_括号,
}

func RePolish(str []string) (ret []string, err error) {
	s := stack.New(10)

	var tmp string
	var itf interface{}
	for i := range str {
		tmp = str[i]

		if op, ok := mOp[tmp]; ok {
			switch op {
			case OP_L_括号:
				s.Push("(")
			case OP_R_括号:
				// 找到左括号为止
				for {
					if itf, err = s.Pop(); err != nil {
						return
					}
					if itf.(string) == "(" {
						break
					}
					ret = append(ret, itf.(string))
				}
			default:
				var optmp int
				var ok bool
				for !s.IsEmpty() {
					// 从stack中找到比当前op运算优先级更小的op为止
					if itf, err = s.Top(); err != nil {
						return
					}
					if optmp, ok = mOp[itf.(string)]; !ok {
						err = errors.New("stack中存在不是操作符的字符.")
						return
					}
					if optmp < op {
						break
					}
					if itf, err = s.Pop(); err != nil {
						return
					}
					ret = append(ret, itf.(string))
				}
				// 操作符还是要入stack
				s.Push(tmp)
			}
		} else {
			// 操作数
			ret = append(ret, tmp)
		}

		// fmt.Printf("%d  %s\t", i, tmp)
		// s.Do(fmt.Print)
		// fmt.Println("\t", ret)
	}

	for !s.IsEmpty() {
		if itf, err = s.Pop(); err != nil {
			return
		}
		ret = append(ret, itf.(string))
	}
	return
}

func Calc(str []string) (ret float64, err error) {
	s := stack.New(10)
	var tmp string
	var itf interface{}
	var f1, f2 float64
	for i := range str {
		tmp = str[i]
		switch tmp {
		case "+":
			if f1, f2, err = getTwoValue(s); err != nil {
				return
			}
			s.Push(f2 + f1)
		case "-":
			if f1, f2, err = getTwoValue(s); err != nil {
				return
			}
			s.Push(f2 - f1)
		case "*":
			if f1, f2, err = getTwoValue(s); err != nil {
				return
			}
			s.Push(f2 * f1)
		case "/":
			if f1, f2, err = getTwoValue(s); err != nil {
				return
			}
			s.Push(f2 / f1)

		default:
			// 操作数入stack
			s.Push(tmp)
		}
	}
	if itf, err = s.Pop(); err != nil {
		return
	}

	return Itf2Float(itf)
}

func getTwoValue(s *stack.Stack) (f1, f2 float64, err error) {
	var itf interface{}
	if itf, err = s.Pop(); err != nil {
		return
	}

	if f1, err = Itf2Float(itf); err != nil {
		return
	}

	if itf, err = s.Pop(); err != nil {
		return
	}

	if f2, err = Itf2Float(itf); err != nil {
		return
	}

	return
}

func Itf2Float(itf interface{}) (ret float64, err error) {
	switch itf.(type) {
	case string:
		ret, err = strconv.ParseFloat(itf.(string), 64)
	case float32:
		ret = float64(itf.(float32))
	case float64:
		ret = itf.(float64)
	case int:
		ret = float64(itf.(int))

	case int32:
		ret = float64(itf.(int32))
	default:
		ret = 0
	}
	return
}
