FROM mysql:8.0

ENV MYSQL_DATABASE=blops_me
ENV MYSQL_USER=blops_me
ENV MYSQL_PASSWORD=12345678
ENV MYSQL_ROOT_PASSWORD=87654321

WORKDIR /app

COPY ./assets/schemes.sql /docker-entrypoint-initdb.d/

EXPOSE 3306

CMD ["mysqld"]
