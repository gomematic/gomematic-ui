FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Gomematic UI" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

EXPOSE 9000 80 443
VOLUME ["/var/lib/gomematic"]

ENV GOMEMATIC_UI_ASSETS /usr/share/gomematic
ENV GOMEMATIC_UI_STORAGE /var/lib/gomematic

ENTRYPOINT ["/usr/bin/gomematic-ui"]
CMD ["server"]

RUN apk add --no-cache ca-certificates mailcap

COPY dist/static /usr/share/gomematic
COPY dist/binaries/gomematic-ui-*-linux-amd64 /usr/bin/gomematic-ui
