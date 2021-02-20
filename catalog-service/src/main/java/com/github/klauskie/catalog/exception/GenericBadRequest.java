package com.github.klauskie.catalog.exception;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.BAD_REQUEST)
public class GenericBadRequest extends RuntimeException {
    private static final String errorMsg = "Missing params";

    public GenericBadRequest() {
        super(errorMsg);
    }
    public GenericBadRequest(final String message) {
        super(message);
    }
}
