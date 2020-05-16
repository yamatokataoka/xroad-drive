package com.yamatokataoka.xroaddrive.api.service;

import com.yamatokataoka.xroaddrive.api.domain.Metadata;
import com.yamatokataoka.xroaddrive.api.repository.MetadataRepository;
import com.yamatokataoka.xroaddrive.api.FileException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.io.Resource;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.io.InputStream;

@Service
public class MediaService {

  private static final Logger log = LoggerFactory.getLogger(MediaService.class);
  private final StorageService storageService;
  private final MetadataRepository metadataRepository;

  public MediaService(StorageService storageService, MetadataRepository metadataRepository) {
    this.storageService = storageService;
    this.metadataRepository = metadataRepository;
  }

  @Transactional
  public Metadata store(MultipartFile multipartFile) {
		String originalFilename = StringUtils.cleanPath(multipartFile.getOriginalFilename());
    Long fileSize = multipartFile.getSize();

    Metadata metadata = new Metadata();
    metadata.setFilename(originalFilename);
    metadata.setFilesize(fileSize);

    try (InputStream inputStream = multipartFile.getInputStream()) {
      metadataRepository.save(metadata);
      storageService.store(inputStream, metadata.getId());
      return metadata;
    } catch (Exception e) {
      throw new FileException("Failed to store file", e);
		}
  }

  @Transactional
  public Resource load(String id) {
    try {
      return storageService.load(id);
    } catch (FileException e) {
      throw new FileException("Failed to load file", e);
		}
  }

  @Transactional
  public void delete(String id) {
    Metadata metadata = metadataRepository.findById(id).orElseThrow();
    storageService.delete(id);
    metadataRepository.delete(metadata);
  }
}