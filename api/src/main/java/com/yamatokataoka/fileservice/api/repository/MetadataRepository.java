package com.yamatokataoka.fileservice.api.repository;

import com.yamatokataoka.fileservice.api.domain.Metadata;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface MetadataRepository extends MongoRepository<Metadata, String> {}