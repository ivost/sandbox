package com.example.demo2.model;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository(value = "CityRepository2")
/////////////////////////////////////
//@Repository
public interface CityRepository extends CrudRepository<City, Long> {

}
