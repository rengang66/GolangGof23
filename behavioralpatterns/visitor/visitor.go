// visitor
package main

type Visitor interface {
	VisitCompanyA(company CompanyA)
	VisitCompanyB(company CompanyB)
}

type AbstractCompany struct {
	vistor Visitor
}

func (abstractCompany *AbstractCompany) Accept(vistor Visitor) {
	abstractCompany.vistor = vistor
}

func (abstractCompany *AbstractCompany) doVisit() {
	//abstractCompany.vistor = vistor
}

type CompanyA struct {
	AbstractCompany
}

func (companyA *CompanyA) doVisit() {
	companyA.vistor.VisitCompanyA(*companyA)
}

type CompanyB struct {
	AbstractCompany
}

func (companyB *CompanyB) doVisit() {
	companyB.vistor.VisitCompanyB(*companyB)
}

type AccountingFirm struct {
}

func (accountingFirm *AccountingFirm) VisitCompanyA(company CompanyA) {
	println("对CompanyA类公司进行会计审计工作")
}

func (accountingFirm *AccountingFirm) VisitCompanyB(company CompanyB) {
	println("对CompanyB类公司进行会计审计工作")
}

type TaxBureau struct {
}

func (taxBureau *TaxBureau) VisitCompanyA(company CompanyA) {
	println("对CompanyA类公司进行税务审计工作")
}

func (taxBureau *TaxBureau) VisitCompanyB(company CompanyB) {
	println("对CompanyB类公司进行税务审计工作")
}

type TradeBureau struct {
}

func (tradeBureau *TradeBureau) VisitCompanyA(company CompanyA) {
	println("对CompanyA类公司进行工商年审工作")
}

func (tradeBureau *TradeBureau) VisitCompanyB(company CompanyB) {
	println("对CompanyB类公司进行工商年审工作")
}

func main() {
	println("——————程序开始运行.————————")

	company1 := CompanyA{}
	company2 := CompanyB{}

	accountingFirm := &AccountingFirm{}
	taxBureau := &TaxBureau{}
	tradeBureau := &TradeBureau{}

	println("———对CompanyA类公司进行处理———")
	company1.Accept(accountingFirm)
	company1.doVisit()

	company1.Accept(taxBureau)
	company1.doVisit()

	company1.Accept(tradeBureau)
	company1.doVisit()

	println("———对CompanyB类公司进行处理———")
	company2.Accept(accountingFirm)
	company2.doVisit()

	company2.Accept(taxBureau)
	company2.doVisit()

	company2.Accept(tradeBureau)
	company2.doVisit()

	println("——————程序运行结束.————————")

}
