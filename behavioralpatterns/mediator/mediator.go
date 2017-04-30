// mediator
package main

type Task struct {
	taskName string
}

func (task *Task) initialize(taskName string) {
	task.taskName = taskName
}

func (task *Task) setTaskName(taskName string) {
	task.taskName = taskName
}

func (task *Task) getTaskName() string {
	return task.taskName
}

type Member struct {
	memberName string
}

func (member *Member) initialize(memberName string) {
	member.memberName = memberName
}

func (member *Member) setMemberName(memberName string) {
	member.memberName = memberName
}

func (member *Member) getMemberName() string {
	return member.memberName
}

type Project interface {
	removeMember(member Member)
	addMember(member Member)
	reduceTask(task Task)
	addTask(task Task)
	showProjectContent()
}

type AbstractProject struct {
	projectName string
	leader      string
	taskMap     map[string]Task
	memberMap   map[string]Member
}

func (abstractProject *AbstractProject) init() {
	abstractProject.memberMap = make(map[string]Member)
	abstractProject.taskMap = make(map[string]Task)
}

func (abstractProject *AbstractProject) getProjectName() string {
	return abstractProject.projectName
}

func (abstractProject *AbstractProject) setProjectName(projectName string) {
	abstractProject.projectName = projectName
}

func (abstractProject *AbstractProject) getLeader() string {
	return abstractProject.leader
}

func (abstractProject *AbstractProject) setLeader(leader string) {
	abstractProject.leader = leader
}

func (abstractProject *AbstractProject) initizeTask(taskMap map[string]Task) {
	abstractProject.taskMap = taskMap
}

func (abstractProject *AbstractProject) initizeMember(memberMap map[string]Member) {
	abstractProject.memberMap = memberMap
}

func (abstractProject *AbstractProject) addMember(member Member) {
	abstractProject.memberMap[member.getMemberName()] = member
}

func (abstractProject *AbstractProject) removeMember(member Member) {
	delete(abstractProject.memberMap, member.getMemberName())
}

func (abstractProject *AbstractProject) addTask(task Task) {
	abstractProject.taskMap[task.getTaskName()] = task
}

func (abstractProject *AbstractProject) reduceTask(task Task) {
	delete(abstractProject.taskMap, task.getTaskName())
}

func (abstractProject *AbstractProject) showProjectContent() {
	println("显示" + abstractProject.getProjectName() + "项目成员和任务情况：")
	memberName := ""

	for topic := range abstractProject.memberMap {
		memberName = memberName + topic + ";"
	}
	print("项目成员:" + memberName)

	var taskName = ""
	for topic := range abstractProject.taskMap {
		taskName = taskName + topic + ";"
	}
	println("项目任务:" + taskName)
}

type Mediator struct {
}

func (mediator *Mediator) changMember(project1 Project, project2 Project, member Member) {
	project1.removeMember(member)
	project2.addMember(member)
}

func (mediator *Mediator) changTask(project1 Project, project2 Project, task Task) {
	project1.reduceTask(task)
	project2.addTask(task)
}

type ProjectA struct {
	AbstractProject
}

func (projectA *ProjectA) init(members map[string]Member, tasks map[string]Task) {
	projectA.initizeMember(members)
	projectA.initizeTask(tasks)
	projectA.projectName = "ProjectA"
}

type ProjectB struct {
	AbstractProject
}

func (projectB *ProjectB) init(members map[string]Member, tasks map[string]Task) {
	projectB.initizeMember(members)
	projectB.initizeTask(tasks)
	projectB.projectName = "ProjectB"
}

type TechnicalDirector struct {
	Mediator
	projectA     Project
	projectB     Project
	directorName string
}

func (technicalDirector *TechnicalDirector) getDirectorName() string {
	return technicalDirector.directorName
}

func (technicalDirector *TechnicalDirector) setDirectorName(directorName string) {
	technicalDirector.directorName = directorName
}

func (technicalDirector *TechnicalDirector) getProjectA() Project {
	return technicalDirector.projectA
}

func (technicalDirector *TechnicalDirector) setProjectA(project Project) {
	technicalDirector.projectA = project
}

func (technicalDirector *TechnicalDirector) getProjectB() Project {
	return technicalDirector.projectB
}

func (technicalDirector *TechnicalDirector) setProjectB(project Project) {
	technicalDirector.projectB = project
}

func main() {
	println("——————程序开始运行.————————")

	// 初始化项目人员和任务信息
	Programmer1 := Member{"程序员1"}
	Programmer2 := Member{"程序员2"}
	Designer1 := Member{"设计师1"}
	Designer2 := Member{"设计师2"}

	ProgramTask1 := Task{"编程工作1"}
	ProgramTask2 := Task{"编程工作2"}
	DesignTask1 := Task{"设计工作1"}
	DesignTask2 := Task{"设计工作2"}

	var ProgrammerMap map[string]Member
	ProgrammerMap = make(map[string]Member)

	ProgrammerMap[Programmer1.getMemberName()] = Programmer1
	ProgrammerMap[Programmer2.getMemberName()] = Programmer2

	var DesignerMap map[string]Member
	DesignerMap = make(map[string]Member)

	DesignerMap[Designer1.getMemberName()] = Designer1
	DesignerMap[Designer2.getMemberName()] = Designer2

	var ProgramTaskMap map[string]Task
	ProgramTaskMap = make(map[string]Task)

	ProgramTaskMap[ProgramTask1.getTaskName()] = ProgramTask1
	ProgramTaskMap[ProgramTask2.getTaskName()] = ProgramTask2

	var DesignTaskMap map[string]Task
	DesignTaskMap = make(map[string]Task)

	DesignTaskMap[DesignTask1.getTaskName()] = DesignTask1
	DesignTaskMap[DesignTask2.getTaskName()] = DesignTask2

	projectA := &ProjectA{}
	projectA.init(ProgrammerMap, ProgramTaskMap)

	projectB := &ProjectB{}
	projectB.init(DesignerMap, DesignTaskMap)

	// 进行项目人员和工作的协调
	println("———协调前的情况———")
	projectA.showProjectContent()
	projectB.showProjectContent()

	leader := TechnicalDirector{}
	leader.setProjectA(projectA)
	leader.setProjectB(projectB)

	// 协调两个项目的人员
	leader.changMember(projectA, projectB, Programmer1)
	leader.changMember(projectB, projectA, Designer1)

	// 协调两个项目的任务
	leader.changTask(projectA, projectB, ProgramTask1)
	leader.changTask(projectB, projectA, DesignTask1)

	println("———协调后的情况———")
	projectA.showProjectContent()
	projectB.showProjectContent()

	println("——————程序运行结束.————————")
}
