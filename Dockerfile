FROM scratch

COPY acr-webhooks-server /
COPY etc/passwd /etc/passwd

USER nobody

ENTRYPOINT ["/acr-webhooks-server"]
