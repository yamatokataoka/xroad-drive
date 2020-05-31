package com.yamatokataoka.xroaddrive.api.service;

import com.yamatokataoka.xroaddrive.api.service.MetadataService;
import com.yamatokataoka.xroaddrive.api.repository.MetadataRepository;
import com.yamatokataoka.xroaddrive.api.domain.Metadata;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.time.Instant;
import java.util.Optional;

import static com.yamatokataoka.xroaddrive.api.testBuilder.buildMetadata;
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
    Metadata metadata = buildMetadata("507f1f77bcf86cd799439011", "originalName.txt", 1000L, Instant.parse("2019-10-01T08:25:24.00Z"));
    when(metadataRepository.findById(any())).thenReturn(Optional.of(metadata));
    
    metadataService.getById("507f1f77bcf86cd799439011");

    verify(metadataRepository, times(1)).findById(any());
  }
}