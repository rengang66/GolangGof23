// memento
package main

type Memento struct {
	meetName         string
	meetType         string
	meetdate         string
	meetLeader       string
	meetParticipants string
	decide           string
	meetContent      string
}

func (memento *Memento) initialize(meetName string, meetType string, meetdate string,
	meetLeader string, meetParticipants string, decide string, meetContent string) {
	memento.meetName = meetName
	memento.meetType = meetType
	memento.meetdate = meetdate
	memento.meetLeader = meetLeader
	memento.meetParticipants = meetParticipants
	memento.decide = decide
	memento.meetContent = meetContent
}

func (memento *Memento) getMeetName() string {
	return memento.meetName
}

func (memento *Memento) setMeetName(meetName string) {
	memento.meetName = meetName
}

func (memento *Memento) getMeetType() string {
	return memento.meetType
}

func (memento *Memento) getMeetdate() string {
	return memento.meetdate
}

func (memento *Memento) getMeetLeader() string {
	return memento.meetLeader
}

func (memento *Memento) getMeetParticipants() string {
	return memento.meetParticipants
}

func (memento *Memento) getDecide() string {
	return memento.decide
}

func (memento *Memento) getMeetContent() string {
	return memento.meetContent
}

func (memento *Memento) showContent() {
	println(memento.meetContent)
}

type Meeting struct {
	meetName         string
	meetType         string
	meetdate         string
	meetLeader       string
	meetParticipants string
	decide           string
	meetContent      string
}

func (meeting *Meeting) getMeetName() string {
	return meeting.meetName
}

func (meeting *Meeting) setMeetName(meetName string) {
	meeting.meetName = meetName
}

func (meeting *Meeting) getMeetType() string {
	return meeting.meetType
}

func (meeting *Meeting) setMeetType(meetType string) {
	meeting.meetType = meetType
}

func (meeting *Meeting) getMeetdate() string {
	return meeting.meetdate
}

func (meeting *Meeting) setMeetdate(meetdate string) {
	meeting.meetdate = meetdate
}

func (meeting *Meeting) getMeetLeader() string {
	return meeting.meetLeader
}

func (meeting *Meeting) setMeetLeader(meetLeader string) {
	meeting.meetLeader = meetLeader
}

func (meeting *Meeting) getMeetParticipants() string {
	return meeting.meetParticipants
}

func (meeting *Meeting) setMeetParticipants(meetParticipants string) {
	meeting.meetParticipants = meetParticipants
}

func (meeting *Meeting) getDecide() string {
	return meeting.decide
}

func (meeting *Meeting) setDecide(decide string) {
	meeting.decide = decide
}

func (meeting *Meeting) getMeetContent() string {
	return meeting.meetContent
}

func (meeting *Meeting) setMeetContent(meetContent string) {
	meeting.meetContent = meetContent
}

func (meeting *Meeting) showContent() {
	println(meeting.meetContent)
}

func (meeting *Meeting) doDecide() {
	meeting.meetContent = "会议名称：" + meeting.meetName + ";"
	meeting.meetContent = meeting.meetContent + "会议类型：" + meeting.meetType + ";"
	meeting.meetContent = meeting.meetContent + "会议时间：" + meeting.meetdate + ";"
	meeting.meetContent = meeting.meetContent + "会议主持人：" + meeting.meetLeader + ";"
	meeting.meetContent = meeting.meetContent + "会议参加者：" + meeting.meetParticipants + "\r"
	meeting.meetContent = meeting.meetContent + "会议内容：" + meeting.meetParticipants + ";"
}

func (meeting *Meeting) CreateMemento() Memento {
	mo := Memento{meeting.meetName, meeting.meetType, meeting.meetdate, meeting.meetLeader,
		meeting.meetParticipants, meeting.decide, meeting.meetContent}
	return mo
}

func (meeting *Meeting) recoveryMemento(mo Memento) {
	meeting.meetName = mo.getMeetName()
	meeting.meetType = mo.getMeetType()
	meeting.meetdate = mo.getMeetdate()
	meeting.meetLeader = mo.getMeetLeader()
	meeting.meetParticipants = mo.getMeetParticipants()
	meeting.decide = mo.getDecide()
	meeting.meetContent = mo.getMeetContent()
}

type Caretaker struct {
	meetMementoMap map[string]Memento
}

func (caretaker *Caretaker) init() {
	caretaker.meetMementoMap = make(map[string]Memento)
}

func (caretaker *Caretaker) recoveryMemento(meetdate string) Memento {
	return caretaker.meetMementoMap[meetdate]
}

func (caretaker *Caretaker) saveMemento(meetdate string, memento Memento) {
	caretaker.meetMementoMap[meetdate] = memento
}

// 展示内容
func (caretaker *Caretaker) showContent() {
	println("显示全部的会议情况：")
	for topic := range caretaker.meetMementoMap {
		mo := caretaker.meetMementoMap[topic]
		println("会议时间是:" + topic + " 会议内容是:" + mo.getMeetContent())
	}
}

func main() {
	println("——————程序开始运行.————————")

	meet1 := Meeting{}
	meet1.setMeetName("meetName1")
	meet1.setMeetType("meetType1")
	meet1.setMeetdate("2009-01-01")
	meet1.setMeetLeader("meetLeader1")
	meet1.setMeetParticipants("meetParticipants1")
	meet1.setDecide("decide1")
	meet1.doDecide()
	memento := meet1.CreateMemento()

	caretaker := Caretaker{}
	caretaker.init()
	caretaker.saveMemento(memento.getMeetdate(), memento)

	meet2 := Meeting{}
	meet2.setMeetName("meetName2")
	meet2.setMeetType("meetType2")
	meet2.setMeetdate("2009-02-01")
	meet2.setMeetLeader("meetLeader2")
	meet2.setMeetParticipants("meetParticipants2")
	meet2.setDecide("decide2")
	meet2.doDecide()

	memento = meet2.CreateMemento()

	caretaker.saveMemento(memento.getMeetdate(), memento)

	caretaker.showContent()

	println("——————程序运行结束.————————")
}
