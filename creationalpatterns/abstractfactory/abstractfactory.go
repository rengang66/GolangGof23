// abstractfactory
package main

type Computer interface {
	doUse()
}

type Company interface {
	bulidProduct(Parameter string)
}

type Telephone interface {
	doUse()
}

type NotebookComputer struct {
}

type PersonalComputer struct {
}

type DesktopPhone struct {
}

type Mobile struct {
}

type CompanyA struct {
}

func (computer NotebookComputer) doUse() {
	println("这是笔记本电脑的功能")
}

func (computer PersonalComputer) doUse() {
	println("这是个人计算机的功能")
}

func (telephone DesktopPhone) doUse() {
	println("这是座机电话的功能")
}

func (telephone Mobile) doUse() {
	println("这是手机的功能")
}

func (company *CompanyA) bulidComputer(Parameter string) Computer {
	if Parameter == "个人电脑" {
		computer := PersonalComputer{}
		return computer
	} else if Parameter == "笔记本电脑" {
		computer := NotebookComputer{}
		return computer
	} else {
		return nil
	}
}

func (company *CompanyA) bulidTelephone(Parameter string) Telephone {
	if Parameter == "座机电话" {
		telephone := DesktopPhone{}
		return telephone
	} else if Parameter == "手机" {
		telephone := Mobile{}
		return telephone
	} else {
		return nil
	}
}

func main() {
	println("——————程序开始运行.————————")
	//根据传入的参数得到Computer产品
	company := CompanyA{}
	computer1 := company.bulidComputer("个人电脑")
	computer1.doUse()

	computer2 := company.bulidComputer("笔记本电脑")
	computer2.doUse()

	//根据传入的参数得到Telephone产品
	telephone1 := company.bulidTelephone("座机电话")
	telephone1.doUse()

	telephone2 := company.bulidTelephone("手机")
	telephone2.doUse()

	println("——————程序运行结束.————————")
}
