package marshal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// 抄的

const tagName = "gson"

func isNeedSkipParse(kind reflect.Kind) bool {
	switch kind {
	case reflect.Chan, reflect.Complex128, reflect.Complex64, reflect.Func, reflect.Invalid:
		return true
	default:
		return false
	}
}

func marshal(v any) (string, error) {
	var err error
	// 空接口直接返回 null
	if v == nil {
		return "null", nil
	}

	// 获取空接口的真实类型
	vType := reflect.TypeOf(v)
	vKind := vType.Kind()

	if isNeedSkipParse(vKind) {
		return "", errors.New("this type is not supported")
	}

	switch vKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		// 如果是这些类型直接返回数值即可
		return fmt.Sprintf("%v", reflect.ValueOf(v)), nil
	case reflect.String:
		// 如果是字符串类型，两边需要加上引号 "
		return fmt.Sprintf("\"%v\"", reflect.ValueOf(v)), nil
	case reflect.Map:
		// 通过 reflect.ValueOf().MapKey() 获取 map 的所有 key
		// 然后 通过 reflect.ValueOf().MapIndex 获取对应的 val
		var res []string
		for _, val := range reflect.ValueOf(v).MapKeys() {
			key := val.Interface()
			value := reflect.ValueOf(v).MapIndex(val)
			if !value.CanInterface() || isNeedSkipParse(reflect.TypeOf(value.Interface()).Kind()) {
				continue
			}

			result, err := marshal(value.Interface())
			if err != nil {
				return "", err
			}

			res = append(res, fmt.Sprintf("\"%v\":%v", key, result))
		}
		return fmt.Sprintf("{%v}", strings.Join(res, ",")), nil
	case reflect.Array, reflect.Slice:
		// 数组或者切片
		// 数组/切片的值
		vVal := reflect.ValueOf(v)
		// 数组/切片的长度
		vLen := vVal.Len()
		res := make([]string, vLen)

		// 对数组中的每一个元素都进行一次序列化
		for i := 0; i < vLen; i++ {
			res[i], err = marshal(vVal.Index(i).Interface())
			if err != nil {
				return "", err
			}
		}
		// 对每个元素之间加上逗号后返回结果
		return fmt.Sprintf("[%v]", strings.Join(res, ",")), nil
	case reflect.Struct:
		// 结构体
		// 思路和数组切片差不多，只不过要获取每个字段的标签
		vVal := reflect.ValueOf(v)
		vFieldNum := vVal.NumField()
		var res []string

		for i := 0; i < vFieldNum; i++ {
			// 处理每一个字段
			field := vVal.Field(i)   // 字段
			fieldType := vVal.Type() // 字段类型

			if !field.CanInterface() || isNeedSkipParse(fieldType.Kind()) {
				// 如果这个类型不能序列化，则继续遍历
				continue
			}

			key := fieldType.Name() // 该字段名称

			tag := fieldType.Field(i).Tag.Get(tagName) // 获取对应的 tag 内容，例如 `mytag:"abcd"`
			// 如果不为空是直接用的字段名，可以自己加入个性化处理，驼峰自动分割啥的
			if tag != "" {
				key = tag
			}

			// 同样是一个个处理
			result, err := marshal(field.Interface())
			if err != nil {
				return "", err
			}

			// 以 "key":val 的格式添加
			res = append(res, fmt.Sprintf("\"%v\":%v", key, result))
		}
		// 每个元素间加逗号
		return fmt.Sprintf("{%v}", strings.Join(res, ",")), nil
	default:
		return "", nil
	}
}

func Marshal(v any) ([]byte, error) {
	res, err := marshal(v)
	if err != nil {
		return nil, err
	}
	return []byte(res), err
}
