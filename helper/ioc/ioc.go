package ioc

import (
	"errors"
	"fmt"
	"reflect"
)

// IOCContainer manages dependency injection.

var services map[string]reflect.Value = map[string]reflect.Value{} 
var servicesByString map[string]reflect.Value = map[string]reflect.Value{} 


// Register registers a service instance by its type.
func RegisterSingleton(serviceTypeObj interface{}, service interface{}) error {
	serviceType := reflect.TypeOf(serviceTypeObj).Elem().String()
	_, exists := services[serviceType]
	if exists {
		return fmt.Errorf("service of type %s already exists", serviceType)
	}
	services[serviceType] = reflect.ValueOf(service)
	return nil
}

// Register registers a service instance by name.
func RegisterSingletonWihString(serviceName string, service interface{}) error {
	_, exists := servicesByString[serviceName]
	if exists {
		return fmt.Errorf("service of type %s already exists", serviceName)
	}
	servicesByString[serviceName] = reflect.ValueOf(service)
	return nil
}

// Resolve retrieves a service instance by its type.
func GetInstanceSingleton(serviceType interface{}) (interface{}, error) {
	t := reflect.TypeOf(serviceType).Elem().String() // Get the actual type
	service, exists := services[t]
	if !exists {
		return nil, errors.New("service not found: " + t)
	}
	return service.Interface(), nil
}

// Resolve retrieves a service instance by name.
func GetInstanceSingletonString(serviceName string) (interface{}, error) {
	service, exists := servicesByString[serviceName]
	if !exists {
		return nil, errors.New("service not found: " + serviceName)
	}
	return service.Interface(), nil
}
