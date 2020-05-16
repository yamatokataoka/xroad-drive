package com.yamatokataoka.xroaddrive.api.service;

import com.yamatokataoka.xroaddrive.api.StorageProperties;
import com.yamatokataoka.xroaddrive.api.FileException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.io.Resource;
import org.springframework.core.io.UrlResource;
import org.springframework.stereotype.Service;
import org.springframework.util.FileSystemUtils;

import java.io.IOException;
import java.io.FileOutputStream;
import java.io.InputStream;
import java.net.MalformedURLException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

@Service
public class FileSystemStorageService implements StorageService {

  private static final Logger log = LoggerFactory.getLogger(FileSystemStorageService.class);
	private final Path location;

	public FileSystemStorageService(StorageProperties properties) {
		location = Paths.get(properties.getLocation());
	}

	@Override
	public void init() {
		try {
      log.info("Initialize file location");
			Files.createDirectories(location);
		} catch (IOException e) {
      log.error("Failed to initialize file location", e);
		}
	}

	@Override
	public void store(InputStream inputStream, String id) {
    try(FileOutputStream fileOutputStream = new FileOutputStream(this.location.resolve(id).toFile())) {
      log.info("Store file: {}", id);
      inputStream.transferTo(fileOutputStream);
		} catch (IOException e) {
      log.error("Failed to store file", e);

      throw new FileException(e);
		}
	}

	@Override
	public Path resolve(String id) {
		return location.resolve(id);
	}

	@Override
	public Resource load(String id) {
    log.info("Load file: {}", id);

    try {
			Path file = resolve(id);
			Resource resource = new UrlResource(file.toUri());
			if (!resource.exists() || !resource.isReadable()) {
        log.error("Failed to load file or does not exist");
			}
      return resource;
		} catch (MalformedURLException e) {
      log.error("Failed to load file", e);

      throw new FileException("Failed to load file", e);
		}
	}

	@Override
	public void delete(String id) {
    log.info("Delete file: {}", id);
		
    try {
      Files.deleteIfExists(resolve(id));
    } catch (Exception e) {
      log.error("Failed to delete file", e);

      throw new FileException("Failed to delete file", e);
    }
	}
}