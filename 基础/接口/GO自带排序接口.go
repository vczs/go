package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//创建一个结构体
type Student struct {
	name string
	age int
}
//创建一个Student结构体类型的切片
type Students []Student
//让该切片实现下面的三种方法，即就是让Students切片实现了sort包中的Sort接口
func (students Students) Len() int{
	return len(students) //返回Students切片的长度
}
func (students Students) Less(i int , j int) bool {
	return students[i].age > students[j].age //返回Students切片的排序规则：i<j是从小到大，i>j是从大到小，==不排序
}
func (students Students) Swap(i int , j int) {
	temp := students[i] //交换切片中下标为i和j的元素位置
	students[i] = students[j]
	students[j] = temp
	//以上三行代码可以简化为：students[i],students[j] = students[j],students[i] 体现了go语言的简洁性
}
func main(){
	var students Students //实例化切片
	//给切片的各个元素赋值
	for i := 0 ; i < 6 ; i++ {
		student := Student{
			name : fmt.Sprintf("学生%d",rand.Intn(100)),
			age : rand.Intn(100),
		}
		students = append(students,student) //将赋值好的元素添加到切片中
	}
	fmt.Println("排序前：")
	for _ , v := range students{ //排序前遍历
		fmt.Println(v)
	}
	//使用go自带包里的sort包中的Sort方法进行排序。
	sort.Sort(students) //参数是一个接口变量，前面Students切片实现了该接口，students又是Students数据类型的，所以可以传入进行排序。
	fmt.Println("排序后：")
	for _ , v := range students{ //排序后遍历
		fmt.Println(v)
	}
}