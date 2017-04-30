// interpreter
package main

type AbstractExpression interface {
	Interpret(project Project)
}

type Project struct {
	projectName string
}

func (project *Project) setProjectName(projectName string) {
	project.projectName = projectName
}

func (project *Project) getProjectName() string {
	return project.projectName
}

type FinancialDepExpression struct {
}

func (financialDepExpression *FinancialDepExpression) Interpret(project Project) {
	println("财务部对" + project.getProjectName() + "的理解。")
}

type MarketDepExpression struct {
}

func (marketDepExpression *MarketDepExpression) Interpret(project Project) {
	println("市场部对" + project.getProjectName() + "的理解。")
}

type TechnicalDepExpression struct {
}

func (technicalDepExpression *TechnicalDepExpression) Interpret(project Project) {
	println("技术部对" + project.getProjectName() + "的理解。")
}

func main() {
	println("——————程序开始运行.————————")
	project1 := Project{}
	project1.setProjectName("ProjectA")
	var expressList []AbstractExpression

	expressList = append(expressList, &FinancialDepExpression{})
	expressList = append(expressList, &MarketDepExpression{})
	expressList = append(expressList, &TechnicalDepExpression{})

	for _, val := range expressList {
		val.Interpret(project1)
	}

	println("——————程序运行结束.————————")
}
