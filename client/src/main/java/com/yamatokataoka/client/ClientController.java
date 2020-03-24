package com.yamatokataoka.client;

import org.springframework.core.io.FileSystemResource;
import org.springframework.core.io.Resource;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

@RestController
@RequestMapping("/api")
public class ClientController {
  
  @PostMapping("/upload")
  public ResponseEntity<?> upload() {
    RestTemplate restTemplate = new RestTemplate();
    
    HttpHeaders headers = new HttpHeaders();
    headers.setContentType(MediaType.MULTIPART_FORM_DATA);

    String serverUrl = "http://localhost:8080/api/upload";

    MultiValueMap<String, Object> body = new LinkedMultiValueMap<>();

    try {
      body.add("file", getTestFile());
      body.add("file", getTestFile());
      body.add("file", getTestFile());
    } catch (IOException ie) {
      ie.printStackTrace();
      return new ResponseEntity<>(ie.getMessage(), HttpStatus.INTERNAL_SERVER_ERROR);
    }

    HttpEntity<MultiValueMap<String, Object>> requestEntity = new HttpEntity<>(body, headers);

    return restTemplate.postForEntity(serverUrl, requestEntity, String.class);
  }

  public Resource getTestFile() throws IOException {
    Path testFile = Files.createTempFile("test-file", ".txt");
    System.out.println("Creating Test File: " + testFile);
    Files.write(testFile, "Hello World !!, This is a test file.".getBytes());
    return new FileSystemResource(testFile.toFile());
  }
}