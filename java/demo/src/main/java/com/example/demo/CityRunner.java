package com.example.demo;

import com.example.demo.repository.CityRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
@Slf4j
public class CityRunner implements CommandLineRunner {

    private final CityRepository cityRepository;
    public CityRunner(CityRepository cityRepository) {
        this.cityRepository = cityRepository;
    }

    @Override
    public void run(String... args) throws Exception {

//        log.info("=== CityRunner run");
//        var cities = cityRepository.findAll();
//        cities.forEach(city -> log.info("{}", city));
//
//        cities = cityRepository.findByNameEndingWithAndPopulationLessThan("es", 4_000_000);
//        cities.forEach(city -> log.info("{}", city));

    }
}
