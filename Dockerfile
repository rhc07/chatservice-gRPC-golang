FROM golang:1.16.4-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app/cmd/server
RUN go build -o main .
COPY . .
CMD [ "/app/server/main" ]