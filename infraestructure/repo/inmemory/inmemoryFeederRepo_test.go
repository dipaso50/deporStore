package inmemory

import (
	"deportStore/domain"
	"testing"
)

func TestSave(t *testing.T) {
	repo := NewInmemoryRepo()

	repo.SaveUnique(domain.Product{SKUCode: "Test"})

	if repo.GetUniqueCount() != 1 {
		t.Errorf("Expected %d unique products , got %d", 1, repo.GetUniqueCount())
	}

	repo.SaveDuplicate(domain.Product{SKUCode: "Test"})

	if repo.GetDuplicateCount() != 1 {
		t.Errorf("Expected %d duplicated products , got %d", 1, repo.GetDuplicateCount())
	}

	repo.SaveDiscartedSKU("Test")

	if repo.GetDiscartCount() != 1 {
		t.Errorf("Expected %d discarted products , got %d", 1, repo.GetDiscartCount())
	}
}

func TestGetUniqueProduct(t *testing.T) {
	sku := "test"
	repo := NewInmemoryRepo()

	repo.SaveUnique(domain.Product{SKUCode: sku})

	if _, exist := repo.GetUniqueProduct(sku); !exist {
		t.Errorf("Expecting one product, got 0\n")
	}

}
