# because go mod is required for project dependencies,
# there is no build stage. You need to build the binary
# using "CGO_ENABLED=0 GOOS=linux go build -o portal *.go"
# and then execute the dockerfile with the binary in the same
# directory as the dockerfile

#build stage
# FROM golang:alpine AS builder 
# RUN mkdir /app 
# ADD . /app
# WORKDIR /app
# RUN apk add --no-cache git
# RUN go get -d
# RUN CGO_ENABLED=0 GOOS=linux go build -o portal *.go

#production image
FROM scratch
COPY ./portal .
EXPOSE 8080
CMD ["./portal"]