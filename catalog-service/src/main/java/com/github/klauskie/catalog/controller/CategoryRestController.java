package com.github.klauskie.catalog.controller;

import com.github.klauskie.catalog.entity.Category;
import com.github.klauskie.catalog.entity.MenuItem;
import com.github.klauskie.catalog.entity.Restaurant;
import com.github.klauskie.catalog.exception.GenericBadRequest;
import com.github.klauskie.catalog.repository.CategoryRepository;
import com.github.klauskie.catalog.repository.MenuItemRepository;
import com.github.klauskie.catalog.repository.RestaurantRepository;
import com.github.klauskie.catalog.util.Constant;
import org.springframework.cache.annotation.CachePut;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.List;
import java.util.Map;
import java.util.Set;

@CrossOrigin(origins = Constant.ORIGINS, allowedHeaders = Constant.ALLOWED_HEADERS)
@RestController
@RequestMapping("/api/*")
public class CategoryRestController {

    @Resource
    private CategoryRepository categoryRepository;

    @Resource
    private MenuItemRepository itemRepository;

    @Resource
    private RestaurantRepository restaurantRepository;

    @GetMapping("/category")
    public List<Category> getCategories() {
        return categoryRepository.findAll();
    }

    @GetMapping("/category/{UUID}")
    @Cacheable(value = "categories", key = "#uuid")
    public Category getCategoryByUUID(@PathVariable(value = "UUID") String uuid) {
        return categoryRepository.findByCategoryId(uuid);
    }

    // TODO: GET list of categories by restaurantID + queryParams
    // TODO: Cache list and refresh on user request
    // @Cacheable(value = "category-list", key = "#uuid")
    @GetMapping("/category/restaurant/{UUID}")
    public Set<Category> getCategoryByRestaurantUUID(@PathVariable(value = "UUID") String uuid) {
        return categoryRepository.findAllByRestaurantId(uuid);
    }

    @PostMapping("/category")
    public void createCategory(@RequestBody Map<String, String> requestMap) {
        Restaurant existingRestaurant = restaurantRepository.findByRestaurantId(requestMap.get(Constant.KEY_RESTAURANT_ID));
        if (existingRestaurant == null) {
            throw new GenericBadRequest("No Restaurant found with given Id.");
        }

        Category newCategory = new Category(requestMap);
        newCategory.setRestaurant(existingRestaurant);

        try {
            categoryRepository.save(newCategory);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }
    }

    @PutMapping("/category")
    @CachePut(value = "categories", key = "#requestMap?.get('categoryId')")
    public Category updateCategory(@RequestBody Map<String, String> requestMap) {
        Category existingCategory = categoryRepository.findByCategoryId(requestMap.get(Constant.KEY_CATEGORY_ID));
        if (existingCategory == null) {
            throw new GenericBadRequest("No Category found with given Id.");
        }

        Category newCategory = new Category(requestMap);
        existingCategory.update(newCategory);

        try {
            categoryRepository.save(existingCategory);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }
        return existingCategory;
    }

    @PutMapping("/category/add/{itemUUID}")
    @CachePut(value = "categories", key = "#requestMap?.get('categoryId')")
    public Category appendItemToCategory(@PathVariable(value = "itemUUID") String itemUUID, @RequestBody Map<String, String> requestMap) {
        Category existingCategory = categoryRepository.findByCategoryId(requestMap.get(Constant.KEY_CATEGORY_ID));
        if (existingCategory == null) {
            throw new GenericBadRequest("No Category found with given Id.");
        }

        MenuItem item = itemRepository.findByItemId(itemUUID);
        if (item == null) {
            throw new GenericBadRequest("No Item found with given Id.");
        }

        existingCategory.appendItem(item);

        try {
            categoryRepository.save(existingCategory);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }
        return existingCategory;
    }

    // TODO: DELETE category by Id
}
