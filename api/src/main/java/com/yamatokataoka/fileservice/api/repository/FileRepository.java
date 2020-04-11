package com.yamatokataoka.fileservice.api.repository;

import com.yamatokataoka.fileservice.api.domain.File;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface FileRepository extends MongoRepository<File, String> {}