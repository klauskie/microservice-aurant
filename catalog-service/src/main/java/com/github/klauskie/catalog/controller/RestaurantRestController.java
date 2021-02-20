package com.github.klauskie.catalog.controller;

import com.github.klauskie.catalog.entity.Category;
import com.github.klauskie.catalog.entity.Restaurant;
import com.github.klauskie.catalog.exception.GenericBadRequest;
import com.github.klauskie.catalog.helper.RestaurantCsvParserHelper;
import com.github.klauskie.catalog.repository.CategoryRepository;
import com.github.klauskie.catalog.repository.MenuItemRepository;
import com.github.klauskie.catalog.repository.RestaurantRepository;
import com.github.klauskie.catalog.util.Constant;
import com.opencsv.CSVReader;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import javax.annotation.Resource;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.Reader;
import java.util.List;
import java.util.Map;
import java.util.Set;

@CrossOrigin(origins = Constant.ORIGINS, allowedHeaders = Constant.ALLOWED_HEADERS)
@RestController
@RequestMapping("/api/*")
public class RestaurantRestController {

    @Resource
    private CategoryRepository categoryRepository;

    @Resource
    private RestaurantRepository restaurantRepository;

    @Resource
    private MenuItemRepository itemRepository;

    @GetMapping("/restaurant")
    public List<Restaurant> getRestaurants() {
        return restaurantRepository.findAll();
    }

    @GetMapping("/restaurant/{UUID}")
    public Restaurant getRestaurantByUUID(@PathVariable(value = "UUID") String uuid) {
        return restaurantRepository.findByRestaurantId(uuid);
    }

    @GetMapping("/restaurant/categories/{UUID}")
    public Set<Category> getCategoryByRestaurantUUID(@PathVariable(value = "UUID") String uuid) {
        Restaurant restaurant = restaurantRepository.findByRestaurantId(uuid);
        if (restaurant == null) {
            throw new GenericBadRequest("No Restaurant found with given Id.");
        }
        return restaurant.getCategories();
    }

    @PostMapping("/restaurant")
    public void createRestaurant(@RequestBody Map<String, String> requestMap) {
        Restaurant newRestaurant = new Restaurant(requestMap);

        try {
            restaurantRepository.save(newRestaurant);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }
    }

    @PutMapping("/restaurant")
    public void updateRestaurant(@RequestBody Map<String, String> requestMap) {
        Restaurant existingRestaurant = restaurantRepository.findByRestaurantId(requestMap.get(Constant.KEY_RESTAURANT_ID));
        if (existingRestaurant == null) {
            throw new GenericBadRequest("No Restaurant found with given Id.");
        }

        Restaurant newRestaurant = new Restaurant(requestMap);
        existingRestaurant.update(newRestaurant);

        try {
            restaurantRepository.save(existingRestaurant);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }
    }

    // TODO: PUT import csv with category and menu
    @PostMapping("/restaurant/upload-csv/{restaurantUUID}")
    public List<String[]> handleCsvDataImport(@RequestParam("file") MultipartFile file,
                                    @PathVariable(value = "restaurantUUID") String uuid) {

        Restaurant existingRestaurant = restaurantRepository.findByRestaurantId(uuid);
        if (existingRestaurant == null) {
            throw new GenericBadRequest("No Restaurant found with given Id.");
        }

        if (file.isEmpty()) {
            throw new GenericBadRequest("Please select a CSV file to upload.");
        }

        List<String[]> records;

        try (Reader reader = new BufferedReader(new InputStreamReader(file.getInputStream()))) {
            CSVReader csvReader = new CSVReader(reader);
            records = csvReader.readAll();

        } catch (Exception ex) {
            throw new GenericBadRequest(ex.getMessage());
        }

        RestaurantCsvParserHelper helper = new RestaurantCsvParserHelper(records, uuid, restaurantRepository, categoryRepository, itemRepository);
        try {
            helper.Run();
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }

        return records;
    }
    // TODO: DELETE restaurant by Id
}
