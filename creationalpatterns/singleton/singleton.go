// singleton
package main

type SaleMan struct {
	name    string
	service string
}

func (saleMan *SaleMan) initialize(name string, service string) {
	saleMan.name = name
	saleMan.service = service
}

func (saleMan *SaleMan) getName() string {
	return saleMan.name
}

func (saleMan *SaleMan) setName(name string) {
	saleMan.name = name
}

func (saleMan *SaleMan) getService() string {
	return saleMan.service
}

func (saleMan *SaleMan) setService(service string) {
	saleMan.service = service
}

type ServiceManager struct {
	saleMan *SaleMan
}

func (serviceManager *ServiceManager) initialize(saleMan *SaleMan) {
	serviceManager.saleMan = saleMan
}

func (serviceManager *ServiceManager) getSaleManService() *SaleMan {
	if serviceManager.saleMan != nil {
		return serviceManager.saleMan
	} else {
		return nil
	}
}

func main() {
	println("——————程序开始运行.————————")

	saleMan := &SaleMan{"小刘", "小刘的服务"}
	service := ServiceManager{saleMan}

	print("第一次获得服务：")
	saleman := service.getSaleManService()
	println(saleman.getService())

	print("第二次获得服务：")
	saleman = service.getSaleManService()
	println(saleman.getService())

	print("第三次获得服务：")
	saleman = service.getSaleManService()
	println(saleman.getService())

	println("——————程序运行结束.————————")
}
