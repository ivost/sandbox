package com.baeldung.properties.yamlmap.pojo;

import lombok.Data;
import lombok.Getter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.PropertySource;
import org.springframework.stereotype.Component;

import java.util.Date;
import java.util.List;
import java.util.Map;

@Data
@Component
@ConfigurationProperties(prefix = "atom")
// @PropertySource("classpath:/com/${my.placeholder:default/path}/app.properties")
// Assuming that "my.placeholder" is present in one of the property sources already registered,
// e.g. system properties or environment variables,
// the placeholder will be resolved to the corresponding value.
// If not, then "default/path" will be used as a default.
// Expressing a default value (delimited by colon ":") is optional.
// If no default is specified and a property cannot be resolved,
// an IllegalArgumentException will be thrown.
// In cases where a given property key exists in more than one .properties file,
// the last @PropertySource annotation processed will 'win' and override.
//@PropertySource(value = "classpath:atom.yml", factory = YamlPropertySourceFactory.class)
@PropertySource(value = "classpath:atom.yml")
public class ServerProperties {
    /*
    In certain situations, it may not be possible or practical to tightly control property source ordering
    when using @PropertySource annotations.
    For example, if the @Configuration classes above were registered via component-scanning,
    the ordering is difficult to predict.
    In such cases - and if overriding is important - it is recommended that the user
    fall back to using the programmatic PropertySource API. See ConfigurableEnvironment and MutablePropertySources javadocs for details.
     */
    private Map<String, String> application;
    private Map<String, List<String>> config;
    private Map<String, Credential> users;
    private Date time;
    private Boolean flag;

    @Getter  //@Setter
    public static class Credential {
        private String username;
        private String password;
        private int num;
        private boolean flag;
    }
}