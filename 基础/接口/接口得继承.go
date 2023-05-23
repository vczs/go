package main


/**
一个接口可以继承其他多个接口（继承的接口中不能有相同的方法），如果A接口继承了B和C接口，那么要实现A接口，也必须要将B和C接口的方法全部实现，缺一不可。
 */
type A interface {
	aa()
	bb()
}
type B interface {
	//aa()  //继承的接口中不能有相同的方法，B中的aa()和A中的aa()相同
	cc()
}
type C interface {
	A
	B //继承的接口中不能有相同的方法，如果有会冲突，编译不能通过
}
