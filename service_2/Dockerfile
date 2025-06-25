FROM python:3.13-slim

WORKDIR /app
COPY pyproject.toml .

RUN pip install flask
COPY app.py .

EXPOSE 8002

CMD ["python", "app.py"]
