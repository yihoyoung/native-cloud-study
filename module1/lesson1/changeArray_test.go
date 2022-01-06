package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc        string
		index       int
		expectValue string
	}{
		{
			"测试第一个数组元素",
			0,
			"I",
		},
		{
			"测试第3个数组元素",
			2,
			"smart",
		},
		{
			"测试第5个数组元素",
			4,
			"strong",
		},
	}

	result := ChangeArray()

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result[tC.index] != tC.expectValue {
				t.Errorf("数组元素不正确：元素为第 %d 个，期待值为： %s, 对象为 %s", tC.index, tC.expectValue, result[tC.index])
			}
		})
	}
}
