# Resource Not Found

Playground for handling resource not found errors.

### Testing

```console
$ go test -v ./aws/internal/experimental/playground/notfound/...
=== RUN   Test_ResourceRead
=== RUN   Test_ResourceRead/Valid_thing_new
2021/02/12 14:33:06 &service.Thing{ThingId:(*string)(0xc0004b0060), Status:(*string)(0xc0004b0070)}
=== RUN   Test_ResourceRead/Valid_thing_old
2021/02/12 14:33:06 &service.Thing{ThingId:(*string)(0xc0004b0130), Status:(*string)(0xc0004b0140)}
=== RUN   Test_ResourceRead/Not_found_thing_new
=== RUN   Test_ResourceRead/Not_found_thing_old
2021/02/12 14:33:06 [WARN] Service Thing (thing-2) not found, removing from state
=== RUN   Test_ResourceRead/Empty_result_thing_new
=== RUN   Test_ResourceRead/Empty_result_thing_old
2021/02/12 14:33:06 [WARN] Service Thing (thing-0) not found, removing from state
=== RUN   Test_ResourceRead/Erroring_thing_new
=== RUN   Test_ResourceRead/Erroring_thing_old
--- PASS: Test_ResourceRead (0.00s)
    --- PASS: Test_ResourceRead/Valid_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Valid_thing_old (0.00s)
    --- PASS: Test_ResourceRead/Not_found_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Not_found_thing_old (0.00s)
    --- PASS: Test_ResourceRead/Empty_result_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Empty_result_thing_old (0.00s)
    --- PASS: Test_ResourceRead/Erroring_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Erroring_thing_old (0.00s)
=== RUN   Test_ResourceDelete
=== RUN   Test_ResourceDelete/Valid_thing
2021/02/12 14:33:06 [DEBUG] Waiting for state to become: []
2021/02/12 14:33:06 [TRACE] Waiting 200ms before next try
2021/02/12 14:33:07 [TRACE] Waiting 400ms before next try
=== RUN   Test_ResourceDelete/Not_found_thing
=== RUN   Test_ResourceDelete/Erroring_thing
--- PASS: Test_ResourceDelete (0.60s)
    --- PASS: Test_ResourceDelete/Valid_thing (0.60s)
    --- PASS: Test_ResourceDelete/Not_found_thing (0.00s)
    --- PASS: Test_ResourceDelete/Erroring_thing (0.00s)
PASS
ok  	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error	0.619s
?   	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error/deleter	[no test files]
?   	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error/finder	[no test files]
=== RUN   Test_ThingDeleted
=== RUN   Test_ThingDeleted/Valid_thing
2021/02/12 14:33:06 [DEBUG] Waiting for state to become: []
2021/02/12 14:33:06 [TRACE] Waiting 200ms before next try
2021/02/12 14:33:07 [TRACE] Waiting 400ms before next try
=== RUN   Test_ThingDeleted/Not_found_thing
2021/02/12 14:33:07 [DEBUG] Waiting for state to become: []
=== RUN   Test_ThingDeleted/Erroring_thing
2021/02/12 14:33:07 [DEBUG] Waiting for state to become: []
=== RUN   Test_ThingDeleted/Empty_result_thing
2021/02/12 14:33:07 [DEBUG] Waiting for state to become: []
--- PASS: Test_ThingDeleted (0.60s)
    --- PASS: Test_ThingDeleted/Valid_thing (0.60s)
    --- PASS: Test_ThingDeleted/Not_found_thing (0.00s)
    --- PASS: Test_ThingDeleted/Erroring_thing (0.00s)
    --- PASS: Test_ThingDeleted/Empty_result_thing (0.00s)
PASS
ok  	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/error/waiter	0.636s
?   	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/namevaluesfilters	[no test files]
=== RUN   Test_ResourceRead
=== RUN   Test_ResourceRead/Valid_thing_new
2021/02/12 14:33:10 &service.Thing{ThingId:(*string)(0xc00004a3e0), Status:(*string)(0xc00004a3f0)}
=== RUN   Test_ResourceRead/Valid_thing_old
2021/02/12 14:33:10 &service.Thing{ThingId:(*string)(0xc00004a4b0), Status:(*string)(0xc00004a4c0)}
=== RUN   Test_ResourceRead/Not_found_thing_new
=== RUN   Test_ResourceRead/Not_found_thing_old
2021/02/12 14:33:10 [WARN] Service Thing (thing-2) not found, removing from state
=== RUN   Test_ResourceRead/Empty_result_thing_new
=== RUN   Test_ResourceRead/Empty_result_thing_old
2021/02/12 14:33:10 [WARN] Service Thing (thing-0) not found, removing from state
=== RUN   Test_ResourceRead/Erroring_thing_new
=== RUN   Test_ResourceRead/Erroring_thing_old
--- PASS: Test_ResourceRead (0.00s)
    --- PASS: Test_ResourceRead/Valid_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Valid_thing_old (0.00s)
    --- PASS: Test_ResourceRead/Not_found_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Not_found_thing_old (0.00s)
    --- PASS: Test_ResourceRead/Empty_result_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Empty_result_thing_old (0.00s)
    --- PASS: Test_ResourceRead/Erroring_thing_new (0.00s)
    --- PASS: Test_ResourceRead/Erroring_thing_old (0.00s)
=== RUN   Test_ResourceDelete
=== RUN   Test_ResourceDelete/Valid_thing
2021/02/12 14:33:10 [DEBUG] Waiting for state to become: []
2021/02/12 14:33:10 [TRACE] Waiting 200ms before next try
2021/02/12 14:33:10 [TRACE] Waiting 400ms before next try
=== RUN   Test_ResourceDelete/Not_found_thing
2021/02/12 14:33:11 [DEBUG] Waiting for state to become: []
=== RUN   Test_ResourceDelete/Erroring_thing
--- PASS: Test_ResourceDelete (0.60s)
    --- PASS: Test_ResourceDelete/Valid_thing (0.60s)
    --- PASS: Test_ResourceDelete/Not_found_thing (0.00s)
    --- PASS: Test_ResourceDelete/Erroring_thing (0.00s)
PASS
ok  	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil	0.626s
?   	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/deleter	[no test files]
?   	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/finder	[no test files]
=== RUN   Test_ThingDeleted
=== RUN   Test_ThingDeleted/Valid_thing
2021/02/12 14:33:10 [DEBUG] Waiting for state to become: []
2021/02/12 14:33:10 [TRACE] Waiting 200ms before next try
2021/02/12 14:33:10 [TRACE] Waiting 400ms before next try
=== RUN   Test_ThingDeleted/Not_found_thing
2021/02/12 14:33:10 [DEBUG] Waiting for state to become: []
=== RUN   Test_ThingDeleted/Erroring_thing
2021/02/12 14:33:10 [DEBUG] Waiting for state to become: []
=== RUN   Test_ThingDeleted/Empty_result_thing
2021/02/12 14:33:10 [DEBUG] Waiting for state to become: []
--- PASS: Test_ThingDeleted (0.60s)
    --- PASS: Test_ThingDeleted/Valid_thing (0.60s)
    --- PASS: Test_ThingDeleted/Not_found_thing (0.00s)
    --- PASS: Test_ThingDeleted/Erroring_thing (0.00s)
    --- PASS: Test_ThingDeleted/Empty_result_thing (0.00s)
PASS
ok  	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/nil/waiter	0.626s
?   	github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service	[no test files]
```
