package com.baeldung.properties.yaml;

import com.baeldung.properties.yaml.factory.YamlPropertySourceFactory;
import lombok.*;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.PropertySource;
import org.springframework.stereotype.Component;

import java.util.List;

@Data
@Component
@Configuration
@ConfigurationProperties(prefix = "yaml")
//@PropertySource(value = "classpath:foo.yml", factory = YamlPropertySourceFactory.class)
@PropertySource(value = "classpath:foo.yml")
public class YamlFooProperties {
    private String name;
    private List<String> aliases;
}
