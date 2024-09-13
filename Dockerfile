FROM debian:stable-slim
COPY DockerWebServer /bin/DockerWebServer
CMD ["/bin/DockerWebServer"]
