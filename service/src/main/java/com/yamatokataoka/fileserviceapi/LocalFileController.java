package com.yamatokataoka.fileserviceapi;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.core.io.Resource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpStatus;
import org.springframework.beans.factory.annotation.Autowired;

import com.yamatokataoka.fileserviceapi.storage.StorageException;
import com.yamatokataoka.fileserviceapi.storage.StorageService;

@RestController
@RequestMapping("/api")
public class LocalFileController {

  private final StorageService storageService;

  @Autowired
	public LocalFileController(StorageService storageService) {
		this.storageService = storageService;
	}

  @PostMapping("/upload")
  public ResponseEntity<?> upload(@RequestParam("file") MultipartFile[] files) {
    try {
      for (MultipartFile file : files) {
        storageService.store(file);
      }
    } catch (StorageException se) {
      // se.printStackTrace();
      return new ResponseEntity<>(se.getMessage(), HttpStatus.INTERNAL_SERVER_ERROR);
    }
    return new ResponseEntity<>(HttpStatus.OK);
  }

  @GetMapping("/download/{id}")
  public ResponseEntity<Resource> download(@PathVariable String id) {
    Resource file = storageService.loadAsResource(id);

    return ResponseEntity.ok()
      .header(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" 
        + file.getFilename() + "\"")
      .body(file);
  }
}