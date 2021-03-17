FROM golang:1.15.8
RUN apt-get update -y
RUN apt-get install -y \
        git python jq curl vim wget \
        graphviz gv

WORKDIR /src

COPY . .
