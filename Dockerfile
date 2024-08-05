FROM brendonmascarenhas/mysql_dump


ENV max_allowed_packet=1G

ENV MYSQL_ROOT_PASSWORD=DevBrendon!

# Define the entrypoint to run the MySQL server
ENTRYPOINT ["docker-entrypoint.sh"]

# The default command that runs when the container starts
CMD ["mysqld"]