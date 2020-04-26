package com.yamatokataoka.fileservice.api.controller;

import com.yamatokataoka.fileservice.api.service.MediaService;
import com.yamatokataoka.fileservice.api.FileException;
import com.yamatokataoka.fileservice.api.domain.Metadata;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.Resource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.List;

@RestController
@RequestMapping("/api")
public class MediaController {

  @Autowired
  private MediaService mediaService;

  @PostMapping(
    path = "/upload",
    consumes = "multipart/form-data")
  public Metadata upload(@RequestParam("file") List<MultipartFile> files) {
    if (files.size() != 1) {
      throw new FileException("Exact one file is required.");
    }
    Metadata metadata = mediaService.store(files.get(0));
    return metadata;
  }

  @GetMapping("/download/{id}")
  public ResponseEntity<Resource> download(@PathVariable String id) {
    Resource fileResource = mediaService.load(id);

    return ResponseEntity.ok()
      .header(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" + fileResource.getFilename() + "\"")
      .body(fileResource);
  }

  @DeleteMapping("/delete/{id}")
  public ResponseEntity delete(@PathVariable String id) {
    mediaService.delete(id);

    return ResponseEntity.noContent()
      .build();
  }
}