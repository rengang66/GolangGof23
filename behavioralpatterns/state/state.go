// state
package main

type State interface {
	doWork(project Project)
}

type Project struct {
	projectName  string
	currentState State
}

func (project *Project) getProjectName() string {
	return project.projectName
}

func (project *Project) setProjectName(projectName string) {
	project.projectName = projectName
}

func (project *Project) getCurrentState() State {
	return project.currentState
}

func (project *Project) setCurrentState(state State) {
	project.currentState = state
}

func (project *Project) doCurrentWork() {
	project.currentState.doWork(*project)
}

type ProjectBuilderState struct {
}

func (projectBuilderState *ProjectBuilderState) doWork(project Project) {
	println(project.getProjectName() + "正在进行立项工作")
}

type ProjectDevelopmentState struct {
}

func (projectDevelopmentState *ProjectDevelopmentState) doWork(project Project) {
	println(project.getProjectName() + "正在进行开发工作")
}

type ProjectEndState struct {
}

func (projectEndState *ProjectEndState) doWork(project Project) {
	println(project.getProjectName() + "正在进行结项工作")
}

type ProjectMaintenanceState struct {
}

func (projectMaintenanceState *ProjectMaintenanceState) doWork(project Project) {
	println(project.getProjectName() + "正在进行维护工作")
}

type ProjectRunState struct {
}

func (projectRunState *ProjectRunState) doWork(project Project) {
	println(project.getProjectName() + "正在进行试运行工作")
}

func main() {
	println("——————程序开始运行.————————")

	bullder := &ProjectBuilderState{}
	development := &ProjectDevelopmentState{}
	run := &ProjectRunState{}
	maintenance := &ProjectMaintenanceState{}
	end := &ProjectEndState{}

	projectA := Project{}
	projectA.setProjectName("projectA")

	println("——设置项目为立项状态———")
	projectA.setCurrentState(bullder)
	projectA.doCurrentWork()

	println("——设置项目为开发状态———")
	projectA.setCurrentState(development)
	projectA.doCurrentWork()

	println("——设置项目为试运行状态———")
	projectA.setCurrentState(run)
	projectA.doCurrentWork()

	println("——设置项目为维护状态———")
	projectA.setCurrentState(maintenance)
	projectA.doCurrentWork()

	println("——设置项目为结项状态———")
	projectA.setCurrentState(end)
	projectA.doCurrentWork()

	println("——————程序运行结束.————————")
}
