package com.yamatokataoka.fileservice.api;

import com.yamatokataoka.fileservice.api.domain.Metadata;

public class testBuilder {
  public static Metadata buildMetadata(String id, String filename, Long filesize) {
    Metadata metadata = new Metadata();
    metadata.setId(id);
    metadata.setFilename(filename);
    metadata.setFilesize(filesize);
    return metadata;
  }
}