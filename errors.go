package main

import "errors"

// ErrInsuficientArguments when there aren't enough arguments for a type
var ErrInsuficientArguments = errors.New("Insuficient arguments")

// ErrTooManyArguments when there are too many arguments for a type
var ErrTooManyArguments = errors.New("Too many arguments")

// ErrInvalidArgument when an argument cannot be parsed
var ErrInvalidArgument = errors.New("Invalid argument")

// ErrIndexVertexRequired when an `Index` doesn't have a `Vertex`
var ErrIndexVertexRequired = errors.New("Index vertex required")

// ErrMissingTexCoord when an `Index` does not have a `TexCoord` but a list of `TexCoord`s is provided
var ErrMissingTexCoord = errors.New("Missing texture coordinate")

// ErrMissingNormal when an `Index` does not have a `Normal` but a list of `Normal`s is provided
var ErrMissingNormal = errors.New("Missing normal")

// ErrVerticeNotFound when a vertice is not found in the vertices slice (for its index)
var ErrVerticeNotFound = errors.New("Vertice not found")

// ErrTexCoordNotFound when a texture coordinate is not found in the texture coordinates slice (for its index)
var ErrTexCoordNotFound = errors.New("Texture coordinate not found")

// ErrNormalNotFound when a normal is not found in the normals slice (for its index)
var ErrNormalNotFound = errors.New("Normal not found")

// ErrInvalidFormat when source does not pass format validation
var ErrInvalidFormat = errors.New("Invalid format")
