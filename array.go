package main

import "errors"

const (
	DefaultCapacity = 16
	CapacityThreshold = 1024
)
//动态数组实现
type Array struct {
	data []interface{}
	capacity int
	size int
}

//实例化
func New(capacity int) *Array  {
	if capacity == 0{
		return nil
	}
	return &Array{
		make([]interface{},capacity,capacity),
		capacity,
		0,
	}
}

func Default() *Array {
	return New(DefaultCapacity)
}

//判断index是否超出范围
func (arr *Array) checkIndex (index int) error{
	if index < 0 || index > arr.capacity{
		return errors.New("index out of range")
	}
	return nil
}

//插入元素
func (arr *Array) Add(index int,val interface{}) error {
	if _err := arr.checkIndex(index);_err != nil {
		return _err
	}
	//判断是否需要扩容
	if arr.size >= arr.capacity{
		if arr.capacity >= CapacityThreshold{
			arr.resize(int(float64(arr.capacity) * 1.25))
		}else{
			arr.resize(arr.capacity << 1)
		}
	}
	copy(arr.data[index + 1 :],arr.data[index : ])
	arr.data[index] = val
	arr.size++
	return nil
}

//扩容函数
func (arr *Array) resize (newCap int) {
	_data := make([]interface{},newCap,newCap)
	copy(_data,arr.data)
	arr.data = _data
	arr.capacity = newCap
}

//插入至尾部
func (arr *Array) Append (val interface{}) error {
	return arr.Add(arr.size,val)
}

//根据索引查找
func (arr *Array) Find (index int) (interface{},error) {
	if _err := arr.checkIndex(index);_err != nil{
		return nil,_err
	}
	return arr.data[index],nil
}

func (arr *Array) Empty() bool {
	return arr.size == 0
}

//是否包含
func (arr *Array) Contains (val interface{}) bool{
	if arr.Empty(){
		return false
	}
	for _,_v := range arr.data{
		if _v == val{
			return true
		}
	}
	return false
}

//删除索引处的值
func (arr *Array) Del (index int) error {
	//因为容量为0实际数据是nil,容量为1，实际元素的index是0
	if _err := arr.checkIndex(index + 1);_err != nil{
		return _err
	}
	if arr.data[index] == nil{
		return nil
	}
	arr.data[index] = nil
	copy(arr.data[index:],arr.data[index + 1:])
	arr.data[arr.size - 1] = nil
	arr.size--
	return nil
}

func (arr *Array) Size() int {
	return arr.size
}
