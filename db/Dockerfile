FROM mysql:latest

ARG admin_psw
ARG admin_user
ARG name
ARG root_psw

# Admin will be granted superuser permissions on the created database.
ENV MYSQL_DATABASE=$name
ENV MYSQL_PASSWORD=$admin_psw
ENV MYSQL_ROOT_PASSWORD=$root_psw
ENV MYSQL_USER=$admin_user

COPY ./setup.sql /docker-entrypoint-initdb.d

EXPOSE 3306