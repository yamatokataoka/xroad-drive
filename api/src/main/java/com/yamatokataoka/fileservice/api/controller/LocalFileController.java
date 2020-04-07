package com.yamatokataoka.fileservice.api.controller;

import com.yamatokataoka.fileservice.api.service.FileService;
import com.yamatokataoka.fileservice.api.service.StorageService;
import com.yamatokataoka.fileservice.api.StorageException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.Resource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

@RestController
@RequestMapping("/api")
public class LocalFileController {

  @Autowired
  private FileService fileService;

  @Autowired
  private StorageService storageService;

  @PostMapping("/upload")
  public ResponseEntity<?> upload(@RequestParam("file") MultipartFile[] files) {
    try {
      for (MultipartFile file : files) {
        fileService.store(file);
      }
    } catch (StorageException se) {
      return new ResponseEntity<>(se.getMessage(), HttpStatus.INTERNAL_SERVER_ERROR);
    }
    return new ResponseEntity<>(HttpStatus.OK);
  }

  @GetMapping("/download/{id}")
  public ResponseEntity<Resource> download(@PathVariable String id) {
    Resource file = storageService.loadAsResource(id);

    return ResponseEntity.ok()
      .header(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" + file.getFilename() + "\"")
      .body(file);
  }
}