package com.baeldung.properties.yaml;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.MockitoJUnitRunner;

import java.util.ArrayList;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

//@RunWith(SpringJUnit4ClassRunner.class)
//@RunWith(SpringRunner.class)
@RunWith(MockitoJUnitRunner.class)
public class YamlFooPropertiesTest {

    @Mock
    YamlFooProperties prop;

    @Before
    public void init() {
        prop = Mockito.mock(YamlFooProperties.class);
        Mockito.when(prop.getName()).thenReturn("foo");
        Mockito.when(prop.getAliases()).thenReturn(List.of("foo", "bar"));
    }

    @Test
    public void fooTest() {
        assertThat(prop.getName()).isEqualTo("foo");
        assertThat(prop.getAliases()).contains("bar", "foo");
    }
}