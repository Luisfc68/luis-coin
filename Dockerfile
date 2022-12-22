ARG GO_VERSION=1.19
ARG GETH_VERSION=v1.10.26

FROM ethereum/client-go:alltools-${GETH_VERSION} AS contractCompiler

COPY ./src/contracts/LuisCoin.sol .

RUN apk add --update nodejs npm && npm install -g solc
RUN solcjs --optimize --abi LuisCoin.sol -o build && mv build/*.abi build/LuisCoin.abi
RUN solcjs --optimize --bin LuisCoin.sol -o build && mv build/*.bin build/LuisCoin.bin
RUN abigen --abi=/build/LuisCoin.abi --bin=./build/LuisCoin.bin --pkg=contracts --out=/build/LuisCoin.go

FROM golang:${GO_VERSION} AS builder

WORKDIR /src

COPY ./ ./
COPY --from=contractCompiler /build/LuisCoin.go ./src/contracts/

RUN make build

FROM scratch AS runner
COPY --from=builder /server /server

ENV PORT=:8080

EXPOSE 8080
ENTRYPOINT ["/server"]