package feeder

import (
	"deportStore/infraestructure/repo/inmemory"
	"sync"
	"testing"
)

func TestAddProducts(t *testing.T) {
	immRepo := inmemory.NewInmemoryRepo()
	serv := NewFeederService(immRepo)

	addProducts(serv, []string{"KASL-3001", "LPOS-3001"})
	addProducts(serv, []string{"KASL-3002", "KASL-3001"})
	addProducts(serv, []string{"KASL-3003", "LPOS-324"})

	uniqueProdNumber := 4
	duplicateProdNumber := 1
	invalidFormatProdNumber := 1

	if serv.repo.GetUniqueCount() != uniqueProdNumber {
		t.Errorf("Expected %d unique products, got: %d", uniqueProdNumber, serv.repo.GetUniqueCount())
	}

	if serv.repo.GetDuplicateCount() != duplicateProdNumber {
		t.Errorf("Expected %d duplicated products, got: %d ", duplicateProdNumber, serv.repo.GetDuplicateCount())

	}

	if serv.repo.GetDiscartCount() != invalidFormatProdNumber {
		t.Errorf("Expected %d discarted products, got: %d", invalidFormatProdNumber, serv.repo.GetDiscartCount())
	}
}

func TestConcurrentAddProducts(t *testing.T) {
	immRepo := inmemory.NewInmemoryRepo()
	serv := NewFeederService(immRepo)

	var wg sync.WaitGroup

	wg.Add(3)

	go func(ss FeederService) {
		defer wg.Done()
		addProducts(serv, []string{"KASL-3001", "LPOS-3001"})
	}(serv)

	go func(ss FeederService) {
		defer wg.Done()
		addProducts(serv, []string{"KASL-3002", "KASL-3001"})
	}(serv)

	go func(ss FeederService) {
		defer wg.Done()
		addProducts(serv, []string{"KASL-3003", "LPOS-324"})
	}(serv)

	uniqueProdNumber := 4
	duplicateProdNumber := 1
	invalidFormatProdNumber := 1

	wg.Wait()

	if serv.repo.GetUniqueCount() != uniqueProdNumber {
		t.Errorf("Expected %d unique products, got: %d", uniqueProdNumber, serv.repo.GetUniqueCount())
	}

	if serv.repo.GetDuplicateCount() != duplicateProdNumber {
		t.Errorf("Expected %d duplicated products, got: %d ", duplicateProdNumber, serv.repo.GetDuplicateCount())

	}

	if serv.repo.GetDiscartCount() != invalidFormatProdNumber {
		t.Errorf("Expected %d discarted products, got: %d", invalidFormatProdNumber, serv.repo.GetDiscartCount())
	}
}

func addProducts(ss FeederService, prods []string) {

	for _, prd := range prods {
		ss.RegisterProduct(prd)
	}
}
