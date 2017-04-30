// prototype
package main

type AbstractPrototype struct {
	introductionMap map[string]string
}

func (introduction *AbstractPrototype) initialize() {
	introduction.introductionMap = make(map[string]string)
}

type CompanyBaseIntroduction struct {
	AbstractPrototype
}

func (introduction CompanyBaseIntroduction) cloneMyself() CompanyBaseIntroduction {
	introduction.introductionMap["公司介绍"] = "这是公司基本介绍"
	return introduction
}

func (introduction *CompanyBaseIntroduction) addSomeIntrouction(topic string, content string) {
	introduction.introductionMap[topic] = content
}

func (introduction *CompanyBaseIntroduction) deleteSomeIntrouction(topic string) {
	delete(introduction.introductionMap, topic)
}

func (introduction *CompanyBaseIntroduction) showIntrouction() {
	for topic := range introduction.introductionMap {
		println(topic, ":", introduction.introductionMap[topic])
	}
}

func main() {
	println("——————程序开始运行.————————")
	Company := CompanyBaseIntroduction{}
	Company.initialize()

	departA := Company.cloneMyself()
	departA.addSomeIntrouction("部门A介绍", "这是部门A介绍的内容")
	departA.showIntrouction()

	departB := departA.cloneMyself()
	departB.deleteSomeIntrouction("部门A介绍")
	departB.addSomeIntrouction("部门B介绍", "这是部门B介绍的内容")
	departB.showIntrouction()

	println("——————程序运行结束.————————")
}
