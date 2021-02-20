package com.github.klauskie.catalog.entity;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.github.klauskie.catalog.util.Constant;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Map;
import java.util.Objects;
import java.util.Set;

@Entity
@Table(name = "restaurant")
public class Restaurant implements Serializable {

    @Id
    @GeneratedValue(generator = "uuid4")
    @GenericGenerator(name = "uuid4", strategy = "org.hibernate.id.UUIDGenerator")
    private String restaurantId;

    private String name;
    private String address;
    private String openHour;
    private String closeHour;

    @JsonIgnore
    @OneToMany(mappedBy = "restaurant")
    private Set<Category> categories;

    /* Constructors */

    public Restaurant() {}

    public Restaurant(Map<String, String> dict) {
        if (dict.containsKey(Constant.KEY_NAME)) {
            this.setName(dict.get(Constant.KEY_NAME));
        }

        if (dict.containsKey(Constant.KEY_ADDRESS)) {
            this.setAddress(dict.get(Constant.KEY_ADDRESS));
        }

        if (dict.containsKey(Constant.KEY_OPEN_HOUR)) {
            this.setOpenHour(dict.get(Constant.KEY_OPEN_HOUR));
        }

        if (dict.containsKey(Constant.KEY_CLOSE_HOUR)) {
            this.setCloseHour(dict.get(Constant.KEY_CLOSE_HOUR));
        }
    }

    public void update(Restaurant that) {
        this.setName(that.getName());
        this.setAddress(that.getAddress());
        this.setOpenHour(that.getOpenHour());
        this.setCloseHour(that.getCloseHour());
    }

    /* Getters and Setters */

    public String getRestaurantId() {
        return restaurantId;
    }

    public void setRestaurantId(String id) {
        this.restaurantId = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getOpenHour() {
        return openHour;
    }

    public void setOpenHour(String openHour) {
        this.openHour = openHour;
    }

    public String getCloseHour() {
        return closeHour;
    }

    public void setCloseHour(String closeHour) {
        this.closeHour = closeHour;
    }

    public Set<Category> getCategories() {
        return categories;
    }

    public void setCategories(Set<Category> categories) {
        this.categories = categories;
    }

    /* Defaults */

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Restaurant that = (Restaurant) o;
        return restaurantId.equals(that.restaurantId) && name.equals(that.name);
    }

    @Override
    public int hashCode() {
        return Objects.hash(restaurantId, name);
    }
}
