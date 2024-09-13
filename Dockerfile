FROM debian:stable-slim
COPY DockerWebServer /bin/DockerWebServer
ENV PORT=8080
CMD ["/bin/DockerWebServer"]
