FROM ubuntu

WORKDIR /app

ADD helloworld /app

EXPOSE 8080

CMD ["./helloworld"]
