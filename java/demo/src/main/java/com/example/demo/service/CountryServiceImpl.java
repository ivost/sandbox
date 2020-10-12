package com.example.demo.service;

import com.example.demo.model.Country;
import com.example.demo.repository.CountryRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
@Slf4j
public class CountryServiceImpl implements CountryService {

    private final CountryRepository repository;

    public CountryServiceImpl(CountryRepository repository) {

        this.repository = repository;
    }

    @Override
    public List<Country> findAll() {
        log.debug("findAll");
        return (List<Country>) repository.findAll();
    }

    @Override
    public Optional<Country> find(Long countryId) {
        log.debug("find {}", countryId);
        return repository.findById(countryId);
    }
}
