FROM scratch
EXPOSE 8080
ENTRYPOINT ["/src"]
COPY ./bin/ /