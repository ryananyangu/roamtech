############################
# STEP 1 build executable binary
############################
# golang alpine 1.14
FROM golang:alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/rometech
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/rometech .

############################
# STEP 2 build a small image
############################
FROM scratch


# Copy our static executable
COPY --from=builder /go/bin/rometech /go/bin/rometech

# Use an unprivileged user.
USER appuser:appuser

ARG PROJECT_ID


# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# COPY ./${SERVICE_ACCOUNT_FNAME}.json /


ENV PROJECT_ID=${PROJECT_ID}

# Run the flowbuilder binary.
EXPOSE 8080

ENTRYPOINT ["/go/bin/rometech"]