package com.github.klauskie.catalog.entity;

import com.github.klauskie.catalog.util.Constant;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.io.Serializable;
import java.util.HashSet;
import java.util.Map;
import java.util.Objects;
import java.util.Set;

@Entity
@Table(name = "category")
public class Category implements Serializable {

    @Id
    @GeneratedValue(generator = "uuid4")
    @GenericGenerator(name = "uuid4", strategy = "org.hibernate.id.UUIDGenerator")
    private String categoryId;

    private String name;

    @ManyToOne
    @JoinColumn(name = "restaurant_id")
    Restaurant restaurant;

    @ManyToMany(cascade = {CascadeType.ALL})
    @JoinTable(
            name = "category_menuItem",
            joinColumns = @JoinColumn(name = "category_id"),
            inverseJoinColumns = @JoinColumn(name = "item_id"))
    Set<MenuItem> menuItems;


    /* Constructors */

    public Category() {}

    public Category(String name, Restaurant restaurant) {
        this.name = name;
        this.restaurant = restaurant;
        this.menuItems = new HashSet<>();
    }

    public Category(Map<String, String> dict) {
        if (dict.containsKey(Constant.KEY_NAME)) {
            this.setName(dict.get(Constant.KEY_NAME));
        }
    }

    public void update(Category that) {
        if (that.getName() != null) {
            this.setName(that.getName());
        }
    }

    /* Getters and Setters */

    public String getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(String id) {
        this.categoryId = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Restaurant getRestaurant() {
        return restaurant;
    }

    public void setRestaurant(Restaurant restaurant) {
        this.restaurant = restaurant;
    }

    public Set<MenuItem> getMenuItems() {
        return menuItems;
    }

    public void setMenuItems(Set<MenuItem> menuItems) {
        this.menuItems = menuItems;
    }

    public void appendItem(MenuItem item) {
        if (item != null) {
            this.menuItems.add(item);
        }
    }

    /* Defaults */

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Category category = (Category) o;
        return categoryId.equals(category.categoryId) && name.equals(category.name) && restaurant.equals(category.restaurant);
    }

    @Override
    public int hashCode() {
        return Objects.hash(categoryId, name, restaurant);
    }
}
