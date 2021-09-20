package feeder

import (
	"fmt"
	"sync"
)

var mutexRegister = &sync.Mutex{}
var mutex = &sync.Mutex{}

type IFeederService interface {
	RegisterProduct(product string)
	LimitReached() bool
	AcceptConnection()
	Report()
}
type FeederService struct {
	uniqueProducts      map[string]string
	duplicatedProducts  map[string]string
	discartedProducts   map[string]string
	currentClientNumber int
	maxClients          int
}

func NewFeederService(maxClients int) FeederService {
	return FeederService{
		uniqueProducts:      make(map[string]string),
		duplicatedProducts:  make(map[string]string),
		discartedProducts:   make(map[string]string),
		currentClientNumber: 0,
		maxClients:          maxClients,
	}
}

func (serv FeederService) RegisterProduct(product string) {

	mutexRegister.Lock()
	defer mutexRegister.Unlock()

	if !validFormat(product) {
		serv.discartedProducts[product] = product
		return
	}

	if _, exist := serv.uniqueProducts[product]; exist {
		serv.duplicatedProducts[product] = product
		return
	}

	serv.uniqueProducts[product] = product
}

func (ser FeederService) LimitReached() bool {
	return ser.currentClientNumber == ser.maxClients
}

func (ser FeederService) AcceptConnection() {
	mutex.Lock()
	defer mutex.Unlock()
	ser.currentClientNumber++
}

func (ser FeederService) Report() {
	fmt.Printf("Received %d unique product skus, %d duplicates, %d discard values \n",
		len(ser.uniqueProducts), len(ser.duplicatedProducts), len(ser.discartedProducts))
}

func (ser FeederService) GracefullShutdown() {
}
