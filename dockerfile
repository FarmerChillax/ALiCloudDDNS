FROM python:3.8-slim-buster

WORKDIR /ddnsCore

COPY . .

RUN pip3 install -r requirements.txt

CMD [ "python3", "ddnsCore.py" ]