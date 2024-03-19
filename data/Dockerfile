FROM mongo:latest
WORKDIR /json_data

COPY ./json_data/ .
COPY ./seed.sh .

RUN chmod +x seed.sh

CMD ["sh", "-c", "/json_data/seed.sh"]