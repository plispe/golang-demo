FROM alpine:3.4
MAINTAINER Petr Pliska <petr.pliska@post.cz>
RUN apk add --update bash
COPY golang-demo golang-demo
ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/golang-demo"]
