package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.service.MediaService;
import com.yamatokataoka.fileservice.api.service.StorageService;
import com.yamatokataoka.fileservice.api.repository.MetadataRepository;
import com.yamatokataoka.fileservice.api.domain.Metadata;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.mock.web.MockMultipartFile;

import java.time.LocalDateTime;
import java.util.Optional;

import static com.yamatokataoka.fileservice.api.testBuilder.buildMetadata;
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
    Metadata metadata = buildMetadata("507f1f77bcf86cd799439011", "originalName.txt", 1000L, LocalDateTime.of(2020, 2, 25, 0, 0));
    when(metadataRepository.findById(any())).thenReturn(Optional.of(metadata));

    mediaService.delete("string of id");

    verify(storageService, times(1)).delete(any());
    verify(metadataRepository, times(1)).findById(any());
    verify(metadataRepository, times(1)).delete(any());
  }
}