package com.example.demo.service;

import com.example.demo.model.Country;

import java.util.List;
import java.util.Optional;

public interface CountryService {
    List<Country> findAll();
    Optional<Country> find(Long countryId);
}
