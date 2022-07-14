package modules

import "sync"

var (
	store map[string]interface{}
	lock  sync.Mutex
)

func Init() {
	store = map[string]interface{}{}
	lock = sync.Mutex{}
}

type Module = interface{}

func getModule(moduleName string) (Module, bool) {
	lock.Lock()
	defer lock.Unlock()

	value, ok := store[moduleName]
	return value, ok
}

func setModule(moduleName string, module Module) {
	lock.Lock()
	defer lock.Unlock()

	store[moduleName] = module
}

func Register(moduleName string, moduleCreate func(string) (Module, error)) Module {
	if module, ok := getModule(moduleName); ok {
		return module
	}
	module, err := moduleCreate(moduleName)
	if err != nil {
		panic(err)
	}
	setModule(moduleName, module)
	return module
}

//func Unregister(moduleName string) {
//	lock.Lock()
//	defer lock.Unlock()
//
//	delete(store, moduleName)
//}
