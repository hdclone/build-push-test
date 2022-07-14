package mock

//go:generate solc --overwrite --base-path=. --abi MockBridge.sol -o ./abi/
//go:generate solc --overwrite --base-path=. --bin MockBridge.sol -o ./bin/
//go:generate abigen --abi ./abi/MockBridge.abi --bin ./bin/MockBridge.bin --pkg mock --out MockBridge.go
