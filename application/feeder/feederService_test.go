package feeder

import (
	"sync"
	"testing"
)

const maxAllowed = 5

func TestClientNumberRestriction(t *testing.T) {

	serv := NewFeederService(maxAllowed)

	for i := 0; i < 10; i++ {
		if !serv.LimitReached() {
			serv.AcceptConnection()
		}
	}

	if serv.currentClientNumber > maxAllowed {
		t.Errorf("Client number greater than expected, expected: %d, got: %d", maxAllowed, serv.currentClientNumber)
	}
}

func TestConcurrentClientNumberRestriction(t *testing.T) {
	serv := NewFeederService(maxAllowed)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(ss FeederService) {
			defer wg.Done()
			if !ss.LimitReached() {
				ss.AcceptConnection()
			}
		}(serv)
	}

	wg.Wait()

	if serv.currentClientNumber > maxAllowed {
		t.Errorf("Client number greater than expected, expected: %d, got: %d", maxAllowed, serv.currentClientNumber)
	}
}

func TestAddProducts(t *testing.T) {
	serv := NewFeederService(maxAllowed)

	addProducts(serv, []string{"KASL-3001", "LPOS-3001"})
	addProducts(serv, []string{"KASL-3002", "KASL-3001"})
	addProducts(serv, []string{"KASL-3003", "LPOS-324"})

	uniqueProdNumber := 4
	duplicateProdNumber := 1
	invalidFormatProdNumber := 1

	if len(serv.uniqueProducts) != uniqueProdNumber {
		t.Errorf("Expected %d unique products, got: %d {%s}", uniqueProdNumber, len(serv.uniqueProducts), serv.uniqueProducts)
	}

	if len(serv.duplicatedProducts) != duplicateProdNumber {
		t.Errorf("Expected %d duplicated products, got: %d {%s}", duplicateProdNumber, len(serv.duplicatedProducts), serv.duplicatedProducts)

	}

	if len(serv.discartedProducts) != invalidFormatProdNumber {
		t.Errorf("Expected %d discarted products, got: %d {%s}", invalidFormatProdNumber, len(serv.discartedProducts), serv.discartedProducts)
	}
}

func TestConcurrentAddProducts(t *testing.T) {
	serv := NewFeederService(maxAllowed)

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

	if len(serv.uniqueProducts) != uniqueProdNumber {
		t.Errorf("Expected %d unique products, got: %d {%s}", uniqueProdNumber, len(serv.uniqueProducts), serv.uniqueProducts)
	}

	if len(serv.duplicatedProducts) != duplicateProdNumber {
		t.Errorf("Expected %d duplicated products, got: %d {%s}", duplicateProdNumber, len(serv.duplicatedProducts), serv.duplicatedProducts)

	}

	if len(serv.discartedProducts) != invalidFormatProdNumber {
		t.Errorf("Expected %d discarted products, got: %d {%s}", invalidFormatProdNumber, len(serv.discartedProducts), serv.discartedProducts)
	}
}

func addProducts(ss FeederService, prods []string) {
	if !ss.LimitReached() {
		ss.AcceptConnection()

		for _, prd := range prods {
			ss.RegisterProduct(prd)
		}
	}
}
