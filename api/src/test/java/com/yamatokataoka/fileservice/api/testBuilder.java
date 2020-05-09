package com.yamatokataoka.fileservice.api;

import com.yamatokataoka.fileservice.api.domain.Metadata;

import java.time.LocalDateTime;

public class testBuilder {
  public static Metadata buildMetadata(String id, String filename, Long filesize, LocalDateTime createdDateTime) {
    Metadata metadata = new Metadata();
    metadata.setId(id);
    metadata.setFilename(filename);
    metadata.setFilesize(filesize);
    metadata.setCreatedDateTime(createdDateTime);
    return metadata;
  }
}