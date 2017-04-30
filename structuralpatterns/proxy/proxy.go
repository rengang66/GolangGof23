// proxy
package main

type AbstractOrganization interface {
	request()
}

type Corporation struct {
}

func (corporation *Corporation) request() {
	println("这是公司的答复。")
}

type Agency struct {
	corporation *Corporation
}

func (agency *Agency) request() {
	if agency.corporation == nil {
		agency.corporation = &Corporation{}
	}
	agency.corporation.request()
}

func main() {
	println("——————程序开始运行.————————")
	agency := Agency{}
	println("——这是第一次答复————")
	agency.request()
	println("——这是第二次答复————")
	agency.request()
	println("——这是第三次答复————")
	agency.request()

	println("——————程序运行结束.————————")

}
