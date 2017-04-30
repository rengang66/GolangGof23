// factorymethod
package main

type Company interface {
	BulidProduct()
}

type Product interface {
	DoUse()
}

type ProductA struct {
	productType      string
	productParameter string
}

type ProductB struct {
	productType      string
	productParameter string
}

type CompanyA struct {
}

func (Pro *ProductA) initialize() {
	Pro.productType = "ProductA"
	Pro.productParameter = "A"
}

func (Pro *ProductA) setProductType(s string) {
	Pro.productType = s
}

func (Pro *ProductA) getProductType() string {
	return Pro.productType
}

func (Pro *ProductA) setProductParameter(s string) {
	Pro.productParameter = s
}

func (Pro *ProductA) getProductParameter() string {
	return Pro.productParameter
}

func (Pro ProductA) DoUse() {
	println("这是ProductA实现的功能")
}

func (Pro *ProductB) initialize() {
	Pro.productType = "ProductB"
	Pro.productParameter = "B"
}

func (Pro *ProductB) setProductType(s string) {
	Pro.productType = s
}

func (Pro *ProductB) getProductType() string {
	return Pro.productType
}

func (Pro *ProductB) setProductParameter(s string) {
	Pro.productParameter = s
}

func (Pro *ProductB) getProductParameter() string {
	return Pro.productParameter
}

func (Pro ProductB) DoUse() {
	println("这是ProductB实现的功能")
}

func (company *CompanyA) BulidProduct(Parameter string) Product {
	if Parameter == "A" {
		product := ProductA{}
		product.initialize()
		return product
	} else if Parameter == "B" {
		product := ProductB{}
		product.initialize()
		return product
	} else {
		return nil
	}
}

func main() {
	println("——————程序开始运行.————————")
	company1 := CompanyA{}
	product := company1.BulidProduct("A")
	product.DoUse()

	//根据传入的参数得到ProductB产品
	product = company1.BulidProduct("B")
	product.DoUse()

	println("——————程序运行结束.————————")

}
