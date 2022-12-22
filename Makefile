contract:
	mkdir build
	solcjs --optimize --abi src/contracts/LuisCoin.sol -o build && mv build/*.abi build/LuisCoin.abi
	solcjs --optimize --bin src/contracts/LuisCoin.sol -o build && mv build/*.bin build/LuisCoin.bin
	abigen --abi=./build/LuisCoin.abi --bin=./build/LuisCoin.bin --pkg=contracts --out=./src/contracts/LuisCoin.go

run:
	go run src/main.go

clean:
	rm build/* && rmdir build
	rm .env
	rm src/contracts/*.go

build:
	go mod download
	CGO_ENABLED=0 go build -installsuffix 'static' -o /server github.com/luisfc68/luis-coin/src

