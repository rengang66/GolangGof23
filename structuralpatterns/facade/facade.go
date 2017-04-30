// facade
package main

type DoWork struct {
}

func (act *DoWork) operate() {
	println("员工：打卡考勤。")
}

type Inspection struct {
}

func (act *Inspection) operate() {
	println("领导视察：端茶倒水。")
}

type Post struct {
}

func (act *Post) operate() {
	println("邮递员：登记收发物品。")
}

type Visit struct {
}

func (act *Visit) operate() {
	println("访客：登记身份证。")
}

type Facade struct {
	visit      Visit
	post       Post
	inspection Inspection
	dowork     DoWork
}

func (facade *Facade) init() {
	facade.visit = Visit{}
	facade.post = Post{}
	facade.inspection = Inspection{}
	facade.dowork = DoWork{}
}

func (facade *Facade) Operate(operation string) {
	if operation == "visit" {
		facade.visit.operate()
	} else if operation == "post" {
		facade.post.operate()
	} else if operation == "inspection" {
		facade.inspection.operate()
	} else if operation == "doWork" {
		facade.dowork.operate()
	} else {
		println("没有对应事项，不能工作。")
	}
}

func main() {

	println("——————程序开始运行.————————")
	facade := Facade{}
	facade.init()

	// 向前台要求访客
	facade.Operate("visit")

	// 向前台提交邮品
	facade.Operate("post")

	// 领导过来视察
	facade.Operate("inspection")

	// 员工上班
	facade.Operate("doWork")

	// study是没有对应的工作接口
	facade.Operate("study")

	println("——————程序运行结束.————————")

}
