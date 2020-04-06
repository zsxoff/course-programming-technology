# -----------------------------------------------------------------------------

FROM gcc:9.3.0 as build

# Copy current dir into container.
ADD . /app
WORKDIR /app

# Compile app.
RUN make all

# -----------------------------------------------------------------------------

FROM debian:10

# Add non-root user.
RUN groupadd -r application && useradd -r -g application application
USER application

# Copy application from build container.
WORKDIR /app
COPY --from=build /app/app .
COPY --from=build /app/libs/libmatrix.so ./libs/libmatrix.so

ENTRYPOINT ["./app"]