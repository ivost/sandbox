package com.example.demo;

import org.springframework.transaction.annotation.Transactional;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
//@ContextConfiguration({"application.yml"})
//@ActiveProfiles("dev")
@Transactional
public @interface TransactionalDevTest { }
