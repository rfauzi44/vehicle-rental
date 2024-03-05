FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -v -o /app/vehicle-rental-api

EXPOSE 443
ENTRYPOINT [ "/app/vehicle-rental-api"]
CMD [ "serve" ]