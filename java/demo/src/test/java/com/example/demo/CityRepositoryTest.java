package com.example.demo;

import com.example.demo.model.City;
import com.example.demo.repository.CityRepository;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.boot.test.autoconfigure.orm.jpa.TestEntityManager;
import org.springframework.test.context.junit4.SpringRunner;

import javax.transaction.Transactional;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertEquals;

@RunWith(SpringRunner.class)
@DataJpaTest
public class CityRepositoryTest {

    @Autowired
    private TestEntityManager entityManager;

    @Autowired
    private CityRepository repository;

    @Test
    @Transactional
    public void testFindAll() {
        var city = new City("New York", 8_500_000);
        entityManager.persist(city);
        var cities = (List<City>) repository.findAll();
        assertEquals(1, cities.size());
        assertThat(cities).extracting(City::getName).containsOnly("New York");
    }
}