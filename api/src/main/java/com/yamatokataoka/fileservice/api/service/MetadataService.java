package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.domain.Metadata;
import com.yamatokataoka.fileservice.api.repository.MetadataRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.Arrays;
import java.util.List;

@Service
public class MetadataService {

  private static final Logger log = LoggerFactory.getLogger(MetadataService.class);
  private final MetadataRepository metadataRepository;

  public MetadataService(MetadataRepository metadataRepository) {
    this.metadataRepository = metadataRepository;
  }

  public List<Metadata> getAll() {
    return metadataRepository.findAll();
  }

  public Metadata getById(String id) {
    return metadataRepository.findById(id).orElseThrow();
  }
}