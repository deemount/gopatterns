package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

/*

 */

// Service represents an external service
type Service interface {
	Do(action string)
}

// ServiceFactory is a component capable of constructing
// a service interface with a given service name
type ServiceFactory interface {
	New(serviceName string) (Service, error)
}

// stub service implementations
type FooService struct {
	Service
}

func NewFooService() Service {
	return &FooService{}
}

func (srv *FooService) Do(action string) {
	log.Println(action)
}

type BarService struct {
	Service
}

func NewBarService() Service {
	return &BarService{}
}

func (srv *BarService) Do(action string) {
	log.Println(action)
}

// default ServiceFactory implementation
type serviceFactory struct {
	mu    sync.Mutex
	cache map[string]Service
}

func (fac *serviceFactory) New(serviceName string) (Service, error) {

	log.Printf("new: arg %s", serviceName)

	// check if the service has already been constructed -- (1)
	serviceName = strings.TrimSpace(serviceName)
	if srv, exists := fac.cache[serviceName]; exists {
		log.Println("new: (1)")
		return srv, nil
	}

	log.Println("new: after first check - lock")
	fac.mu.Lock()
	defer fac.mu.Unlock()
	log.Println("new: unlock - before double check")

	// ensure that the service wasn't already created during
	// lock-acquisition by a similar goroutine -- (2)
	if srv, exists := fac.cache[serviceName]; exists {
		log.Println("new: (2)")
		return srv, nil
	}

	log.Println("new: after double check")

	// construct the service and add it to the cache
	var newService Service

	switch serviceName {
	case "foo":
		newService = NewFooService()
		break
	case "bar":
		barSrv := new(BarService)
		newService = (barSrv).Service
		break
	}

	log.Printf("new: new service %+v", newService)

	if newService == nil {
		return nil, fmt.Errorf("don't know how to construct %s", serviceName)
	}

	fac.cache[serviceName] = newService

	log.Printf("new: factory cache %+v", fac.cache)

	return newService, nil
}

type serviceHandler struct {
	srvFactory ServiceFactory
}

func (handler *serviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// the code below extracts the service name
	// and action name
	pathComponents := strings.Split(r.URL.Path, "/")
	l := len(pathComponents) - 1
	if pathComponents[l] == "/" {
		pathComponents = pathComponents[:l]
		l = len(pathComponents)
	}
	actionName := pathComponents[l]
	srvName := pathComponents[l-1]

	log.Printf("serve http: action %s, service %s", actionName, srvName)

	// construct the service from the factory
	// and call the specified action
	srv, err := handler.srvFactory.New(srvName)
	log.Printf("serve http: service &%+v", srv)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	srv.Do(actionName)
}

func main() {

	factory := serviceFactory{
		cache: make(map[string]Service, 1),
	}

	fooSrv := new(FooService)
	factory.cache["foo"] = fooSrv.Service

	handler := &serviceHandler{
		srvFactory: new(serviceFactory),
	}

	// requests will be of the form
	// /dispatch/<serviceName>/<action>
	http.Handle("/dispatch/", handler)
	http.ListenAndServe(":8080", nil)
}
