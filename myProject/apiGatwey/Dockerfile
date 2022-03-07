FROM golang:1.17-alpine
RUN mkdir apiGatwey
COPY . /apiGatwey
WORKDIR /apiGatwey
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 8080