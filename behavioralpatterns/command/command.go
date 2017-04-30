// command
package main

type Command interface {
	setMan(man Subordinate)
	Execute()
}

type Manager struct {
	command Command
}

func (manager *Manager) setCommand(command Command) {
	manager.command = command
}

func (manager *Manager) ExecuteCommand() {
	manager.command.Execute()
}

type Subordinate struct {
	name string
}

func (subordinate *Subordinate) setName(name string) {
	subordinate.name = name
}

func (subordinate *Subordinate) getName() string {
	return subordinate.name
}

func (subordinate *Subordinate) executeCommand(task string) {
	println(subordinate.name + "执行" + task + "的命令")
}

type PlanCommand struct {
	man Subordinate
}

func (planCommand *PlanCommand) setMan(man Subordinate) {
	planCommand.man = man
}

func (planCommand *PlanCommand) Execute() {
	println("———编写计划———")
	planCommand.man.executeCommand("编写计划")
}

type ReportCommand struct {
	man Subordinate
}

func (reportCommand *ReportCommand) setMan(man Subordinate) {
	reportCommand.man = man
}

func (reportCommand *ReportCommand) Execute() {
	println("———编写计划———")
	reportCommand.man.executeCommand("编写报告")
}

func main() {
	println("——————程序开始运行.————————")
	subordinate := Subordinate{}
	subordinate.setName("小刘")

	command1 := &ReportCommand{}
	command1.setMan(subordinate)
	manager := Manager{}
	manager.setCommand(command1)
	manager.ExecuteCommand()

	command2 := &PlanCommand{}
	command2.setMan(subordinate)
	manager.setCommand(command2)
	manager.ExecuteCommand()

	println("——————程序运行结束.————————")

}
