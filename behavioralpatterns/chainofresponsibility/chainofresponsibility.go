// chainofresponsibility
package main

type SuperMan interface {
	setAnswer(answer string)
	answerTheQuestion(man PrimaryMan)
}

type AbstractSuperMan struct {
	answer            string
	successorSuperMan SuperMan
}

func (abstractSuperMan *AbstractSuperMan) getAnswer() string {
	return abstractSuperMan.answer
}

func (abstractSuperMan *AbstractSuperMan) setAnswer(answer string) {
	abstractSuperMan.answer = answer
}

func (abstractSuperMan *AbstractSuperMan) getSuccessorSuperMan() SuperMan {
	return abstractSuperMan.successorSuperMan
}

func (abstractSuperMan *AbstractSuperMan) setSuccessorSuperMan(man SuperMan) {
	abstractSuperMan.successorSuperMan = man
}

type PrimaryMan struct {
	thisAnswer string
}

func (primaryMan *PrimaryMan) setThisAnswer(answer string) {
	primaryMan.thisAnswer = answer
}

func (primaryMan *PrimaryMan) getThisAnswer() string {
	return primaryMan.thisAnswer
}

func (primaryMan *PrimaryMan) question(answer string) bool {
	if primaryMan.thisAnswer == answer {
		return true
	} else {
		return false
	}
}

type SuperManOne struct {
	AbstractSuperMan
	isAnswer bool
}

func (superManOne *SuperManOne) init() {
	superManOne.isAnswer = false
}

func (superManOne *SuperManOne) answerTheQuestion(man PrimaryMan) {
	if man.question(superManOne.answer) {
		println("SuperManOne回答正确。")
	} else {
		if superManOne.isAnswer {
			println("大家都不能回答这个问题。")
		} else {
			superManOne.isAnswer = true
			if superManOne.successorSuperMan != nil {
				println("SuperManOne不会回答问题，提交下一个。")
				superManOne.successorSuperMan.answerTheQuestion(man)
			}
		}
	}
}

type SuperManTwo struct {
	AbstractSuperMan
	isAnswer bool
}

func (superManTwo *SuperManTwo) init() {
	superManTwo.isAnswer = false
}

func (superManTwo *SuperManTwo) answerTheQuestion(man PrimaryMan) {
	if man.question(superManTwo.answer) {
		println("superManTwo回答正确。")
	} else {
		if superManTwo.isAnswer {
			println("大家都不能回答这个问题。")
		} else {
			superManTwo.isAnswer = true
			if superManTwo.successorSuperMan != nil {
				println("superManTwo不会回答问题，提交下一个。")
				superManTwo.successorSuperMan.answerTheQuestion(man)
			}
		}
	}
}

type SuperManThree struct {
	AbstractSuperMan
	isAnswer bool
}

func (superManTwo *SuperManThree) init() {
	superManTwo.isAnswer = false
}

func (superManTwo *SuperManThree) answerTheQuestion(man PrimaryMan) {
	if man.question(superManTwo.answer) {
		println("SuperManThree回答正确。")
	} else {
		if superManTwo.isAnswer {
			println("大家都不能回答这个问题。")
		} else {
			superManTwo.isAnswer = true
			if superManTwo.successorSuperMan != nil {
				println("SuperManThree不会回答问题，提交下一个。")
				superManTwo.successorSuperMan.answerTheQuestion(man)
			}
		}
	}
}

func main() {
	println("——————程序开始运行.————————")
	superMan1 := &SuperManOne{}
	superMan1.init()
	superMan2 := &SuperManTwo{}
	superMan2.init()
	superMan3 := &SuperManThree{}
	superMan3.init()

	// 形成一个闭环的链
	superMan1.setSuccessorSuperMan(superMan2)
	superMan2.setSuccessorSuperMan(superMan3)
	superMan3.setSuccessorSuperMan(superMan1)

	// 设置提问者的问题和答案
	primaryMan := PrimaryMan{}
	primaryMan.setThisAnswer("ANSWER")

	// 让superMan1能回答这个答案
	superMan1.setAnswer("ANSWER1")

	// 让superMan2能回答这个答案
	superMan2.setAnswer("ANSWER2")

	// 让superMan3能回答这个答案
	superMan3.setAnswer("ANSWER")
	superMan1.answerTheQuestion(primaryMan)

	println("——————程序运行结束.————————")
}
