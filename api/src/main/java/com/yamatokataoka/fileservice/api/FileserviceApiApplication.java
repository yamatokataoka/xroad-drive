package com.yamatokataoka.fileservice.api;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;

import com.yamatokataoka.fileservice.api.service.StorageService;

@SpringBootApplication
@EnableConfigurationProperties(StorageProperties.class)
public class FileserviceApiApplication {

	public static void main(String[] args) {
		SpringApplication.run(FileserviceApiApplication.class, args);
	}

	@Bean
	CommandLineRunner init(StorageService storageService) {
		return (args) -> {
			storageService.init();
		};
	}
}
