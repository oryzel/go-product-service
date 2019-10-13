FROM golang:latest 

LABEL version="1.0"

RUN mkdir /go/src/app

RUN go get -u github.com/golang/dep/cmd/dep

ADD ./main.go /go/src/app
COPY ./Gopkg.toml /go/src/app

WORKDIR /go/src/app 

COPY . .

RUN dep ensure 
RUN go build -o project-service

EXPOSE 1001

CMD ["./project-service"]