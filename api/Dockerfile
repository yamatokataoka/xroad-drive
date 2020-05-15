FROM openjdk:11 AS builder
ENV GRADLE_USER_HOME /fileservice-api/.gradle
COPY . /fileservice-api
WORKDIR /fileservice-api
RUN ./gradlew bootJar

FROM openjdk:11-jre-slim
RUN useradd --shell /bin/bash fileservice-api-user
USER fileservice-api-user
COPY --from=builder /fileservice-api/build/libs/fileservice-api.jar /opt/fileservice/fileservice-api.jar
WORKDIR /etc/fileservice/fileservice-api
EXPOSE 8082
CMD ["java", "-jar", "/opt/fileservice/fileservice-api.jar"]