// flyweight
package main

type Document interface {
	readDocument()
}

type AdministrativeDocument struct {
}

func (doc *AdministrativeDocument) readDocument() {
	println("请阅读行政文档。")
}

type FinancialDocument struct {
}

func (doc *FinancialDocument) readDocument() {
	println("请阅读财务文档。")
}

type TechnicalDocument struct {
}

func (doc *TechnicalDocument) readDocument() {
	println("请阅读技术文档。")
}

type DocumentRepository struct {
	DocumentMap map[string]Document
}

func (repository *DocumentRepository) initialize() {
	repository.DocumentMap = make(map[string]Document)
	repository.DocumentMap["行政文档"] = Document(&AdministrativeDocument{})
	repository.DocumentMap["技术文档"] = Document(&TechnicalDocument{})
	repository.DocumentMap["财务文档"] = Document(&FinancialDocument{})
}

func (repository *DocumentRepository) getDocument(docType string) Document {

	if repository.DocumentMap[docType] != nil {
		return repository.DocumentMap[docType]
	} else {
		println("没有此类文档。")
		return nil
	}
}

func main() {
	println("——————程序开始运行.————————")
	repository := DocumentRepository{}
	repository.initialize()

	doc1 := repository.getDocument("行政文档")

	if doc1 != nil {
		doc1.readDocument()
	}

	doc2 := repository.getDocument("技术文档")
	if doc2 != nil {
		doc2.readDocument()
	}

	doc3 := repository.getDocument("财务文档")
	if doc3 != nil {
		doc3.readDocument()
	}

	doc4 := repository.getDocument("人事文档")
	if doc4 != nil {
		doc4.readDocument()
	}

	println("——————程序运行结束.————————")
}
