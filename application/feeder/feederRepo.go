package feeder

import "deportStore/domain"

type IFeederRepo interface {
	SaveUnique(product domain.Product)
	SaveDuplicate(product domain.Product)
	SaveDiscartedSKU(sku string)
	GetUniqueProduct(sku string) (domain.Product, bool)
	GetUniqueCount() int
	GetDiscartCount() int
	GetDuplicateCount() int
}
