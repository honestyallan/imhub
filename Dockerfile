FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates

# Install bash
RUN apk add --no-cache bash

COPY --from=hub/builder:latest /code/build/hub /usr/bin/hub

EXPOSE 26656 26657 1317 9090

CMD ["hub"]
