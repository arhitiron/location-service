FROM centurylink/ca-certs
EXPOSE 8000
COPY location.bin /location
ENTRYPOINT ["/location"]