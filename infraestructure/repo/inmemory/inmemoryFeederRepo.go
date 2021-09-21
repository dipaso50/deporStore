package inmemory

import "deportStore/domain"

type InMemoryFeederRepo struct {
	uniqueProducts     map[string]domain.Product
	duplicatedProducts map[string]domain.Product
	discartedProducts  map[string]string
}

func NewInmemoryRepo() *InMemoryFeederRepo {
	return &InMemoryFeederRepo{
		uniqueProducts:     make(map[string]domain.Product),
		duplicatedProducts: make(map[string]domain.Product),
		discartedProducts:  make(map[string]string),
	}
}

func (rep *InMemoryFeederRepo) SaveUnique(product domain.Product) {
	rep.uniqueProducts[product.SKUCode] = product
}
func (rep *InMemoryFeederRepo) SaveDuplicate(product domain.Product) {
	rep.duplicatedProducts[product.SKUCode] = product
}
func (rep *InMemoryFeederRepo) SaveDiscartedSKU(sku string) {
	rep.discartedProducts[sku] = sku
}
func (rep *InMemoryFeederRepo) GetUniqueProduct(sku string) (domain.Product, bool) {
	val, exist := rep.uniqueProducts[sku]
	return val, exist
}
func (rep *InMemoryFeederRepo) GetUniqueCount() int {
	return len(rep.uniqueProducts)
}
func (rep *InMemoryFeederRepo) GetDiscartCount() int {
	return len(rep.discartedProducts)
}
func (rep *InMemoryFeederRepo) GetDuplicateCount() int {
	return len(rep.duplicatedProducts)
}
