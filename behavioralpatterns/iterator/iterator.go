// iterator
package main

type Iterator interface {
	next()
	//previous()
}

type Employee struct {
	employeeName string
	education    string
}

func (employee *Employee) setEmployeeName(name string) {
	employee.employeeName = name
}

func (employee *Employee) getEmployeeName() string {
	return employee.employeeName
}

func (employee *Employee) setEducation(education string) {
	employee.education = education
}

func (employee *Employee) getEducation() string {
	return employee.education
}

type EmployeeCollection struct {
	employeeList []Employee
	employeeMax  int
}

func (employeeCollection *EmployeeCollection) addEmployee(employee Employee) {
	employeeCollection.employeeList = append(employeeCollection.employeeList, employee)
	employeeCollection.employeeMax++
}

func (employeeCollection *EmployeeCollection) getEmployee(i int) Employee {
	return employeeCollection.employeeList[i-1]
}

func (employeeCollection *EmployeeCollection) getEmployeeMax() int {
	return employeeCollection.employeeMax
}

func (employeeCollection *EmployeeCollection) setEmployeeMax(i int) {
	employeeCollection.employeeMax = i
}

type ImplementIterator struct {
	employeeCollection EmployeeCollection
	currentIndex       int
}

func (implementIterator *ImplementIterator) init(employeeCollection EmployeeCollection) {
	implementIterator.employeeCollection = employeeCollection
	implementIterator.currentIndex = employeeCollection.getEmployeeMax()
}

func (implementIterator *ImplementIterator) next() *Employee {

	if implementIterator.currentIndex == 0 {
		return nil
	}

	Em := implementIterator.employeeCollection.getEmployee(implementIterator.currentIndex)
	implementIterator.currentIndex = implementIterator.currentIndex - 1
	return &Em
}

func main() {
	println("——————程序开始运行.————————")

	employee1 := Employee{"小王", "学士"}
	employee2 := Employee{"小张", "学士"}
	employee3 := Employee{"小刘", "硕士"}
	employee4 := Employee{"小李", "学士"}
	employee5 := Employee{"小马", "硕士"}

	employeeCollection := EmployeeCollection{}
	employeeCollection.setEmployeeMax(0)

	employeeCollection.addEmployee(employee1)
	employeeCollection.addEmployee(employee2)
	employeeCollection.addEmployee(employee3)
	employeeCollection.addEmployee(employee4)
	employeeCollection.addEmployee(employee5)

	iterator := ImplementIterator{}
	iterator.init(employeeCollection)

	employee := iterator.next()

	for employee != nil {
		if employee.getEducation() == "硕士" {
			print(employee.getEmployeeName() + ";")
		}
		employee = iterator.next()
	}
	println()

	println("——————程序运行结束.————————")
}
