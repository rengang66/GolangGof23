// bridge
package main

/*
type Action interface {
	doAction(depart string, title string)
}
*/

type AbstractAction interface {
	doAction(depart string, title string)
}

/*
func (action AbstractAction) doAction(depart string, title string) {
	println("这是" + depart + "部门的标准工作活动," + "主题是" + title)
}
*/

type AbstractDepartment struct {
	departAction AbstractAction
}

func (department *AbstractDepartment) setAbstractAction(departAction AbstractAction) {
	department.departAction = departAction
}

func (department *AbstractDepartment) action(title string) {
	department.departAction.doAction("", title)
}

type DevelopmentDep struct {
	AbstractDepartment
	departName string
}

func (department *DevelopmentDep) initialize() {
	department.departName = "开发部"
}

func (department *DevelopmentDep) action(title string) {
	department.departAction.doAction(department.departName, title)
}

type FinanceDep struct {
	AbstractDepartment
	departName string
}

func (department *FinanceDep) initialize() {
	department.departName = "财务部"
}

func (department *FinanceDep) action(title string) {
	department.departAction.doAction(department.departName, title)
}

type MarketDep struct {
	AbstractDepartment
	departName string
}

func (department *MarketDep) initialize() {
	department.departName = "市场部"
}

func (department *MarketDep) action(title string) {
	department.departAction.doAction(department.departName, title)
}

type Meeting struct {
	//AbstractAction
}

func (meeting Meeting) doAction(depart string, title string) {
	println("这是" + depart + "会议工作活动," + "主题是" + title)
}

type Training struct {
	//AbstractAction
}

func (training Training) doAction(depart string, title string) {
	println("这是" + depart + "培训工作活动," + "主题是" + title)
}

func scene1() {
	// 场景1：针对开发部的培训工作
	//action := AbstractAction(Training{})
	action := Training{}
	depart := DevelopmentDep{}
	depart.setAbstractAction(action)
	depart.action("提高开发技能")
}

func scene2() {
	// 场景2：针对财务部的会议
	action := Meeting{}
	depart := FinanceDep{}
	depart.setAbstractAction(action)
	depart.action("检查会计制度")
}

func scene3() {
	//场景3：针对市场部的培训
	action := Training{}
	depart := MarketDep{}
	depart.setAbstractAction(action)
	depart.action("沟通技巧")
}

func main() {
	println("——————程序开始运行.————————")

	scene1()
	scene2()
	scene3()

	println("——————程序运行结束.————————")
}
