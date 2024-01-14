FROM alpine:3.17 AS release

COPY --from=ozon/base /source/bin/url /url

CMD ["/url"]