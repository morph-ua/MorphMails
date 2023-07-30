FROM scratch
COPY morph_mails /bin/server
ENTRYPOINT ["server"]
