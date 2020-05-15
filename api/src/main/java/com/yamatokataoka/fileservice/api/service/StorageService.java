package com.yamatokataoka.fileservice.api.service;

import org.springframework.core.io.Resource;
import org.springframework.web.multipart.MultipartFile;

import java.io.InputStream;
import java.nio.file.Path;

public interface StorageService {

	void init();

	void store(InputStream inputStream, String id);

	Path resolve(String id);

	Resource load(String id);

	void delete(String id);

}