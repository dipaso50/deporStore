package feeder

import "sync"

const maxClients = 5

var mutexRegister = &sync.Mutex{}
var mutex = &sync.Mutex{}

type IFeederService interface {
	RegisterProduct(product string)
	LimitReached() bool
	AcceptConnection()
}
type FeederService struct {
	uniqueProducts      map[string]string
	duplicatedProducts  map[string]string
	discartedProducts   map[string]string
	currentClientNumber int
}

func NewFeederService() FeederService {
	return FeederService{
		uniqueProducts:      make(map[string]string),
		duplicatedProducts:  make(map[string]string),
		discartedProducts:   make(map[string]string),
		currentClientNumber: 0,
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
	mutex.Lock()
	defer mutex.Unlock()
	return ser.currentClientNumber == maxClients
}

func (ser FeederService) AcceptConnection() {
	mutex.Lock()
	defer mutex.Unlock()
	ser.currentClientNumber++
}
