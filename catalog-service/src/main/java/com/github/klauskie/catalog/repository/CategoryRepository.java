package com.github.klauskie.catalog.repository;

import com.github.klauskie.catalog.entity.Category;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.Set;

public interface CategoryRepository extends JpaRepository<Category, Long> {
    Category findByCategoryId(String uuid);

    @Query("SELECT c FROM Category c WHERE c.restaurant.restaurantId = :uuid")
    Set<Category> findAllByRestaurantId(@Param(value = "uuid") String uuid);
}
