FROM ubuntu:18.04

USER root

# Based on https://github.com/allthings/chromedriver with minor modifications: the base image (now ubuntu) and chromedriver package name (chromium-chromedriver now)

# Install chromium-chromedriver (which depends on chromium-browser):
RUN export DEBIAN_FRONTEND=noninteractive \
  && apt-get update \
  && apt-get install --no-install-recommends --no-install-suggests -y \
    chromium-chromedriver \
  # Start chromium-browser via wrapper script with --no-sandbox argument:
  && mv /usr/lib/chromium-browser/chromium-browser /usr/lib/chromium-browser/chromium-browser-original \
  && printf '%s\n' '#!/bin/sh' \
    'exec /usr/lib/chromium-browser/chromium-browser-original --no-sandbox "$@"' \
    > /usr/lib/chromium-browser/chromium-browser && chmod +x /usr/lib/chromium-browser/chromium-browser \
  # Remove obsolete files:
  && apt-get clean \
  && rm -rf \
    /tmp/* \
    /usr/share/doc/* \
    /var/cache/* \
    /var/lib/apt/lists/* \
    /var/tmp/*

# Make chromedriver available in the PATH:
RUN ln -s /usr/lib/chromium-browser/chromedriver /usr/local/bin/

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get -y install default-jre-headless && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://commons-repo.ritchiecli.io/install.sh | bash

RUN mkdir /rit
COPY . /rit
RUN sed -i 's/\r//g' /rit/set_umask.sh
RUN sed -i 's/\r//g' /rit/run.sh
RUN chmod +x /rit/set_umask.sh

WORKDIR /app
ENTRYPOINT ["/rit/set_umask.sh"]
CMD ["/rit/run.sh", "chromedriver", "--url-base=/wd/hub", "--port=4444", "--whitelisted-ips="]
