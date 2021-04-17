ARG PROJECT_PATH=C:\git\GoWorkspace\src\github.com\Nerzal\homeautomation

FROM tinygo/tinygo AS build

COPY . .
RUN tinygo build -target=wasm -o html/wasm.wasm dashboard/wasm.go
RUN go build -o api/app api/main.go

FROM gcr.io/distroless/static:nonroot AS service
USER nonroot:nonroot

EXPOSE 8080

ARG PROJECT_PATH
ENV PROJECT_PATH=$PROJECT_PATH

COPY --from=build ${PROJECT_PATH}/dashboard/html /html
COPY --from=build ${PROJECT_PATH}/api/app ./bin/app


CMD ["/bin/app"] 
