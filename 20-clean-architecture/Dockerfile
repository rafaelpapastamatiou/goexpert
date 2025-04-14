FROM golang:1.24.2-bookworm

WORKDIR /app
RUN apt-get update && \
    apt-get install -y curl lsb-release  && \
    apt-get clean && \
    curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - && \
    echo "deb https://packagecloud.io/golang-migrate/migrate/debian/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list && \
    apt-get update && \
    apt-get install -y migrate && \
    apt-get clean

CMD [ "tail", "-f", "/dev/null" ]