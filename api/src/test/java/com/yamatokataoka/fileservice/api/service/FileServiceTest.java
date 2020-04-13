package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.service.FileService;
import com.yamatokataoka.fileservice.api.service.StorageService;
import com.yamatokataoka.fileservice.api.repository.FileRepository;
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
class FileServiceTest {

  @Mock
  private StorageService storageService;

  @Mock
  private FileRepository fileRepository;

  private FileService fileservice;

  @BeforeEach
  public void setup() {
    fileservice = new FileService(storageService, fileRepository);
  }

  @Test
  public void testStoreSuccess() {
    MockMultipartFile mockMultipartFile = new MockMultipartFile("name", "originalFilename.txt", "text/plain", "some text".getBytes());
    fileservice.store(mockMultipartFile);

    verify(storageService, times(1)).store(any(), any());
    verify(fileRepository, times(1)).save(any());
  }
}