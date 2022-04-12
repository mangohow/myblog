FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=abc_123

#定义会被容器自动执行的目录
ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d
# 定义要执行的shell脚本
ENV CREATE_DB_SH create_db.sh

# 拷贝sql文件和脚本文件
COPY blog.sql /blog.sql
COPY $CREATE_DB_SH $AUTO_RUN_DIR
# 赋予脚本执行权限
RUN chmod a+x $AUTO_RUN_DIR/$CREATE_DB_SH

EXPOSE 3306 3306

CMD [ "mysqld" ]
