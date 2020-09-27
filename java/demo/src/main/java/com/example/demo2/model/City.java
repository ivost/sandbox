package com.example.demo2.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;

/*
Caused by: org.hibernate.DuplicateMappingException: The [com.example.demo.model.City] and [com.example.demo2.model.City] entities share the same JPA entity name: [City] which is not allowed!
	at org.hibernate.boot.internal.InFlightMetadataCollectorImpl.addEntityBinding(InFlightMetadataCollectorImpl.java:314)
	at org.hibernate.cfg.AnnotationBinder.bindClass(AnnotationBinder.java:820)
	at org.hibernate.boot.model.source.internal.annotations.AnnotationMetadataSourceProcessorImpl.processEntityHierarchies(AnnotationMetadataSourceProcessorImpl.java:254)
	at org.hibernate.boot.model.process.spi.MetadataBuildingProcess$1.processEntityHierarchies(MetadataBuildingProcess.java:230)
	at org.hibernate.boot.model.process.spi.MetadataBuildingProcess.complete(MetadataBuildingProcess.java:273)
	at org.hibernate.jpa.boot.internal.EntityManagerFactoryBuilderImpl.metadata(EntityManagerFactoryBuilderImpl.java:1224)
	at org.hibernate.jpa.boot.internal.EntityManagerFactoryBuilderImpl.build(EntityManagerFactoryBuilderImpl.java:1255)

By default, the entity name is the unqualified name of the entity class (i.e. the short class name excluding the package name). A different entity name can be set explicitly by using the name attribute of the Entity annotation:

@Entity(name="MyName")
public class MyEntity {

}
 */

@Entity(name = "demo2.city")
//////////////////////////
@Table(name = "cities")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class City {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private int population;

    public City(String name, int population) {
        this.name = name;
        this.population = population;
    }
}
