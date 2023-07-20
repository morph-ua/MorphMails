FROM scratch
COPY helium /bin/server
ENTRYPOINT ["server"]
