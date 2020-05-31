package com.yamatokataoka.xroaddrive.api;

import com.yamatokataoka.xroaddrive.api.domain.Metadata;

import java.time.Instant;

public class testBuilder {
  public static Metadata buildMetadata(String id, String filename, Long filesize, Instant createdDateTime) {
    Metadata metadata = new Metadata();
    metadata.setId(id);
    metadata.setFilename(filename);
    metadata.setFilesize(filesize);
    metadata.setCreatedDateTime(createdDateTime);
    return metadata;
  }
}