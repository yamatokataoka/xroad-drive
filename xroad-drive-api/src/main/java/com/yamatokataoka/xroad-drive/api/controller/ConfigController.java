package com.yamatokataoka.xroaddrive.api.controller;

import com.yamatokataoka.xroaddrive.api.XRoadProperties;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/config")
public class ConfigController {

  @Autowired
  private XRoadProperties xRoadProperties;

  @GetMapping
  public XRoadProperties getXRoadConfig() {
    return xRoadProperties;
  }
}
