# write docker file for go?

FROM golang:1.20-alpine

WORKDIR /app

COPY . .
#COPY go.mod ./
#COPY go.sum ./
RUN go mod download

#COPY *.go ./
ENV CONN_URL="root:root@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
RUN go build -o /todolist

EXPOSE 3000

CMD [ "/todolist" ]