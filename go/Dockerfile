# syntax=docker/dockerfile:1                                                               
                                                                                           
# Build the application from source                                                        
FROM golang:1.21.4 AS build-stage                                                          
                                                                                           
WORKDIR /app                                                                               
                                                                                           
COPY *.mod ./                                                                      
RUN go mod download                                                                        
                                                                                           
COPY *.go ./                                                                               
                                                                                           
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]
