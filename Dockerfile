FROM moby/buildkit:v0.9.3
WORKDIR /acs
COPY acs README.md /acs/
ENV PATH=/acs:$PATH
ENTRYPOINT [ "/bhojpur/acs" ]