package com.yamatokataoka.fileservice.api;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;

@SpringBootApplication
@EnableConfigurationProperties(StorageProperties.class)
public class FileserviceApiApplication {

	public static void main(String[] args) {
		SpringApplication.run(FileserviceApiApplication.class, args);
	}
}
