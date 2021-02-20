package com.github.klauskie.catalog.helper;

import com.github.klauskie.catalog.entity.Category;
import com.github.klauskie.catalog.entity.MenuItem;
import com.github.klauskie.catalog.entity.Restaurant;
import com.github.klauskie.catalog.repository.CategoryRepository;
import com.github.klauskie.catalog.repository.MenuItemRepository;
import com.github.klauskie.catalog.repository.RestaurantRepository;

import javax.annotation.Resource;
import java.util.*;

public class RestaurantCsvParserHelper {

    @Resource
    private CategoryRepository categoryRepository;

    @Resource
    private RestaurantRepository restaurantRepository;

    @Resource
    private MenuItemRepository itemRepository;

    private List<String[]> mData;
    private Map<String, Category> categoryMap;
    private Restaurant mRestaurant;
    private String restaurantUUID;

    public RestaurantCsvParserHelper(List<String[]> data, String uuid, RestaurantRepository rORM, CategoryRepository cORM, MenuItemRepository iORM) {
        this.mData = data;
        this.restaurantUUID = uuid;
        this.restaurantRepository = rORM;
        this.categoryRepository = cORM;
        this.itemRepository = iORM;

        this.categoryMap = new HashMap<>();
        this.mRestaurant = restaurantRepository.findByRestaurantId(uuid);
    }

    public void Run() {
        deleteEverythingInRestaurant();
        parse2Map();

        System.out.println("%%%%%% CATEGORY MAP: " + this.categoryMap.size());

        for (Map.Entry<String, Category> entry : this.categoryMap.entrySet()) {
            categoryRepository.save(entry.getValue());
        }
    }

    private void parse2Map() {
        for (String[] row : mData) {
            MenuItem item = ToMenuItem(row);
            List<Category> categories = ToCategorySet(row);
            if (item == null || categories == null) {
                continue;
            }

            for (Category c : categories) {
                if (this.categoryMap.containsKey(c.getName())) {
                    this.categoryMap.get(c.getName()).appendItem(item);
                } else {
                    c.appendItem(item);
                    this.categoryMap.put(c.getName(), c);
                }
            }
        }
    }

    private void deleteEverythingInRestaurant() {
        Set<Category> categories = categoryRepository.findAllByRestaurantId(this.restaurantUUID);
        Set<MenuItem> items = itemRepository.findAllByRestaurantId(this.restaurantUUID);
        categoryRepository.deleteInBatch(categories);
        itemRepository.deleteInBatch(items);
    }

    private MenuItem ToMenuItem(String[] row) {
        if (row.length < 6) {
            return null;
        }
        MenuItem item = new MenuItem();
        item.setName(row[0]);
        item.setDescription(row[1]);
        item.setPrice(row[2]);
        item.setIsAvailable(row[3].equals("true"));
        item.setIsDisplayable(row[4].equals("true"));
        item.setRestaurant(mRestaurant);
        return item;
    }

    private List<Category> ToCategorySet(String[] row) {
        if (row.length < 6) {
            return null;
        }
        String[] categoryNames = row[5].split("\\|");
        List<Category> categories = new ArrayList<>();
        for (String name : categoryNames) {
            categories.add(new Category(name.trim(), mRestaurant));
        }
        return categories;
    }
}
