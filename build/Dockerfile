FROM alpine:3.8

ENV USER_UID=1001 \
    USER_NAME=app-simulator

# install binary
COPY _output/bin/app-simulator /usr/local/bin/app-simulator

COPY bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
