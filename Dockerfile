FROM public.ecr.aws/docker/library/golang:1.20.5 as builder

WORKDIR /code

COPY go.mod go.mod
COPY go.sum go.sum

RUN --mount=type=ssh go mod download

COPY . .

RUN make build-cart
RUN make build-order

FROM builder as cart_service

ENTRYPOINT ["./bin/cart-service"]

FROM builder as order_service

ENTRYPOINT ["./bin/order-service"]

