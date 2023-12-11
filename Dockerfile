FROM golang:latest
COPY go.mod /app/
COPY go.sum /app/
COPY ./ /app
WORKDIR /app/cmd/app
RUN go build main.go
RUN ls -la
CMD ["./main"]


