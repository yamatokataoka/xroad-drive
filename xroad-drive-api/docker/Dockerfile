FROM openjdk:11 AS develop
ENV GRADLE_USER_HOME /xroad-drive-api/.gradle
COPY . /xroad-drive-api
WORKDIR /xroad-drive-api

FROM develop AS build
RUN ./gradlew bootJar

FROM openjdk:11-jre-slim AS production
RUN useradd --shell /bin/bash xroad-drive-api-user
USER xroad-drive-api-user
COPY --from=build /xroad-drive-api/build/libs/xroad-drive-api.jar /opt/xroad-drive/xroad-drive-api.jar
WORKDIR /etc/xroad-drive/xroad-drive-api
EXPOSE 8082
CMD ["java", "-jar", "/opt/xroad-drive/xroad-drive-api.jar"]