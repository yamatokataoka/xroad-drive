package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.domain.File;
import com.yamatokataoka.fileservice.api.repository.FileRepository;
import com.yamatokataoka.fileservice.api.FileException;
import org.bson.types.ObjectId;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.io.InputStream;

@Service
public class FileService {

  private static final Logger log = LoggerFactory.getLogger(FileService.class);
  private final StorageService storageService;
  private final FileRepository fileRepository;

  public FileService(StorageService storageService, FileRepository fileRepository) {
    this.storageService = storageService;
    this.fileRepository = fileRepository;
  }

  @Transactional
  public File store(MultipartFile multipartFile) {
    String id = new ObjectId().toString() ;
		String originalFilename = StringUtils.cleanPath(multipartFile.getOriginalFilename());
    Long size = multipartFile.getSize();

    File file = new File();
    file.setId(id);
    file.setName(originalFilename);
    file.setSize(size);

    try (InputStream inputStream = multipartFile.getInputStream()) {
      storageService.store(inputStream, id);
      fileRepository.save(file);
      return file;
    } catch (Exception e) {
      throw new FileException("Failed to store file", e);
		}
  }
}