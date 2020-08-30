package com.yamatokataoka.xroaddrive.api.controller;

import com.yamatokataoka.xroaddrive.api.service.MetadataService;
import com.yamatokataoka.xroaddrive.api.controller.MetadataController;
import com.yamatokataoka.xroaddrive.api.domain.Metadata;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.core.type.TypeReference;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;

import java.time.Instant;
import java.util.Arrays;
import java.util.List;

import static com.yamatokataoka.xroaddrive.api.testBuilder.buildMetadata;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(MetadataController.class)
public class MetadataControllerTest {

  @Autowired
  private MockMvc mockMvc;

  @Autowired
  private ObjectMapper objectMapper;

  @MockBean
  private MetadataService metadataService;

  private Metadata metadata = buildMetadata("507f1f77bcf86cd799439011", "originalName.txt", 1000L, Instant.parse("2019-10-01T08:25:24.00Z"));

  @Test
  void testGetMetadataList() throws Exception {
    
    when(metadataService.getAll()).thenReturn(Arrays.asList(metadata));

    MvcResult mvcResult = mockMvc.perform(get("/api/metadata"))
      .andExpect(status().isOk())
      .andReturn();

    TypeReference<List<Metadata>> typeReference = new TypeReference<List<Metadata>>() {};
    String actualRsponseAsString = mvcResult.getResponse().getContentAsString();
    List<Metadata> actualResponseBody = objectMapper.readValue(actualRsponseAsString, typeReference);
    Metadata actualResponseMetadata = actualResponseBody.get(0);

    assertAll(
      () -> assertEquals(metadata.getId(), actualResponseMetadata.getId()),
      () -> assertEquals(metadata.getFilename(), actualResponseMetadata.getFilename()),
      () -> assertEquals(metadata.getFilesize(), actualResponseMetadata.getFilesize()),
      () -> assertEquals(metadata.getCreatedDateTime(), actualResponseMetadata.getCreatedDateTime())
    );
  }

  @Test
  void testGetMetadataById() throws Exception {
    
    when(metadataService.getById(any())).thenReturn(metadata);

    MvcResult mvcResult = mockMvc.perform(get("/api/metadata/{id}", "507f1f77bcf86cd799439011"))
      .andExpect(status().isOk())
      .andReturn();

    String actualRsponseAsString = mvcResult.getResponse().getContentAsString();
    Metadata actualResponseBody = objectMapper.readValue(actualRsponseAsString, Metadata.class);

    assertAll(
      () -> assertEquals(metadata.getId(), actualResponseBody.getId()),
      () -> assertEquals(metadata.getFilename(), actualResponseBody.getFilename()),
      () -> assertEquals(metadata.getFilesize(), actualResponseBody.getFilesize()),
      () -> assertEquals(metadata.getCreatedDateTime(), actualResponseBody.getCreatedDateTime())
    );
  }
}