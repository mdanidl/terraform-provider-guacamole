FROM mysql:latest
ADD config/my.cnf /etc/mysql/my.cnf
ADD config/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf
RUN mkdir -p /var/log/mysql && chown mysql:mysql /var/log/mysql
COPY schema/00* /docker-entrypoint-initdb.d/
