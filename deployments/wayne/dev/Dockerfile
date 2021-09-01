FROM golang:1.17-alpine AS base
RUN apt update -y
LABEL base=true

FROM base AS runner
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .
RUN mkdir -p build/db/migrations
RUN cp secrets.ex.json build/secrets.json
RUN go build -o ./build/elf .
RUN cp db/migrations/* build/db/migrations/

FROM base
ENV AWS_REGION="${AWS_REGION}"
# setup the app and run it
WORKDIR /app
COPY --from=builder /wayne/build/ ./
RUN chmod +x elf
ENTRYPOINT ./elf
RUN echo "STARTED SERVER. If server unavailable, just try restarting the api container"