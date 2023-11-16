# map-struct - a go library

MapStruct is a lightweight Go library that provides functions for mapping data from a map to a struct using reflection.

## Installation

To use MapStruct in your Go project, simply run:

 ```bash
    go get github.com/b0a04gl/map-struct
 ```


## Usage

Here's how to get started with this query builder:

1. Import the package:

    ```go
    import "github.com/b0a04gl/map-struct"
    ```

2. Sample mapping 

    ```go
        package main

        import (
            "fmt"
            "github.com/b0a04gl/map-struct"
        )

            type Person struct {
                Name    string
                Age     int
                Address string
            }

            func main() {
                // Example data to be mapped
                data := map[string]interface{}{
                "Name":    "Tommy Shelby",
                "Age":     34,
                "Address": "441, Baking street, Buckingham",
                }

            // Create an instance of the target struct
            var person Person

            // Map data to the struct using MapToStruct function
            err := mapstruct.MapToStruct(data, &person)

            // Check for errors
            if err != nil {
            fmt.Println("Error:", err)
            } else {
            // Use the mapped struct
                fmt.Println("Mapped Person:", person)
            }
        }
     ```

