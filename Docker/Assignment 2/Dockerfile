FROM openjdk:latest
VOLUME /tmp
ARG JAVA_OPTS
ENV JAVA_OPTS=$JAVA_OPTS
COPY classroom-0.0.1-SNAPSHOT.jar devops.jar
EXPOSE 12346
ENTRYPOINT exec java $JAVA_OPTS -jar devops.jar
