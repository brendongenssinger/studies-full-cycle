FROM mysql:5.7-debian



#RUN apt-get update && apt-get install -y vim || apt-get update -o Acquire::ForceIPv4=true && apt-get install -y vim


ENV MYSQL_ROOT_PASSWORD=DevBrendon!

# Define the entrypoint to run the MySQL server
ENTRYPOINT ["docker-entrypoint.sh"]

# The default command that runs when the container starts
CMD ["mysqld"]