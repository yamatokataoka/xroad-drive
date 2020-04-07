package com.yamatokataoka.fileservice.api.domain;

public class File extends BaseEntity {

  private String name;
  private String url;
  private Long size;

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getUrl() {
    return url;
  }

  public void setUrl(String url) {
    this.url = url;
  }

  public Long getSize() {
    return size;
  }

  public void setSize(Long size) {
    this.size = size;
  }
}