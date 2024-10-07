FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o api-server blops-me/cmd/api-server

FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update && apt install -y \
    python3 \
    python3-pip \
    python3-venv \
    python3-virtualenv \
    libmagic1 \
    libmagic-dev \
    && apt clean

WORKDIR /app

COPY . .
COPY --from=builder /app/api-server .

RUN virtualenv /opt/venv
RUN /opt/venv/bin/pip install python-magic
RUN /opt/venv/bin/pip install pypdf
RUN /opt/venv/bin/pip install python-docx
RUN /opt/venv/bin/pip install python-pptx
RUN /opt/venv/bin/pip install xlrd
RUN /opt/venv/bin/pip install openpyxl
RUN /opt/venv/bin/pip install cryptography

EXPOSE 8010

CMD ["./api-server"]