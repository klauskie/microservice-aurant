package com.github.klauskie.catalog.controller;

import com.github.klauskie.catalog.entity.Category;
import com.github.klauskie.catalog.entity.MenuItem;
import com.github.klauskie.catalog.entity.Restaurant;
import com.github.klauskie.catalog.exception.GenericBadRequest;
import com.github.klauskie.catalog.repository.CategoryRepository;
import com.github.klauskie.catalog.repository.MenuItemRepository;
import com.github.klauskie.catalog.repository.RestaurantRepository;
import com.github.klauskie.catalog.util.Constant;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.*;

@CrossOrigin(origins = Constant.ORIGINS, allowedHeaders = Constant.ALLOWED_HEADERS)
@RestController
@RequestMapping("/api/*")
public class MenuItemRestController {

    @Resource
    private CategoryRepository categoryRepository;

    @Resource
    private MenuItemRepository itemRepository;

    @Resource
    private RestaurantRepository restaurantRepository;

    @GetMapping("/item")
    public List<MenuItem> getMenuItems() {
        return itemRepository.findAll();
    }

    // TODO: GET items by restaurantID and queryParams (categoryId, pagination, ...)
    @GetMapping("/item/{restaurantID}")
    public Set<MenuItem> getMenuItems_withQueryParams(@PathVariable(value = "restaurantID") String restaurantUUID,
                                                      @RequestParam(value="category", required=false) Optional<String> categoryUUID) {

        Restaurant restaurant = restaurantRepository.findByRestaurantId(restaurantUUID);
        if (restaurant == null) {
            throw new GenericBadRequest("No Restaurant found with given Id.");
        }

        Set<MenuItem> items = new HashSet<>();

        if (categoryUUID.isPresent()) {
            Category category = categoryRepository.findByCategoryId(categoryUUID.get());
            if (category != null) {
                items = category.getMenuItems();
            }
        } else {
            items = itemRepository.findAllByRestaurantId(restaurantUUID);
        }
        return items;
    }

    @PostMapping("/item")
    public void createMenuItem(@RequestBody Map<String, Object> requestMap) {
        if (!requestMap.containsKey(Constant.KEY_RESTAURANT_ID)) {
            throw new GenericBadRequest("No Restaurant Id provided.");
        }
        Restaurant existingRestaurant = restaurantRepository.findByRestaurantId((String) requestMap.get(Constant.KEY_RESTAURANT_ID));
        if (!requestMap.containsKey(Constant.KEY_CATEGORY_ID)) {
            throw new GenericBadRequest("No Category Id provided.");
        }
        Category existingCategory = categoryRepository.findByCategoryId((String) requestMap.get(Constant.KEY_CATEGORY_ID));

        if (existingRestaurant == null) {
            throw new GenericBadRequest("No Restaurant found with given Id.");
        }

        if (existingCategory == null) {
            throw new GenericBadRequest("No Categories found with given Id.");
        }

        MenuItem newItem = new MenuItem(requestMap);
        newItem.setRestaurant(existingRestaurant);
        try {
            itemRepository.save(newItem);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }

        try {
            existingCategory.appendItem(newItem);
            categoryRepository.save(existingCategory);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }
    }

    @PostMapping("/item-definition")
    public Map<String, MenuItem> itemListDefinition(@RequestBody List<String> itemListReq) {
        Map<String, MenuItem> menuItemMap = new HashMap<>();
        for (String itemID : itemListReq) {
            if (!menuItemMap.containsKey(itemID)) {
                menuItemMap.put(itemID, itemRepository.findByItemId(itemID));
            }
        }
        return menuItemMap;
    }

    @PutMapping("/item")
    public void updateMenuItem(@RequestBody Map<String, Object> requestMap) {
        MenuItem existingItem = itemRepository.findByItemId((String) requestMap.get(Constant.KEY_ITEM_ID));
        if (existingItem == null) {
            throw new GenericBadRequest("No Item found with given Id.");
        }

        MenuItem newItem = new MenuItem(requestMap);
        existingItem.update(newItem);

        try {
            itemRepository.save(existingItem);
        } catch (DataIntegrityViolationException e) {
            throw new GenericBadRequest(e.getMessage());
        }
    }

    // TODO: DELETE item by itemId
}
