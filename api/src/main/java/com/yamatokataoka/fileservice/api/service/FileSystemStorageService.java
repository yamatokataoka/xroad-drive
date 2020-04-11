package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.StorageProperties;
import com.yamatokataoka.fileservice.api.StorageException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.io.Resource;
import org.springframework.core.io.UrlResource;
import org.springframework.stereotype.Service;
import org.springframework.util.FileSystemUtils;

import java.io.IOException;
import java.io.InputStream;
import java.net.MalformedURLException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;
import java.util.stream.Stream;
import java.util.UUID;

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
    try {
      log.info("Store file: {}", id);
			Files.copy(inputStream, this.location.resolve(id), StandardCopyOption.REPLACE_EXISTING);
		} catch (IOException e) {
      log.error("Failed to store file", e);

      throw new StorageException(e);
		}
	}

	@Override
	public Stream<Path> loadAll() {
    Stream<Path> streams = null;
		try {
			return Files.walk(this.location, 1)
				.filter(path -> !path.equals(this.location))
				.map(this.location::relativize);
		} catch (IOException e) {
      log.error("Failed to read stored files", e);

      throw new StorageException("Failed to read stored files", e);
		}
	}

	@Override
	public Path load(String id) {
		return location.resolve(id);
	}

	@Override
	public Resource loadAsResource(String id) {
    try {
      log.info("Read file: {}", id);
			Path file = load(id);
			Resource resource = new UrlResource(file.toUri());
			if (!resource.exists() || !resource.isReadable()) {
        log.error("Failed to read file or does not exist");
			}
      return resource;
		} catch (MalformedURLException e) {
      log.error("Failed to read file", e);

      throw new StorageException("Failed to read file", e);
		}
	}

	@Override
	public void deleteAll() {
		FileSystemUtils.deleteRecursively(location.toFile());
	}
}