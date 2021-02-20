package com.github.klauskie.catalog.repository;

import com.github.klauskie.catalog.entity.MenuItem;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.Set;


@Repository
public interface MenuItemRepository extends JpaRepository<MenuItem, Long> {
    MenuItem findByItemId(String uuid);

    @Query("SELECT item FROM MenuItem item WHERE item.restaurant.restaurantId = :uuid")
    Set<MenuItem> findAllByRestaurantId(@Param(value = "uuid") String uuid);
}
