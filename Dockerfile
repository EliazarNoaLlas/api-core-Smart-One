FROM surnet/alpine-wkhtmltopdf:3.8-0.12.5-full as wkhtmltopdf
FROM alpine:3.8
MAINTAINER Soporte <soporte@smartc.pe>

# wkhtmltopdf install dependencies
RUN apk add --no-cache \
        libstdc++ \
        libx11 \
        libxrender \
        libxext \
        libssl1.0 \
        ca-certificates \
        fontconfig \
        freetype \
        ttf-droid \
        ttf-freefont \
        ttf-liberation

# wkhtmltopdf copy bins from ext image
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /usr/local/bin/wkhtmltopdf
#COPY --from=wkhtmltopdf /bin/wkhtmltoimage /usr/local/bin/wkhtmltoimage
#COPY --from=wkhtmltopdf /bin/libwkhtmltox* /usr/local/bin/

# Install required packages
RUN apk update

RUN apk add dmidecode
RUN apk add ca-certificates
RUN apk add openssl
RUN apk --no-cache add tzdata

COPY ./app /home/app

EXPOSE 80

CMD ["/home/app"]
