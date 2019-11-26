############################
# STEP 1 build app using go-builder
############################
# go 1.12 alpine builder image
FROM 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:1.12 as builder

# args can be passed with --build-arg name=value
ARG PKG=github.com/braincorp/boss_pod_template
#ARG GIT_COMMIT=notset
#ARG LDFLAGS="-s -w -X pkg/version.REVISION="
#ARG LDFLAGS='-s -w -X github.com/braincorp/boss_pod_template/pkg/version.REVISION=todo'

# this is the working directory (user already created in builder img
WORKDIR /home/appuser

COPY . .
# go modules prep work

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN GOPROXY=https://proxy.golang.org go mod download

RUN go mod download && \
    go mod verify && \
    GIT_COMMIT=$(git rev-list -1 HEAD) && \
    go build -ldflags "-s -w \
               -X github.com/braincorp/boss_pod_template/pkg/version.REVISION=${GIT_COMMIT}" \
       -a -o bin/podinfo cmd/podinfo/* && \
    go build -ldflags "-s -w \
               -X github.com/braincorp/boss_pod_template/pkg/version.REVISION=${GIT_COMMIT}" \
       -a -o bin/podcli cmd/podcli/*

############################
# STEP 2 create standalone image
############################
FROM scratch

EXPOSE 8080

## Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /home/appuser

# Copy UI
COPY ./ui ./ui
# Copy binaries
COPY --from=builder /home/appuser/bin/pod* ./
COPY --from=builder /home/appuser/bin/podcli /usr/local/bin/
#RUN chown -R appuser:appuser ./
# Use an unprivileged user.
USER appuser
CMD ["./podinfo"]
