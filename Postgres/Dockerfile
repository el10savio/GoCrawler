FROM library/postgres:12.1

COPY Postgres/init.sql /docker-entrypoint-initdb.d/
RUN chmod u+x /docker-entrypoint-initdb.d/init.sql

RUN apt-get update && apt-get -y install sudo

RUN unlink /docker-entrypoint.sh
RUN ln -s usr/local/bin/docker-entrypoint.sh /

#Expose Port 5432
EXPOSE 5432/tcp

# Allow connections from anywhere
RUN echo "host    all    all    0.0.0.0/0    md5" >> /var/lib/postgresql/data/pg_hba.conf

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["postgres"]
