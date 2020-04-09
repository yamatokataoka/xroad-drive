package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.domain.File;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import java.io.IOException;
import java.io.InputStream;
import java.util.UUID;

@Service
public class FileService {

  private static final Logger log = LoggerFactory.getLogger(FileService.class);
  private final StorageService storageService;

  public FileService(StorageService storageService) {
    this.storageService = storageService;
  }

  public File store(MultipartFile multipartFile) {
    String fileId = UUID.randomUUID().toString();
		String originalFilename = StringUtils.cleanPath(multipartFile.getOriginalFilename());
		String url = ServletUriComponentsBuilder.fromCurrentContextPath()
			.path("/download/")
			.path(fileId)
			.toUriString();
    Long size = multipartFile.getSize();

    File file = new File();
    file.setName(originalFilename);
    file.setUrl(url);
    file.setSize(size);

    try (InputStream inputStream = multipartFile.getInputStream()) {
      storageService.store(inputStream, fileId);
    } catch (IOException e) {
      log.error("Failed to read file", e);
		}
    return file;
  }
}