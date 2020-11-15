# two_sum_golang

## 題目解讀

### 題目來源

[two_sum](https://leetcode.com/problems/two-sum/)

### 原文:

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

### 解讀:

給定一個整數的陣列 nums 還有一個整數 target

寫一個 function  來找出陣列中 兩個元素 

符合 兩個元素的和 等於 target

然後返回這兩個元素的index

## 初步解法:
### 初步觀察:
首先是 假設 array nums中 存在兩個 index i, j 使得 target = nums[i] + nums[j]

則代表 i跟j存在以下關係

nums[j] = target - num[i]

也就是只要先取出一個值 假設是 k = nums[p]

然後逐步測試 其他 元素是否等於 target - k

就可以找到目標
### 初步設計:
給定整數 array nums 給定一整數 target

先建立一個 整數map , valueMap

step 1: let index = 0 , len = len(nums)

step 2: valueMap[nums[index]]=index

step 3: check valueMap[target-nums[index]] exists if exists return index, valueMap[target-nums[index]] else go to step 4

step 4: index = index + 1  if index <  len  go to step 1, else return []

## 遇到的困難

### 驗證的部份

雖然說題目上說是不用管index 順序

但為了檢驗方便

還是要做一個簡單排序 由小至大

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

## 我的github source code
```golang
package two_sum_golang

func twoSum(nums []int, target int) []int {
	results := []int{}
	valueMap := make(map[int]int)
	for idx, val := range nums {
		value, ok := valueMap[target-nums[idx]]
		if ok && value != idx {
			if idx < value {
				return []int{idx, value}
			} else {
				return []int{value, idx}
			}
		} else {
			valueMap[val] = idx
		}
	}
	return results
}
```
## 測資的撰寫
```golang
package two_sum_golang

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Example2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Example3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

```