FROM golang:alpine
COPY . ./simpleChatInGo
WORKDIR /go/simpleChatInGo/src/public/script
RUN apk add npm
RUN npm install typescript -g
RUN tsc *.ts; rm -rf *.ts
RUN npm uninstall -g typescript
RUN apk del npm
WORKDIR /go/simpleChatInGo/src/
RUN go build main.go

CMD ["./main"]