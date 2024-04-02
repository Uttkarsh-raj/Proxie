FROM golang:1.20.5-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY . /app/

EXPOSE 7000

RUN go build -o main.exe

CMD [ "/app/main.exe" ]