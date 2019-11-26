FROM maven:alpine as build-env

WORKDIR /code

# Prepare by downloading dependencies
ADD pom.xml /code/pom.xml
RUN ["mvn", "dependency:resolve"]
RUN ["mvn", "verify"]

# Adding source, compile and package into a fat jar
ADD src /code/src
RUN ["mvn", "package"]

FROM openjdk:8-jre-alpine

WORKDIR /app
COPY --from=build-env /code/target/hello-jar-with-dependencies.jar /app/app.jar
CMD ["java", "-jar", "app.jar"]
