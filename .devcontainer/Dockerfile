ARG VARIANT="jammy"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}


COPY core-libs.sh /tmp/
RUN yes | chmod +x /tmp/core-libs.sh \
    && bash /tmp/core-libs.sh \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/core-libs.sh

COPY install-dev-tools.sh /tmp/
RUN chmod +x /tmp/install-dev-tools.sh
