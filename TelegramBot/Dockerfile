FROM python:alpine

ENV PYTHONUNBUFFERED=1

ENV PYTHONDONTWRITEBYTECODE=1

WORKDIR /app

COPY ./requirements.txt ./

RUN pip install --upgrade pip

RUN pip install --no-cache-dir -r ./requirements.txt

COPY ./ ./

EXPOSE 3000

CMD ["python main.py"]