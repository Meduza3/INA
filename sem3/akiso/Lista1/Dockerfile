FROM debian:latest

COPY . /app

WORKDIR /app

RUN apt-get update && apt-get install -y g++ make

CMD ["make", "test"]