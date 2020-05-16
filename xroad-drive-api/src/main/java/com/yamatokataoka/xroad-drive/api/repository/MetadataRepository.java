package com.yamatokataoka.xroaddrive.api.repository;

import com.yamatokataoka.xroaddrive.api.domain.Metadata;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface MetadataRepository extends MongoRepository<Metadata, String> {}