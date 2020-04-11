package com.yamatokataoka.fileservice.api.service;

import org.springframework.core.io.Resource;
import org.springframework.web.multipart.MultipartFile;

import java.io.InputStream;
import java.nio.file.Path;
import java.util.stream.Stream;

public interface StorageService {

	void init();

	void store(InputStream inputStream, String id);

	Stream<Path> loadAll();

	Path resolve(String id);

	Resource loadAsResource(String id);

	void deleteAll();

}