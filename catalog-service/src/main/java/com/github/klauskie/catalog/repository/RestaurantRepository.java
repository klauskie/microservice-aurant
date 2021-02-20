package com.github.klauskie.catalog.repository;

import com.github.klauskie.catalog.entity.Restaurant;
import org.springframework.data.jpa.repository.JpaRepository;

public interface RestaurantRepository extends JpaRepository<Restaurant, Long> {
    Restaurant findByRestaurantId(String uuid);
}
