FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o server
EXPOSE 8086
CMD [ "/app/server" ]
