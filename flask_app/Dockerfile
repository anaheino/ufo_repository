FROM python:3.8-slim
WORKDIR /app
COPY . /app
WORKDIR /app/src
COPY src/requirements.txt .
RUN pip install --trusted-host pypi.python.org -r requirements.txt
EXPOSE 5000
ENV NAME FlaskApp
CMD ["python", "app.py"]