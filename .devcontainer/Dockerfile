ARG VARIANT="jammy"
ARG FROM_PLATFORM=linux/arm64
FROM --platform=${FROM_PLATFORM} mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

#RUN apt update
#RUN apt install -y musl-tools musl-dev
#RUN apt-get install -y build-essential
#RUN yes | apt install gcc-x86-64-linux-gnu

COPY core-libs.sh /tmp/
RUN yes | chmod +x /tmp/core-libs.sh \
    && bash /tmp/core-libs.sh \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/core-libs.sh

RUN echo "export PATH=\$PATH:/home/vscode/go/bin" >> /home/vscode/.bashrc
RUN echo "export GOPATH=/home/vscode/go" >> /home/vscode/.bashrc

RUN chown -R vscode:root /home/vscode/go

COPY install-dev-tools.sh /tmp/
RUN chmod +x /tmp/install-dev-tools.sh
