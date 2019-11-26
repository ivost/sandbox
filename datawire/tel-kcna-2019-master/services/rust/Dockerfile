FROM rust:buster as build-env

WORKDIR /usr/src/app
RUN USER=root cargo init
COPY Cargo.toml .
RUN cargo fetch
RUN cargo build --release
RUN rm -r target/release/app target/release/.fingerprint/app-*
COPY src src
RUN cargo build --release

FROM debian:buster

RUN apt-get -qq update && apt-get -qq install libssl1.1
COPY --from=build-env /usr/src/app/target/release/app /app
CMD ["/app"]
