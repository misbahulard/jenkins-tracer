# ============================= Builder Container ==============================

# Image for build golang code based on alpine golang
FROM golang:alpine AS build-env

# Run update
RUN apk add --update --no-cache ca-certificates git

# Run build on binary
WORKDIR /src
ADD . /src
RUN cd /src &&  go build -o jenkins-tracer

# ============================= Main Container =================================

# Image for apps supervisord and others based on alpine only
FROM alpine
LABEL maintainer="Misbahul Ardani"

# ============================= System =========================================

# Run update
RUN apk add --update --no-cache ca-certificates git libc6-compat curl jq

# # add needed some needed package for Testing only
# RUN apk add --no-cache libc6-compat curl jq

# Setting Time Zone on the container
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata \
    && cp /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone


# ============================= Jenkins Tracer =================================

# Create home directory for log agent and for the logging
RUN mkdir -p /opt/jenkins-tracer \
    && mkdir -p /var/log/jenkins-tracer

# Copy the config
COPY config.references.yaml /opt/jenkins-tracer/config.yaml

# copy the binary
COPY --from=build-env /src/jenkins-tracer /usr/local/bin/jenkins-tracer

# Entrypoint
ENTRYPOINT ["jenkins-tracer"]