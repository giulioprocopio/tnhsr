FROM mysql:8.0

ARG db_admin_psw
ARG db_admin_user
ARG db_name
ARG db_root_psw

# Admin will be granted superuser permissions on the created database.
ENV MYSQL_DATABASE=$db_name
ENV MYSQL_PASSWORD=$db_admin_psw
ENV MYSQL_ROOT_PASSWORD=$db_root_psw
ENV MYSQL_USER=$db_admin_user

COPY ./deploy/db/setup.sql /docker-entrypoint-initdb.d

EXPOSE 3306