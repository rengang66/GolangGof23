// adapter
package main

type Customer struct {
	customerName string
	requirement  string
}

type Designer struct {
	designerName string
	works        string
	Designs      string
}

type Analyst struct {
	Customer
	analystName string
	designer    Designer
	request     string
}

func (customer *Customer) setCustomerName(customerName string) {
	customer.customerName = customerName
}

func (customer *Customer) getCustomerName() string {
	return customer.customerName
}

func (customer *Customer) setRequirement(requirement string) {
	customer.requirement = requirement
}

func (customer *Customer) getRequirement() string {
	return customer.requirement
}

func (customer *Customer) commitRequirement() string {
	return customer.getRequirement()
}

func (designer *Designer) setName(designerName string) {
	designer.designerName = designerName
}

func (designer *Designer) getName() string {
	return designer.designerName
}

func (designer *Designer) getWorks() string {
	return designer.works
}

func (designer *Designer) setDesigns(designs string) {
	designer.Designs = designs
}

func (designer *Designer) getFinishWork() string {
	designer.designToWorks()
	return designer.works
}

func (designer *Designer) designToWorks() {
	println("———按照需求分析设计转化为具体工作———")
	designer.works = designer.Designs + designer.getName() + "按照需求设计内容，完成工作。"
}

func (analyst *Analyst) getName() string {
	return analyst.analystName
}

func (analyst *Analyst) setName(analystName string) {
	analyst.analystName = analystName
}

func (analyst *Analyst) getDesigner() Designer {
	return analyst.designer
}

func (analyst *Analyst) setDesigner(designer Designer) {
	analyst.designer = designer
}

func (analyst *Analyst) setRequest(request string) {
	analyst.request = request
}

func (analyst *Analyst) getFinishworks() string {
	designs := analyst.requestToDesign()
	analyst.designer.setDesigns(designs)
	return analyst.designer.getFinishWork()
}

func (analyst *Analyst) requestToDesign() string {
	println("———分析员按照用户需求转化为需求分析和设计———")
	return analyst.getName() + "按照" + analyst.request + "，形成需求设计内容。"
}

func main() {
	println("——————程序开始运行.————————")

	customer := Customer{}
	customer.setCustomerName("客户小王")
	customer.setRequirement("客户小王的需求")

	designer := Designer{"开发员小张", "", ""}

	analyst := Analyst{}
	analyst.setName("分析员小刘")

	//客户把用户需求提交给分析员
	analyst.setRequest(customer.commitRequirement())

	//分析员经过转化后提供给开发人员
	analyst.setDesigner(designer)

	//得到满足客户的需求的工作产品
	println(analyst.getFinishworks())

	println("——————程序运行结束.————————")
}
