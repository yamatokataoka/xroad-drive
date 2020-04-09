package com.yamatokataoka.fileservice.api.controller;

import com.yamatokataoka.fileservice.api.service.FileService;
import com.yamatokataoka.fileservice.api.service.StorageService;
import com.yamatokataoka.fileservice.api.domain.File;
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

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/api")
public class LocalFileController {

  @Autowired
  private FileService fileService;

  @Autowired
  private StorageService storageService;

  @PostMapping("/upload")
  public List<File> upload(@RequestParam("file") MultipartFile[] files) {
    List<File> fileList = new ArrayList<File>();
    for (MultipartFile multipartFile : files) {
      File file = fileService.store(multipartFile);
      fileList.add(file);
    }
    return fileList;
  }

  @GetMapping("/download/{id}")
  public ResponseEntity<Resource> download(@PathVariable String id) {
    Resource file = storageService.loadAsResource(id);

    return ResponseEntity.ok()
      .header(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" + file.getFilename() + "\"")
      .body(file);
  }
}