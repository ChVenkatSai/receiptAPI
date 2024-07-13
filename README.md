# Receipt Points Calculator API

## Project Structure

```
receiptAPI/
├── cmd/
│ └── receiptAPI/
│ │ └── main.go
├── internal/
│ ├── api/
│ │ └── handler.go
│ ├── service/
│ │ └── service.go
│ └── storage/
│ │ └──  storage.go
├── pkg/
│ └── models/
│ │ └── receipt.go
├── go.mod
├── go.sum
```

## Setup

### To run the project:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/ChVenkatSai/receiptAPI.git
   cd receiptAPI
2. **Initialize the Go module:**
   ```sh
   go mod tidy
3. **Go to folder containing main.go and run it:**
   ```sh
   cd cmd/receiptAPI
   go run main.go

Use Postman or curl to the localhost (port:8080) and test the application. 
