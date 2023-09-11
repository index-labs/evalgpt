FROM python:3.11.5-slim-bookworm

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY bin/evalgpt .

CMD ["evalgpt", "--run-as-server", "--listen-addr", "0.0.0.0:8080"]