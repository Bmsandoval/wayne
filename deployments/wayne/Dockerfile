FROM amazonlinux:2 AS base
RUN yum update -y
LABEL base=true

FROM base AS builder
# Install golang
RUN amazon-linux-extras install -y golang1.11
# Build release
WORKDIR /wayne
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .
RUN mkdir -p build/migrations
RUN cp secrets.ex.json build/secrets.json
RUN go build -o ./build/elf .
RUN cp migrations/* build/migrations/

FROM base
ENV AWS_REGION="${AWS_REGION}"
# setup the app and run it
WORKDIR /app
COPY --from=builder /wayne/build/ ./
RUN chmod +x elf
ENTRYPOINT ./elf
RUN echo "STARTED SERVER. If server unavailable, just try restarting the api container"