FROM alpine:3.22.2

WORKDIR /app

COPY main .

RUN chmod +x ./main

EXPOSE ${SERVER_PORT}

CMD ["./main"]