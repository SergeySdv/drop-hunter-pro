FROM alpine
WORKDIR "/opt"
ADD migrations migrations
ADD ./cry cry
RUN chmod +x /opt/cry

ENTRYPOINT ["/opt/cry"]