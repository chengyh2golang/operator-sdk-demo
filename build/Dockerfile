FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/operator-sdk-demo \
    USER_UID=1001 \
    USER_NAME=operator-sdk-demo

# install operator binary
COPY build/_output/bin/operator-sdk-demo ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
