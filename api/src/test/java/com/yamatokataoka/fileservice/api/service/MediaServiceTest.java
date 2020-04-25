package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.service.MediaService;
import com.yamatokataoka.fileservice.api.service.StorageService;
import com.yamatokataoka.fileservice.api.repository.MetadataRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.mock.web.MockMultipartFile;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;

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
    MockMultipartFile mockMultipartFile = new MockMultipartFile("name", "originalFilename.txt", "text/plain", "some text".getBytes());
    mediaService.store(mockMultipartFile);

    verify(storageService, times(1)).store(any(), any());
    verify(metadataRepository, times(1)).save(any());
  }
}