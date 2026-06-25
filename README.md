<img width="1400" height="420" alt="image" src="https://github.com/user-attachments/assets/ed106e1b-fd58-4464-9c9b-0827d4937b30" />
# Learning Go

My Go learning journey documented through code — from basics to building real projects.

## Structure

```
golang/
├── 01. Hello Go
├── 02. Variables
├── 03. Functions
├── 04. Pointers
├── 05. Arrays and Slices
├── 06. Control Flow
├── 07. Maps
├── 08. Structs
├── 09. Interfaces
├── 10. Errors
├── 11. Goroutines
└── Projects/
    ├── go-webserver
    └── go-crud
```

## Language Fundamentals

### [01. Hello Go](./01.%20Hello%20Go/)
First Go program. Understanding `package main`, `func main()`, and how Go files are structured.

### [02. Variables](./02.%20Variables/)
| File | Concept |
|------|---------|
| `01_basic_declaration` | `var` keyword |
| `02_multiple_vars` | Declaring multiple variables |
| `03_variables_types` | Type system |
| `04_reassignment` | Mutating variables |
| `05_static_typing` | Why Go is statically typed |
| `06_declaration_and_initialization` | `:=` shorthand |
| `07_zero_values` | Default values for each type |
| `08_strong_typing` | No implicit conversion |
| `09_multiple_declaration` | Group declarations |
| `10_constants` | `const` keyword |
| `11_constants_and_types` | Typed constants |
| `12_grouped_declarations` | `const` and `var` blocks |

### [03. Functions](./03.%20Functions/)
| File | Concept |
|------|---------|
| `01_basic_functions` | Defining and calling functions |
| `02_parameters` | Single parameter |
| `03_multiple_parameters` | Multiple parameters |
| `04_return` | Return values |
| `05_return_multiple_values` | Multiple return values |
| `06_named_return_values` | Named returns |
| `07_named_function_literals` | Functions as values |
| `08_closures` | Closures |

### [04. Pointers](./04.%20Pointers/)
| File | Concept |
|------|---------|
| `01_pointer_address` | `&` operator, memory addresses |
| `02_change_via_pointer` | `*` operator, dereferencing |
| `03_function_without_pointers` | Pass by value |
| `04_function_with_pointer` | Pass by pointer |

### [05. Arrays and Slices](./05.%20Arrays%20and%20Slices/)
| File | Concept |
|------|---------|
| `01_arrays` | Fixed-size arrays |
| `02_array_literal` | Array initialization |
| `03_slice_from_array` | Slicing arrays |
| `04_slice_modifies_array` | Slices share memory |
| `05_make_slice` | `make()` for slices |
| `06_append_and_capacity` | `append()`, cap vs len |
| `07_slice_capacity_overflow` | When Go reallocates |
| `08_delete_elements_in_slice` | Delete with `append` trick |
| `09_iteration` | `for range` over slices |

### [06. Control Flow](./06.%20Control%20Flow/)
| File | Concept |
|------|---------|
| `01_if` to `05` | if, else, else-if, conditions |
| `06_switch` | Switch statements |
| `07_for_loop_counter` | Classic for loop |
| `08_for_loop_condition` | While-style loop |
| `09_infinite_for_with_break` | Infinite loop + break |
| `10_for_range` | Range iteration |
| `11_for_range_break_continue` | break and continue |

### [07. Maps](./07.%20Maps/)
| File | Concept |
|------|---------|
| `01_basic_maps` | Map syntax |
| `02_check_key_exists` | Comma-ok idiom |
| `03_delete_from_map` | `delete()` |
| `04_for_range_maps` | Iterating maps |
| `05_empty_map_nil` | Nil maps |
| `06_make_map` | `make()` for maps |
| `07_maps_in_functions` | Passing maps |

### [08. Structs](./08.%20Structs/)
| File | Concept |
|------|---------|
| `01_struct_type` | Defining structs |
| `02_struct_literal` | Initializing structs |
| `03_struct_pass_by_value` | Structs are values |
| `04_struct_pass_by_pointer` | Pointer receivers |
| `05_function_as_field` | Function fields |
| `06_methods` | Methods on structs |
| `07_string_method` | Stringer interface |
| `08_struct_vs_map` | When to use which |

### [09. Interfaces](./09.%20Interfaces/)
| File | Concept |
|------|---------|
| `01_basic_interface` | Defining interfaces |
| `02_multiple_implementations` | Implicit implementation |
| `03_interface_as_parameter` | Polymorphism |
| `04_stringer_interface` | Built-in interfaces |
| `05_custom_slice_type` | Type aliases |
| `06_interface_and_struct` | Composition |
| `07_multiple_interfaces` | Implementing many interfaces |
| `08_slice_iteration` | Interface slices |
| `09_empty_interface` | `any` / `interface{}` |

### [10. Errors](./10.%20Errors/)
| File | Concept |
|------|---------|
| `01_panic_runtime_error` | What panic is |
| `02_avoiding_runtime_error` | Safe checks |
| `03_defer_named_function` | `defer` with named functions |
| `04_defer_anonymous_function` | `defer` with closures |
| `05_recover` | Recovering from panic |
| `06_return_error` | Returning errors |
| `07_create_error` | `errors.New()` |
| `08_custom_error_and_panic` | Custom error types |
| `09_function_returns_error` | Error handling patterns |
| `10_new_error` | `fmt.Errorf()` wrapping |

### [11. Goroutines](./11.%20Goroutines/)
| File | Concept |
|------|---------|
| `01_concurrency` | What concurrency means |
| `02_main_goroutine` | The main goroutine |
| `03_create_goroutine` | `go` keyword |
| `04_multiple_goroutines` | Running many goroutines |
| `05_waitgroup` | `sync.WaitGroup` |
| `06_channels` | Channel basics |
| `07_one_channel_multiple_goroutines` | Fan-out pattern |

---

## Projects

### [go-webserver](./Projects/go-webserver/)
A basic HTTP server using Go's standard library.

- Static file serving
- Form handling
- Route registration with `net/http`

**Run:**
```bash
cd Projects/go-webserver
go run main.go
```

---

### [go-crud](./Projects/go-crud/)
A production-structured REST API for managing movies.

**Architecture:** Handler → Service → Repository → PostgreSQL

**Stack:** `net/http` · PostgreSQL · `lib/pq` · `golang-migrate` · `godotenv`

**Endpoints:**

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/movies` | List all movies |
| GET | `/movies/{id}` | Get movie by ID |
| POST | `/movies` | Create a movie |
| PUT | `/movies/{id}` | Update a movie |
| DELETE | `/movies/{id}` | Delete a movie |

**Run:**
```bash
cd Projects/go-crud
cp .env.example .env   # fill in your Postgres credentials
migrate -path ./migrations -database "postgres://..." up
go run ./cmd/api/
```

---

## Progress

- [x] Language fundamentals (variables → goroutines)
- [x] HTTP server
- [x] REST API with PostgreSQL
- [ ] Middleware
- [ ] Testing
- [ ] Authentication
