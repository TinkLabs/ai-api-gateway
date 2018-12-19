FROM bananabb/ubuntu-go:go.1.3.1

# Copy project to container
ADD . /var/src

# Setup project
RUN cd /var/src/gateway \
    dep ensure

# Basic setup
WORKDIR /var/src/gateway
EXPOSE 80 443

CMD ["go","run","main.go"]