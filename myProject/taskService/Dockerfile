FROM golang:1.17-alpine
RUN mkdir taskService
COPY . /taskService
WORKDIR /taskService
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9000