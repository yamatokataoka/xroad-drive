package com.yamatokataoka.fileservice.api.service;

import com.yamatokataoka.fileservice.api.StorageException;
import com.yamatokataoka.fileservice.api.StorageFileNotFoundException;
import com.yamatokataoka.fileservice.api.StorageProperties;
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

	private final Path location;

	public FileSystemStorageService(StorageProperties properties) {
		location = Paths.get(properties.getLocation());
	}

	@Override
	public void init() {
		try {
			Files.createDirectories(location);
		} catch (IOException e) {
			throw new StorageException("Failed to initialize file location", e);
		}
	}

	@Override
	public void store(InputStream inputStream, String id) {
    try {
			Files.copy(inputStream, this.location.resolve(id), StandardCopyOption.REPLACE_EXISTING);
		} catch (IOException e) {
			throw new StorageException("Failed to store file " + id, e);
		}
	}

	@Override
	public Stream<Path> loadAll() {
		try {
			return Files.walk(this.location, 1)
				.filter(path -> !path.equals(this.location))
				.map(this.location::relativize);
		} catch (IOException e) {
			throw new StorageException("Failed to read files", e);
		}
	}

	@Override
	public Path load(String id) {
		return location.resolve(id);
	}

	@Override
	public Resource loadAsResource(String id) {
		try {
			Path file = load(id);
			Resource resource = new UrlResource(file.toUri());
			if (resource.exists() || resource.isReadable()) {
				return resource;
			} else {
				throw new StorageFileNotFoundException("Failed to read file: " + id);
			}
		} catch (MalformedURLException e) {
			throw new StorageFileNotFoundException("Could not read file: " + id, e);
		}
	}

	@Override
	public void deleteAll() {
		FileSystemUtils.deleteRecursively(location.toFile());
	}
}