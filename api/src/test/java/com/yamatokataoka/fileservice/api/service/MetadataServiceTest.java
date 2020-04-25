package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.service.MetadataService;
import com.yamatokataoka.fileservice.api.repository.MetadataRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;

@ExtendWith(MockitoExtension.class)
public class MetadataServiceTest {

  @Mock
  private MetadataRepository metadataRepository;

  private MetadataService metadataService;

  @BeforeEach
  public void setup() {
    metadataService = new MetadataService(metadataRepository);
  }

  @Test
  public void testGetAll() {
    
    metadataService.getAll();

    verify(metadataRepository, times(1)).findAll();
  }

  @Test
  public void testGetById() {
    
    metadataService.getById("string of id");

    verify(metadataRepository, times(1)).findById(any());
  }
}