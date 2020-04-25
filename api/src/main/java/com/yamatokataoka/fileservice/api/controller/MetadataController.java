package com.yamatokataoka.fileservice.api.controller;

import com.yamatokataoka.fileservice.api.service.MetadataService;
import com.yamatokataoka.fileservice.api.domain.Metadata;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/metadata")
public class MetadataController {

  @Autowired
  private MetadataService metadataService;
  
  @GetMapping
  public List<Metadata> getMetadataList() {
    return metadataService.getAll();
  }

  @GetMapping("/{id}")
  public Metadata getMetadataById(@PathVariable String id) {
    return metadataService.getById(id);
  }
}