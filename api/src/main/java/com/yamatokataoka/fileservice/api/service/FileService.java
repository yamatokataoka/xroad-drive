package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.domain.File;
import com.yamatokataoka.fileservice.api.StorageException;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import java.io.IOException;
import java.io.InputStream;
import java.util.UUID;

@Service
public class FileService {

  private final StorageService storageService;

  public FileService(StorageService storageService) {
    this.storageService = storageService;
  }

  public void store(MultipartFile multipartFile) {
    String id = UUID.randomUUID().toString();
		String originalFilename = StringUtils.cleanPath(multipartFile.getOriginalFilename());
		String url = ServletUriComponentsBuilder.fromCurrentContextPath()
			.path("/download/")
			.path(id)
			.toUriString();

    try(InputStream inputStream = multipartFile.getInputStream()) {
      File file = new File();
      file.setId(id);
      file.setName(originalFilename);
      file.setSize(multipartFile.getSize());
  
      storageService.store(inputStream, id);
    } catch (IOException e) {
			throw new StorageException("Failed to store file " + id, e);
		}
  }
}