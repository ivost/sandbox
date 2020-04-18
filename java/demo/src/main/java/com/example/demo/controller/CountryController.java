package com.example.demo.controller;

import com.example.demo.model.Country;
import com.example.demo.service.CountryService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Optional;

@RestController
public class CountryController {

    final CountryService countryService;
//    public CountryController() {
//        this.countryService = new CountryServiceImpl();
//    }
public CountryController(CountryService countryService) {
    this.countryService = countryService;
}

    @GetMapping("/countries")
    public List<Country> getCountries() {
        return countryService.findAll();
    }

    @GetMapping("/countries/{countryId}")
    public Optional<Country> getCountry(@PathVariable Long countryId) {
        return countryService.find(countryId);
    }
}
