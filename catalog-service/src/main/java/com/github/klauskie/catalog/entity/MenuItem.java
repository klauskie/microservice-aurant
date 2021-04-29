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
@Table(name = "menuItem")
public class MenuItem implements Serializable {
    @Id
    @GeneratedValue(generator = "uuid4")
    @GenericGenerator(name = "uuid4", strategy = "org.hibernate.id.UUIDGenerator")
    private String itemId;

    private String name;
    private String description;
    private String price;
    private Boolean isAvailable;
    private Boolean isDisplayable;

    @OneToMany(mappedBy = "restaurant", cascade = {CascadeType.ALL})
    private Set<Category> categories;

    @ManyToOne
    @JoinColumn(name = "restaurant_id")
    Restaurant restaurant;

    /* Constructors */

    public MenuItem() {}

    public MenuItem(Map<String, Object> dict) {
        if (dict.containsKey(Constant.KEY_NAME)) {
            this.setName((String) dict.get(Constant.KEY_NAME));
        }
        if (dict.containsKey(Constant.KEY_DESCRIPTION)) {
            this.setDescription((String) dict.get(Constant.KEY_DESCRIPTION));
        }
        if (dict.containsKey(Constant.KEY_PRICE)) {
            this.setPrice((String) dict.get(Constant.KEY_PRICE));
        }
        if (dict.containsKey(Constant.KEY_IS_AVAILABLE)) {
            this.setIsAvailable((Boolean) dict.get(Constant.KEY_IS_AVAILABLE));
        }
        if (dict.containsKey(Constant.KEY_IS_DISPLAYABLE)) {
            this.setIsDisplayable((Boolean) dict.get(Constant.KEY_IS_DISPLAYABLE));
        }
    }

    public void update(MenuItem that) {
        if (that.getName() != null) {
            this.setName(that.getName());
        }
        if (that.getDescription() != null) {
            this.setDescription(that.getDescription());
        }
        if (that.getPrice() != null) {
            this.setPrice(that.getPrice());
        }
        if (that.getIsAvailable() != null) {
            this.setIsAvailable(that.getIsAvailable());
        }
        if (that.getIsDisplayable() != null) {
            this.setIsDisplayable(that.getIsDisplayable());
        }
    }

    /* Getters and Setters */

    public String getItemId() {
        return itemId;
    }

    public void setItemId(String id) {
        this.itemId = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getPrice() {
        return price;
    }

    public void setPrice(String price) {
        this.price = price;
    }

    public Boolean getIsAvailable() {
        return isAvailable;
    }

    public void setIsAvailable(Boolean available) {
        isAvailable = available;
    }

    public Boolean getIsDisplayable() {
        return isDisplayable;
    }

    public void setIsDisplayable(Boolean displayable) {
        isDisplayable = displayable;
    }

    @JsonIgnore
    public Set<Category> getCategories() {
        return categories;
    }

    public void setCategories(Set<Category> categories) {
        this.categories = categories;
    }

    public Restaurant getRestaurant() {
        return restaurant;
    }

    public void setRestaurant(Restaurant restaurant) {
        this.restaurant = restaurant;
    }

    /* Defaults */

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        MenuItem menuItem = (MenuItem) o;
        return itemId.equals(menuItem.itemId) && name.equals(menuItem.name);
    }

    @Override
    public int hashCode() {
        return Objects.hash(itemId, name);
    }
}
