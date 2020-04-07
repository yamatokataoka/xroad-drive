package com.yamatokataoka.fileservice.api.repository;

import com.yamatokataoka.fileservice.api.domain.File;
import org.springframework.data.repository.CrudRepository;

public interface FileRepository extends CrudRepository<File, Long> {}