package feeder

import (
	"deportStore/domain"
	"fmt"
	"sync"
)

var mutexRegister = &sync.Mutex{}

type IFeederService interface {
	RegisterProduct(product string)
	PrintReport()
}
type FeederService struct {
	repo IFeederRepo
}

func NewFeederService(rep IFeederRepo) FeederService {
	return FeederService{
		repo: rep,
	}
}

func (serv FeederService) RegisterProduct(sku string) {

	mutexRegister.Lock()
	defer mutexRegister.Unlock()

	repo := serv.repo

	if !validFormat(sku) {
		repo.SaveDiscartedSKU(sku)
		return
	}

	if _, exist := repo.GetUniqueProduct(sku); exist {
		repo.SaveDuplicate(domain.Product{SKUCode: sku})
		return
	}

	repo.SaveUnique(domain.Product{SKUCode: sku})
}

func (ser FeederService) PrintReport() {
	repo := ser.repo
	uniqueCount := repo.GetUniqueCount()
	duplicatedCount := repo.GetDuplicateCount()
	discartedCount := repo.GetDiscartCount()

	fmt.Printf("Received %d unique product skus, %d duplicates, %d discard values \n", uniqueCount, duplicatedCount, discartedCount)
}
