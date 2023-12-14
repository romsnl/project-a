FROM golang:1.21-alpine AS buildo

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/project-a

FROM scratch

WORKDIR /

COPY --from=buildo /build/bin/project-a /project-a

EXPOSE 3001

CMD [ "./project-a" ]