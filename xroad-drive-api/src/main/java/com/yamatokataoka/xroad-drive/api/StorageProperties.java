package com.yamatokataoka.xroaddrive.api;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "xroaddrive.storage")
public class StorageProperties {

	private String location;

  public String getLocation() {
    return location;
  }

  public void setLocation(String location) {
    this.location = location;
  }
}