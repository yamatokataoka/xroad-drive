package com.yamatokataoka.fileservice.api.controller;

import com.yamatokataoka.fileservice.api.controller.MediaController;
import com.yamatokataoka.fileservice.api.domain.Metadata;
import com.yamatokataoka.fileservice.api.service.MediaService;
import com.yamatokataoka.fileservice.api.service.StorageService;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.mock.web.MockMultipartFile;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;

import static com.yamatokataoka.fileservice.api.testBuilder.buildMetadata;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.multipart;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(MediaController.class)
public class MediaControllerTest {

  @Autowired
  private MockMvc mockMvc;

  @Autowired
  private ObjectMapper objectMapper;

  @MockBean
  private MediaService mediaService;

  @MockBean
  private StorageService storageService;

  @Test
  void testUpload() throws Exception {

    Metadata metadata = buildMetadata("507f1f77bcf86cd799439011", "originalName.txt", 1000L);
    MockMultipartFile mockMultipartFile = new MockMultipartFile("file", "originalFilename.txt", "text/plain", "some text".getBytes());

    when(mediaService.store(any())).thenReturn(metadata);

    MvcResult mvcResult = mockMvc.perform(multipart("/api/upload")
        .file(mockMultipartFile))
      .andExpect(status().isOk())
      .andReturn();

    String actualRsponseAsString = mvcResult.getResponse().getContentAsString();
    Metadata actualResponseBody = objectMapper.readValue(actualRsponseAsString, Metadata.class);

    assertAll(
      () -> assertEquals(metadata.getId(), actualResponseBody.getId()),
      () -> assertEquals(metadata.getFilename(), actualResponseBody.getFilename()),
      () -> assertEquals(metadata.getFilesize(), actualResponseBody.getFilesize())
    );
  }

  @Test
  void testDownload() throws Exception {

    Metadata metadata = buildMetadata("507f1f77bcf86cd799439011", "originalName.txt", 1000L);
    byte[] mockMultipartFileContent = "some text".getBytes();
    MockMultipartFile mockMultipartFile = new MockMultipartFile("file", "originalFilename.txt", "text/plain", mockMultipartFileContent);

    when(storageService.load(any())).thenReturn(mockMultipartFile.getResource());

    MvcResult mvcResult = mockMvc.perform(get("/api/download/{id}", "507f1f77bcf86cd799439011"))
      .andExpect(status().isOk())
      .andReturn();

    byte[] actualRsponseAsByte = mvcResult.getResponse().getContentAsByteArray();

    assertArrayEquals(mockMultipartFileContent, actualRsponseAsByte);
  }
}