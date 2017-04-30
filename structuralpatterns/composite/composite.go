// composite
package main

type Organization interface {
	init(name string)
	showOrganizationStructure(parentName string)
}

type AbstractOrganization struct {
	organizationName string
	organizationType string
}

func (abstractOrganization *AbstractOrganization) getOrganizationName() string {
	return abstractOrganization.organizationName
}

func (abstractOrganization *AbstractOrganization) setOrganizationName(name string) {
	abstractOrganization.organizationName = name
}

type Corporation struct {
	AbstractOrganization
	componentList []Organization
}

func (organization *Corporation) init(name string) {
	organization.organizationName = name
	organization.organizationType = "Corporation"
}

func (organization *Corporation) add(org Organization) {
	organization.componentList = append(organization.componentList, org)
}

func (organization *Corporation) showOrganizationStructure(parentName string) {
	organization.organizationName = organization.getOrganizationName()
	println(parentName + organization.organizationName)
	for _, v := range organization.componentList {
		v.showOrganizationStructure(parentName + organization.organizationName + "--")
	}
}

type Department struct {
	AbstractOrganization
	componentList []Organization
}

func (organization *Department) init(name string) {
	organization.organizationName = name
	organization.organizationType = "Department"
}

func (organization *Department) add(org Organization) {
	organization.componentList = append(organization.componentList, org)
}

func (organization *Department) showOrganizationStructure(parentName string) {
	organization.organizationName = organization.getOrganizationName()
	println(parentName + organization.organizationName)
	for _, v := range organization.componentList {
		v.showOrganizationStructure(parentName + organization.organizationName + "--")
	}
}

type Factory struct {
	AbstractOrganization
	componentList []Organization
}

func (organization *Factory) init(name string) {
	organization.organizationName = name
	organization.organizationType = "Factory"
}

func (organization *Factory) add(org Organization) {
	organization.componentList = append(organization.componentList, org)
}

func (organization *Factory) showOrganizationStructure(parentName string) {
	organization.organizationName = organization.getOrganizationName()
	println(parentName + organization.organizationName)
	for _, v := range organization.componentList {
		v.showOrganizationStructure(parentName + organization.organizationName + "--")
	}
}

type WorkTeam struct {
	AbstractOrganization
	componentList []Organization
}

func (organization *WorkTeam) init(name string) {
	organization.organizationName = name
	organization.organizationType = "WorkTeam"
}

func (organization *WorkTeam) add(org Organization) {
	println("这是叶子节点，下面没有内容")
}

func (organization *WorkTeam) showOrganizationStructure(parentName string) {
	organization.organizationName = organization.getOrganizationName()
	println(parentName + organization.organizationName)
}

type Workshop struct {
	AbstractOrganization
	componentList []Organization
}

func (organization *Workshop) init(name string) {
	organization.organizationName = name
	organization.organizationType = "Workshop"
}

func (organization *Workshop) add(org Organization) {
	println("这是叶子节点，下面没有内容")
}

func (organization *Workshop) showOrganizationStructure(parentName string) {
	organization.organizationName = organization.getOrganizationName()
	println(parentName + organization.organizationName)
}

func scene1() {
	// 场景1，形成公司-部门-项目组的公司结构
	corporation := Corporation{}
	corporation.init("公司A")

	depart1 := &Department{}
	depart1.init("部门1")

	depart2 := &Department{}
	depart2.init("部门2")

	depart3 := &Department{}
	depart3.init("部门3")

	corporation.add(depart1)
	corporation.add(depart2)
	corporation.add(depart3)

	team1 := &WorkTeam{}
	team1.init("项目组1")
	team2 := &WorkTeam{}
	team2.init("项目组2")

	depart1.add(team1)
	depart1.add(team2)

	corporation.showOrganizationStructure("")
}

func scene2() {
	// 场景2，形成公司-工厂-车间的工厂结构
	corporation := &Corporation{}
	corporation.init("公司A")

	factory1 := &Factory{}
	factory1.init("工厂1")

	factory2 := &Factory{}
	factory2.init("工厂2")

	factory3 := &Factory{}
	factory3.init("工厂3")

	corporation.add(factory1)
	corporation.add(factory2)
	corporation.add(factory3)

	team1 := &Workshop{}
	team1.init("车间1")
	team2 := &Workshop{}
	team2.init("车间2")

	factory1.add(team1)
	factory1.add(team2)

	corporation.showOrganizationStructure("")
}

func main() {
	println("——————程序开始运行.————————")
	scene1()
	println()
	scene2()

	println("——————程序运行结束.————————")
}
