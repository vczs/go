package model

/**
封装：变量隐藏 只对外提供允许的操作接口  并将变量和操作接口封装在一个类中
 */
type bank struct { //隐藏信息 所以结构体 和 结构体字段的首字母都小写
	account string
	password string
	balance float64
}


/**
工厂模式函数 构造函数
对外提供允许的操作接口 首字母大写
创建一个bank实例
*/
func NewBank(account string,password string) *bank {
	return &bank {
		account : account,
		password :password,
	}
}


/**
对外提供允许的操作接口 首字母大写
修改余额
*/
func (bank *bank) SetBalance(balance float64) {
	bank.balance = balance
}


/**
对外提供允许的操作接口 首字母大写
查询余额
*/
func (bank *bank) GetBalance() float64 {
	return bank.balance
}