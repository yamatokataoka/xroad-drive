package com.yamatokataoka.client;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.core.io.FileSystemResource;
import org.springframework.core.io.Resource;
import org.springframework.context.annotation.Bean;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

@SpringBootApplication
public class ClientApplication {

	public static void main(String[] args) {
		SpringApplication.run(ClientApplication.class, args);
	}

  @Bean
	public RestTemplate restTemplate(RestTemplateBuilder builder) {
		return builder.build();
	}

  @Bean
	public CommandLineRunner run(RestTemplate restTemplate) throws Exception {
    HttpHeaders headers = new HttpHeaders();
    headers.setContentType(MediaType.MULTIPART_FORM_DATA);

    MultiValueMap<String, Object> body = new LinkedMultiValueMap<>();
    body.add("file", getTestFile());
    body.add("file", getTestFile());
    body.add("file", getTestFile());

    HttpEntity<MultiValueMap<String, Object>> requestEntity = new HttpEntity<>(body, headers);
    String serverUrl = "http://localhost:8080/api/upload";

		return args -> {
			ResponseEntity<String> response = restTemplate
        .postForEntity(serverUrl, requestEntity, String.class);
      System.out.println(response);
		};
	}

  public static Resource getTestFile() throws IOException {
      Path testFile = Files.createTempFile("test-file", ".txt");
      System.out.println("Creating Test File: " + testFile);
      Files.write(testFile, "Hello World !!, This is a test file.".getBytes());
      return new FileSystemResource(testFile.toFile());
    }
}
