package com.yamatokataoka.fileservice.api.domain;

import org.springframework.data.annotation.Id;

public class File {

  @Id
  private String id;
  private String name;
  private Long size;

  public File() {}

  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public Long getSize() {
    return size;
  }

  public void setSize(Long size) {
    this.size = size;
  }
}