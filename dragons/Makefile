all:
	go run genmazes.go > data.go
	go build -ldflags "-w" -o gofinal main.go data.go
	strip gofinal

clean:
	rm -f gofinal data.go
