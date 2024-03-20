package utils

import (
	"errors"
)

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal server error")

	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("data not found")

	// ErrNoUpdatedData is an error for when no data is provided to update
	ErrNoUpdatedData = errors.New("no data to update")

	// ErrConflictingData is an error for when data conflicts with existing data
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")

	// ErrDuplicateData is an error for when data already exists in the system
	ErrDuplicateData = errors.New("email or username already exists")

	// ErrTokenDuration is an error for when the token duration format is invalid
	ErrTokenDuration = errors.New("invalid token duration format")

	// ErrTokenCreation is an error for when the token creation fails
	ErrTokenCreation = errors.New("error creating token")

	// ErrExpiredToken is an error for when the access token is expired
	ErrExpiredToken = errors.New("token has expired")

	// ErrInvalidToken is an error for when the access token is invalid
	ErrInvalidToken = errors.New("token is invalid")

	// ErrInvalidCredentials is an error for when the credentials are invalid
	ErrInvalidCredentials = errors.New("invalid email or password")

	// ErrEmptyAuthorizationHeader is an error for when the authorization header is empty
	ErrEmptyAuthorizationHeader = errors.New("authorization header is not provided")

	// ErrInvalidAuthorizationHeader is an error for when the authorization header is invalid
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")

	// ErrInvalidAuthorizationType is an error for when the authorization type is invalid
	ErrInvalidAuthorizationType = errors.New("authorization type is not supported")

	// ErrUnauthorized is an error for when the user is unauthorized
	ErrUnauthorized = errors.New("user is unauthorized to access the resource")

	// ErrForbidden is an error for when the user is forbidden to access the resource
	ErrForbidden = errors.New("user is forbidden to access the resource")

	// ErrFormatFile is an error for when the input file format cannot be opened
	ErrFormatFile = errors.New("unable to open input file format")

	// ErrFileSize is an error for when the input file size exceeds the allowed limit
	ErrFileSize = errors.New("input file size format not allowed, size exceeds")

	// ErrFileExtension is an error for when the input file type is not allowed
	ErrFileExtension = errors.New("input file type format not allowed")
	ErrImageRequired = errors.New("photo_url is required")
)
