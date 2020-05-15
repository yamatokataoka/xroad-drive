package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.service.MetadataService;
import com.yamatokataoka.fileservice.api.repository.MetadataRepository;
import com.yamatokataoka.fileservice.api.domain.Metadata;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.time.LocalDateTime;
import java.util.Optional;

import static com.yamatokataoka.fileservice.api.testBuilder.buildMetadata;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

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
    Metadata metadata = buildMetadata("507f1f77bcf86cd799439011", "originalName.txt", 1000L, LocalDateTime.of(2020, 2, 25, 0, 0));
    when(metadataRepository.findById(any())).thenReturn(Optional.of(metadata));
    
    metadataService.getById("507f1f77bcf86cd799439011");

    verify(metadataRepository, times(1)).findById(any());
  }
}