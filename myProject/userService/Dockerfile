FROM golang:1.17-alpine
RUN mkdir userService
COPY . /userService
WORKDIR /userService
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9001
