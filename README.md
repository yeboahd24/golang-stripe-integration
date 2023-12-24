# Create Product

```bash
curl -X POST http://localhost:8080/mido/products \
     -H "Content-Type: application/json" \
     -d '{"title":"Product 3","description":"Description of Product 3","price":400.50}' \

```

# List Products

```bash
curl -X GET http://localhost:8080/mido/products
```

# Initiate payment

```bash
curl -X POST http://localhost:8080/mido/create-payment-intent \
     -H "Content-Type: application/json" \
     -d '{"id":2}' \


```

# Project Initialization

- install golang
- install project dependencies by running `go mod tidy`
- run project by running `go run main.go`
# golang-stripe-integration
