package com.yamatokataoka.xroaddrive.api.service;

import com.yamatokataoka.xroaddrive.api.service.MediaService;
import com.yamatokataoka.xroaddrive.api.service.StorageService;
import com.yamatokataoka.xroaddrive.api.repository.MetadataRepository;
import com.yamatokataoka.xroaddrive.api.domain.Metadata;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.mock.web.MockMultipartFile;

import java.time.Instant;
import java.util.Optional;

import static com.yamatokataoka.xroaddrive.api.testBuilder.buildMetadata;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class MediaServiceTest {

  @Mock
  private StorageService storageService;

  @Mock
  private MetadataRepository metadataRepository;

  private MediaService mediaService;

  @BeforeEach
  public void setup() {
    mediaService = new MediaService(storageService, metadataRepository);
  }

  @Test
  public void testStore() {
    MockMultipartFile mockMultipartFile = new MockMultipartFile("file", "originalFilename.txt", "text/plain", "some text".getBytes());
    mediaService.store(mockMultipartFile);

    verify(storageService, times(1)).store(any(), any());
    verify(metadataRepository, times(1)).save(any());
  }

  @Test
  public void testLoad() {
    mediaService.load("string of id");

    verify(storageService, times(1)).load(any());
  }

  @Test
  public void testDelete() {
    Metadata metadata = buildMetadata("507f1f77bcf86cd799439011", "originalName.txt", 1000L, Instant.parse("2019-10-01T08:25:24.00Z"));
    when(metadataRepository.findById(any())).thenReturn(Optional.of(metadata));

    mediaService.delete("string of id");

    verify(storageService, times(1)).delete(any());
    verify(metadataRepository, times(1)).findById(any());
    verify(metadataRepository, times(1)).delete(any());
  }
}