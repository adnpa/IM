FROM ubuntu:latest
LABEL authors="hz"

ENTRYPOINT ["top", "-b"]