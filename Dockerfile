FROM golang:1.18 AS dev

WORKDIR /usr/src/app/

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o api-routes


FROM golang:1.18 
COPY --from=dev /usr/src/app/api-routes /app/ 

EXPOSE 3006
#CMD [ "tail", "-f", "/dev/null" ]
CMD [ "/app/api-routes" ]