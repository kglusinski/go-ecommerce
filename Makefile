BUILD_FLAGS := GOARCH=amd64

.PHONY: start-cart start-order build-cart build-order start
build-cart:
	$(BUILD_FLAGS) go build -o ./bin/cart-service ./cmd/cart/...

build-order:
	$(BUILD_FLAGS) go build -o ./bin/order-service ./cmd/order/...

start:
	docker compose up -d