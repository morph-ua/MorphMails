FROM scratch
COPY AtomicEmails /bin/server
ENTRYPOINT ["server"]
