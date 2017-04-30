// observer
package main

type AbstractAddressBook struct {
	employeeList []CompanyEmployee
	addressBook  string
}

func (abstractAddressBook *AbstractAddressBook) getAddressBook() string {
	return abstractAddressBook.addressBook
}

func (abstractAddressBook *AbstractAddressBook) setAddressBook(book string) {
	abstractAddressBook.addressBook = book
}

func (abstractAddressBook *AbstractAddressBook) abstractAddEmployee(employee CompanyEmployee) {
	abstractAddressBook.employeeList = append(abstractAddressBook.employeeList, employee)
}

func (abstractAddressBook *AbstractAddressBook) removeEmployee(employee CompanyEmployee) {
	//abstractAddressBook.employeeList = append(abstractAddressBook.employeeList,employee)
}

type AbstractEmployee struct {
	addressBook         CompanyAddressBook
	addressBookContents string
}

func (abstractEmployee *AbstractEmployee) getAddressBook() CompanyAddressBook {
	return abstractEmployee.addressBook
}

func (abstractEmployee *AbstractEmployee) setAddressBook(book CompanyAddressBook) {
	abstractEmployee.addressBook = book
}

func (abstractEmployee *AbstractEmployee) abstractUpdate(book CompanyAddressBook) {
	abstractEmployee.setAddressBook(book)
	abstractEmployee.addressBookContents = abstractEmployee.addressBook.getAddressBook()
}

type CompanyAddressBook struct {
	AbstractAddressBook
}

func (companyAddressBook *CompanyAddressBook) init() {
	companyAddressBook.addressBook = "旧的通信录"
}

func (companyAddressBook *CompanyAddressBook) addEmployee(employee CompanyEmployee) {
	employee.update(*companyAddressBook)
	companyAddressBook.abstractAddEmployee(employee)
}

func (companyAddressBook *CompanyAddressBook) notice() {
	for _, val := range companyAddressBook.AbstractAddressBook.employeeList {
		val.update(*companyAddressBook)
	}
}

type CompanyEmployee struct {
	AbstractEmployee
	employeeName string
}

func (companyEmployee *CompanyEmployee) init(employeeName string) {
	companyEmployee.employeeName = employeeName
}

func (companyEmployee *CompanyEmployee) setEmployeeName(employeeName string) {
	companyEmployee.employeeName = employeeName
}

func (companyEmployee *CompanyEmployee) getEmployeeName() string {
	return companyEmployee.employeeName
}

func (companyEmployee *CompanyEmployee) update(book CompanyAddressBook) {
	companyEmployee.abstractUpdate(book)
	println(companyEmployee.getEmployeeName() + "更新通讯录。" + "通讯录内容：" + companyEmployee.addressBookContents)
}

func main() {
	println("——————程序开始运行.————————")
	addressBook := CompanyAddressBook{}
	addressBook.init()
	employee1 := CompanyEmployee{}
	employee1.init("employee1")
	employee2 := CompanyEmployee{}
	employee2.init("employee2")

	addressBook.addEmployee(employee1)
	addressBook.addEmployee(employee2)

	println("———更新的通讯录———")
	newAddressBook := "新的通讯录"
	addressBook.setAddressBook(newAddressBook)

	addressBook.notice()

	println("——————程序运行结束.————————")

}
