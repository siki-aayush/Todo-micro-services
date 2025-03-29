FROM alpine:latest

RUN mkdir /app

COPY todoServiceApp /app

CMD ["/app/todoServiceApp"]


