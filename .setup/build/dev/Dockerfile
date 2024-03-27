FROM golang:1.22

ENV PATH="/go/bin:${PATH}"
ENV WORKDIR="/app"

WORKDIR $WORKDIR

COPY . $WORKDIR

RUN apt-get update && \
    go install github.com/golang/mock/mockgen@v1.6.0 && \
    go install github.com/cosmtrek/air@latest && \
    go mod download
    
ENTRYPOINT ["air", "-d"]