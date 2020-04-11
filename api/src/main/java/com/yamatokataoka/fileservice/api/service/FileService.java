package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.domain.File;
import com.yamatokataoka.fileservice.api.repository.FileRepository;
import com.yamatokataoka.fileservice.api.StorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.io.InputStream;
import java.util.UUID;

@Service
public class FileService {

  private static final Logger log = LoggerFactory.getLogger(FileService.class);
  private final StorageService storageService;
  private final FileRepository fileRepository;

  public FileService(StorageService storageService, FileRepository fileRepository) {
    this.storageService = storageService;
    this.fileRepository = fileRepository;
  }

  public File store(MultipartFile multipartFile) {
    String fileId = UUID.randomUUID().toString();
		String originalFilename = StringUtils.cleanPath(multipartFile.getOriginalFilename());
    Long size = multipartFile.getSize();

    File file = new File();
    file.setName(originalFilename);
    file.setSize(size);

    try (InputStream inputStream = multipartFile.getInputStream()) {
      storageService.store(inputStream, fileId);
      fileRepository.save(file);
      return file;
    } catch (Exception e) {
      throw new StorageException("Failed to store file", e);
		}
  }
}