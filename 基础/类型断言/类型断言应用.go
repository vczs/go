package main

import "fmt"

type Student struct{
	name string
}
func typeJudge(variable ...interface{}){
	for index , value := range variable {
		switch value.(type){
		case int,int8,int32,int64 :
			fmt.Printf("第%v个变量%v是\"整数\"类型\n",index+1,value)
		case string :
			fmt.Printf("第%v个变量%v是\"string\"类型\n",index+1,value)
		case float32:
			fmt.Printf("第%v个变量%v是\"float32\"类型\n",index+1,value)
		case float64:
			fmt.Printf("第%v个变量%v是\"float64\"类型\n",index+1,value)
		case bool :
			fmt.Printf("第%v个变量%v是\"布尔\"类型\n",index+1,value)
		case Student :
			fmt.Printf("第%v个变量%v是\"Student\"类型\n",index+1,value)
		case *Student:
			fmt.Printf("第%v个变量%v是\"*Student\"类型\n",index+1,value)
		default:
			fmt.Printf("第%v个变量%v类型无法确定！！！\n",index+1,value)
		}
	}
}
func main(){
	var a int8 = 15
	b := "hello golang!"
	var c = 65.5
	var d bool
	e := 502
	var f float32 = 12.8
	g := Student{"学生g"}
	h := &Student{"学生h"}
	typeJudge(a,b,c,d,e,f,g,h)
}